# Tests

> Last updated: 2026-01-11

## Test Framework

Standard Go testing (`testing` package) with `testify` for assertions.

## Running Tests

```bash
# Run all tests
make test

# Or directly
GOOS=windows GO111MODULE=on GOARCH=amd64 go test -v ./go/...

# Run specific package tests
go test -v ./pkg/base/session/...
go test -v ./pkg/cluster/compute/cluster/...

# Run with verbose output
go test -v -run TestSessionConnect ./pkg/base/session/
```

## Test Organization

Tests are co-located with source files:

```
pkg/
├── base/
│   ├── instance/
│   │   ├── instancemanager.go
│   │   └── instancemanager_test.go
│   └── session/
│       ├── session.go
│       └── session_test.go
├── cluster/
│   └── compute/
│       ├── affinityrule/
│       │   ├── affinityrule.go
│       │   └── affinityrule_test.go
│       ├── cluster/
│       │   ├── cluster.go
│       │   └── cluster_test.go
│       └── ...
├── errors/
│   ├── errors.go
│   └── errors_test.go
└── hardware/
    └── host/
        ├── computersystem.go
        ├── computersystem_test.go
        ├── memory.go
        ├── memory_test.go
        └── ...
```

## Test Files

| File | Tests |
|------|-------|
| `pkg/base/session/session_test.go` | Session creation and connection |
| `pkg/base/instance/instancemanager_test.go` | Instance lifecycle management |
| `pkg/cluster/compute/cluster/cluster_test.go` | Cluster operations |
| `pkg/cluster/compute/clusternode/clusternode_test.go` | Node operations |
| `pkg/cluster/compute/resourcegroup/resourcegroup_test.go` | Resource groups |
| `pkg/cluster/compute/resource/resource_test.go` | Cluster resources |
| `pkg/cluster/compute/affinityrule/affinityrule_test.go` | Affinity rules |
| `pkg/cluster/compute/groupset/groupset_test.go` | Availability sets |
| `pkg/cluster/network/clusternetwork/clusternetwork_test.go` | Cluster networking |
| `pkg/cluster/service/clusterservice_test.go` | Cluster service |
| `pkg/cluster/storage/clustersharedvolume/clustersharedvolume_test.go` | CSV operations |
| `pkg/errors/errors_test.go` | Error handling |
| `pkg/hardware/host/computersystem_test.go` | Computer system info |
| `pkg/hardware/host/memory_test.go` | Memory info |
| `pkg/hardware/host/operatingsystem_test.go` | OS info |

## Writing New Tests

```go
package mypackage

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMyFunction(t *testing.T) {
    // Setup
    host := host.NewWmiLocalHost()

    // Test
    result, err := MyFunction(host)

    // Assert
    require.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, expected, result.Value)
}
```

## Test Requirements

Most tests require:
1. **Windows environment** - Tests use live WMI
2. **Administrator privileges** - For WMI access
3. **Hyper-V installed** - For virtualization tests
4. **Failover Clustering** - For cluster tests
5. **Active cluster** - For cluster-related tests

## Test Dependencies

```go
require (
    github.com/stretchr/testify v1.7.0
)
```

## Notes

- Tests run against live WMI services (no mocking)
- Some tests may require specific Windows Server roles
- Tests should clean up created resources
- Use `t.Skip()` for environment-specific tests
