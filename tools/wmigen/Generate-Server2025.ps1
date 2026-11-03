# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [string]$OutputDirectory,

    [Parameter(Mandatory = $true)]
    [ValidatePattern('^[0-9a-f]{40}$')]
    [string]$SourceCommit,

    [string]$GeneratorPath = (Join-Path $PSScriptRoot 'bin\Release\wmigen.exe')
)

$ErrorActionPreference = 'Stop'
$os = Get-CimInstance -ClassName Win32_OperatingSystem
if ($os.ProductType -eq 1 -or $os.BuildNumber -ne '26100' -or $os.Caption -notlike '*Windows Server 2025*') {
    throw 'Generation requires a genuine Windows Server 2025 host (build 26100), not a client OS or another Server release'
}

$clustering = Get-WindowsFeature -Name Failover-Clustering
if (-not $clustering.Installed) {
    throw 'Failover Clustering must already be installed; this script does not install roles or restart the host'
}
Get-CimClass -Namespace root\MSCluster -ClassName MSCluster_Cluster | Out-Null

$generator = Get-Item -LiteralPath $GeneratorPath
if ($generator.PSIsContainer) {
    throw 'GeneratorPath must identify a built wmigen executable'
}
$output = [IO.Path]::GetFullPath($OutputDirectory)
if (Test-Path -LiteralPath $output) {
    throw 'OutputDirectory must be a new staging directory; existing bindings are never overwritten'
}
New-Item -ItemType Directory -Path $output | Out-Null
$bindings = Join-Path $output 'server2025'
$log = Join-Path $output 'generation.log'
& $generator.FullName -generatesource -format GO -ns root -recurse `
    -rootmodulepath github.com/microsoft/wmi/server2025 -outdir $bindings *> $log
if ($LASTEXITCODE -ne 0) {
    throw "wmigen failed with exit code $LASTEXITCODE; inspect $log. Partial output is not qualified."
}

$files = @(Get-ChildItem -LiteralPath $bindings -Filter '*.go' -Recurse -File)
if ($files.Count -eq 0) {
    throw 'wmigen returned success without generating any Go bindings'
}
$clusterFiles = @(Get-ChildItem -LiteralPath (Join-Path $bindings 'root\mscluster') -Filter '*.go' -File)
if ($clusterFiles.Count -eq 0) {
    throw 'Failover Clustering bindings are missing from the generated output'
}
$namespaces = @($files | Group-Object DirectoryName | Sort-Object Name | ForEach-Object {
    [ordered]@{
        path = $_.Name.Substring($bindings.Length + 1).Replace('\', '/')
        files = $_.Count
    }
})
$version = Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion'
[ordered]@{
    os = [ordered]@{
        caption = $os.Caption
        version = $os.Version
        build = $os.BuildNumber
        revision = $version.UBR
        architecture = $os.OSArchitecture
    }
    generatorSourceCommit = $SourceCommit
    generatorSha256 = (Get-FileHash -LiteralPath $generator.FullName -Algorithm SHA256).Hash.ToLowerInvariant()
    generatedAtUtc = [DateTime]::UtcNow.ToString('o')
    rootModulePath = 'github.com/microsoft/wmi/server2025'
    failoverClusteringInstalled = $true
    traversal = 'root, recursive'
    legacyExclusions = @('root/rsop/user', 'localized ms_409, ms_509 and ms_501 namespaces')
    generatedFiles = $files.Count
    namespaces = $namespaces
    qualification = 'Generated from local class metadata; formatting and Go compilation must still pass'
} | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $output 'generation.json') -Encoding UTF8
Write-Output "Generated $($files.Count) files in $($namespaces.Count) namespaces. Format and compile before checking in."
