# Usage

> Last updated: 2026-01-11

## Prerequisites

- **Go 1.24.3+**
- **Windows** - This library only works on Windows (uses COM/OLE)
- **Administrator privileges** - Required for most WMI operations
- **Hyper-V role** - For virtualization APIs
- **Failover Clustering feature** - For cluster APIs

## Installation

```bash
go get github.com/microsoft/wmi
```

## Building

```bash
# Build all packages (Windows only)
make library

# Or directly with Go
GO111MODULE=on GOARCH=amd64 GOOS=windows go build ./...
```

## Running Tests

```bash
# Run tests
make test

# Or directly
GOOS=windows GO111MODULE=on GOARCH=amd64 go test -v ./go/...
```

## Code Formatting

```bash
make format
# Or
gofmt -s -w pkg
```

## Common Usage Patterns

### Creating a WMI Session

```go
import (
    "github.com/microsoft/wmi/pkg/base/host"
    "github.com/microsoft/wmi/pkg/base/session"
)

// Local connection
whost := host.NewWmiLocalHost()
sess, err := session.GetHostSession("root\\virtualization\\v2", whost)

// Remote connection
whost := host.NewWmiHost("server01", "domain", "username", "password")
sess, err := session.GetHostSession("root\\virtualization\\v2", whost)
```

### Querying Instances

```go
import (
    "github.com/microsoft/wmi/pkg/base/query"
    v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

// Query VMs
q := query.NewWmiQuery("Msvm_ComputerSystem")
q.AddFilter("Caption", "Virtual Machine")

instances, err := sess.QueryInstances(q)
for _, inst := range instances {
    vm, _ := v2.NewMsvm_ComputerSystemEx1(inst)
    name, _ := vm.GetPropertyElementName()
    fmt.Println(name)
}
```

### Using High-Level APIs

```go
import (
    "github.com/microsoft/wmi/pkg/cluster/compute/cluster"
    "github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
)

// Cluster operations
cl, err := cluster.GetCluster(whost)
nodes, err := cl.GetNodes()

// Virtual disk operations
vhd, err := disk.GetVirtualHardDisk(whost, vhdPath)
```

### Cleanup

```go
// Always close instances to release COM resources
defer instance.Close()
defer instances.Close()

// Stop WMI when done
defer session.StopWMI()
```

## Configuration

### Environment Variables

| Variable | Description |
|----------|-------------|
| `GO111MODULE` | Set to `on` for module mode |
| `GOOS` | Must be `windows` |
| `GOARCH` | Typically `amd64` |

### WMI Namespaces

Common namespaces used:
- `root\virtualization\v2` - Hyper-V
- `root\cimv2` - Standard WMI classes
- `root\MSCluster` - Failover Clustering
- `root\microsoft\windows\cluster` - Modern cluster API
- `root\standardcimv2` - Network configuration

## Cross-Compilation Note

This library **cannot be cross-compiled** from Linux/macOS because it depends on Windows COM interfaces through `go-ole`. Build must happen on Windows.
