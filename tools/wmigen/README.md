# WMI code generator

`wmigen` reads class metadata from the local Windows WMI repository and emits Go
or C# wrappers. Some relationship discovery also enumerates instances, so
generation is not strictly metadata-only. The Microsoft MIT-licensed generator is
maintained here alongside its Go runtime and generated bindings.

## Build

On Windows, install Visual Studio Build Tools with MSBuild and the .NET Framework
4.8 targeting pack. From a Developer PowerShell prompt at the repository root:

```powershell
msbuild .\tools\wmigen\wmicodegen.csproj /p:Configuration=Release
.\tools\wmigen\bin\Release\wmigen.exe -help
```

The dedicated [GitHub Actions build workflow](../../.github/workflows/wmigen.yml)
builds the generator, runs its unit tests and checks its help command on Windows. It runs for relevant
pull requests and pushes to `master`, and can also be started manually. Its
`wmigen-windows-net48` artifact contains a ZIP and SHA256 sidecar; the ZIP includes
the executable, configuration, this README and MIT license. Running the executable
requires Windows and .NET Framework 4.8, but not MSBuild.

Building the generator does not query WMI or install Windows roles. Existing Go
and security pipelines remain separate.

## Unit tests

On Windows, use the build prerequisites above plus the .NET SDK 8 or later.
From the repository root in Developer PowerShell:

```powershell
msbuild .\tools\wmigen\wmicodegen.csproj /p:Configuration=Release
dotnet restore .\tools\wmigen\tests\wmigen.Tests.csproj --source https://api.nuget.org/v3/index.json
dotnet test .\tools\wmigen\tests\wmigen.Tests.csproj --configuration Release --no-restore --logger "trx;LogFileName=wmigen.trx"
```

Run each command only after the previous one succeeds. The test project references
the built Release executable rather than rebuilding the legacy project with
`dotnet` MSBuild. MSTest, its adapter and the .NET test SDK are pinned public
NuGet test-only dependencies; the shipped executable has no new dependencies.
Results are written under `tools/wmigen/tests/TestResults/`.

The focused suite exercises shared qualifier/signature helpers, enum wire values,
exception propagation, C# constructors, properties, parameters and enum files,
plus Go type widths, identifier repairs, enum collisions, namespace imports and
package naming. CLI cases verify that help short-circuits generation even after
generation options. Property tests
use an unbound, unsaved in-memory `ManagementClass` schema; they do not query or
save WMI metadata, enumerate instances or invoke provider methods. Other emitter
tests use an in-memory namespace instead of the discovery-performing
`GOWmiNamespace` constructor. C# emitter cases do not require its separate generated
runtime. No Hyper-V/cluster roles, lab hosts or cloud access
are required. The tests do not assert that every Go keyword is handled: in
particular, the legacy emitter's `return` identifier defect remains a separate
follow-up.

The dedicated build job runs these tests before packaging, including when reused
by the release workflow. The tests do not run the generator's WMI discovery paths.

## Release

The separate [release workflow](../../.github/workflows/wmigen-release.yml) runs
when a maintainer pushes a tag of the form `wmigen/vX.Y.Z`, for example
`wmigen/v1.0.0`. The tagged commit must already be merged into `master`.

The workflow reuses the same Windows build, unit tests, help smoke test and packaging steps.
Only after those succeed does it verify the ZIP checksum and create a GitHub
Release with the ZIP and SHA256 sidecar. It requires no external feed, credentials
or service connection: only the publishing job has repository-content write
permission through `GITHUB_TOKEN`. Actions are pinned to immutable commit hashes.
Tool releases do not replace the Go library's "Latest" release.

The workflow does not create or push tags, generate schema bindings, sign binaries
or install providers. Creating the workflows does not publish a release. If
publishing fails, inspect the run and any existing draft before retrying; the
workflow does not overwrite existing releases or assets.

This follows GitHub's guidance on
[reusing workflows](https://docs.github.com/en/actions/how-tos/reuse-automations/reuse-workflows),
[release management](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)
and [least-privilege workflow permissions](https://docs.github.com/en/actions/reference/security/secure-use).

## Generate Go bindings

Run on the Windows release whose schemas you intend to capture:

```powershell
.\tools\wmigen\bin\Release\wmigen.exe `
    -generatesource -format GO -ns root/cimv2 `
    -rootmodulepath github.com/microsoft/wmi/server2019 `
    -outdir .\server2019
```

`-rootmodulepath` changes generated import paths; it does **not** select an OS or
connect to a different machine. Never label output from a workstation or another
Server release as Server 2025. Providers depend on the roles and features already
installed on the generation host.

Generate into a fresh staging directory, inspect the output, run `gofmt`, and
compile the generated packages before replacing checked-in bindings. Do not
reconfigure a shared host merely to make additional providers available.

Supported generation options:

| Option | Meaning |
|---|---|
| `-generatesource` | Generate source files. |
| `-ns root/cimv2,root/standardcimv2` | Comma-separated local WMI namespaces. |
| `-recurse` | Include child namespaces. |
| `-format GO` | Go output; `CS` retains the legacy C# emitter. |
| `-rootmodulepath <import-prefix>` | Go import prefix; defaults to `github.com/microsoft/wmi/server2019`. |
| `-outdir <directory>` | Output directory; defaults to the current directory. |
| `-help` | Display the original command-line help. |

This initial import preserves the existing generator implementation, including its
legacy command-line behavior. The advertised `-wmiclass` option is not implemented;
namespace generation processes all classes. The C# emitter is retained for
compatibility, but its runtime is not supplied by this Go repository.
