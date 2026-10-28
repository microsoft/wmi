# Windows Server 2025 WMI bindings

This snapshot was generated on Windows Server 2025 Datacenter, build
10.0.26100.32995. It contains 12,225 Go files across 76 namespaces, including
168 files in `root/mscluster`.

[generation.json](generation.json) records the source OS, generator commit and
SHA256, capture time, and complete namespace/file-count inventory. Providers
depend on installed roles and hardware; this snapshot includes host-specific
vendor providers and is not an inventory of every possible Server 2025 provider.
No instance values, host names, or credentials are recorded in the provenance.

## Imports

Use the namespace-specific packages, for example:

```go
import (
    "github.com/microsoft/wmi/server2025/root/cimv2"
    "github.com/microsoft/wmi/server2025/root/mscluster"
    virtualization "github.com/microsoft/wmi/server2025/root/virtualization/v2"
)
```

These are schema bindings. Provider operations still require the appropriate
Windows host, roles, permissions, and runtime state.

## Regeneration

Follow the [Server 2025 generation instructions](../tools/wmigen/README.md#windows-server-2025)
on a genuine Server 2025 host with Failover Clustering already installed. The
wrapper traverses `root` recursively into a new staging directory. It does not
install roles, configure a cluster, or restart the host.

The legacy exclusions are `root/rsop/user` and localized `ms_409`, `ms_509`,
and `ms_501` namespaces. Relationship discovery can read instances, but the
generated wrappers contain class/property/method definitions, not an instance
export.

Before replacing this directory, inspect the provenance and generated files,
then run from the repository root:

```powershell
gofmt -s -w .\server2025
$env:GOOS = 'windows'
go build -p=1 ./...
```

Compilation is not a functional test of every provider method. The existing
host-dependent integration tests can alter VMs and storage; do not run them on a
shared host without a separate, explicitly authorized test plan.
