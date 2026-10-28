# WMI code generator

`wmigen` reads **class metadata**, not instance data, from the local Windows WMI
repository and emits Go or C# wrappers. The Microsoft MIT-licensed generator is
maintained here alongside its Go runtime and generated bindings.

## Build

On Windows, install Visual Studio Build Tools with MSBuild and the .NET Framework
4.8 targeting pack. From a Developer PowerShell prompt at the repository root:

```powershell
msbuild .\tools\wmigen\wmicodegen.csproj /p:Configuration=Release
.\tools\wmigen\bin\Release\wmigen.exe -help
```

The existing Windows pipeline builds the generator and checks its help command.
Building the generator does not require querying WMI or installing Windows roles.

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
