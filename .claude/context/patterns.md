# Patterns

> Last updated: 2026-01-11

## Design Patterns

### 1. Instance Wrapper Pattern

Generated WMI classes embed `*cim.WmiInstance` and provide typed property accessors:

```go
// server2019/root/virtualization/v2/Msvm_ComputerSystem.go
type Msvm_ComputerSystem struct {
    *CIM_ComputerSystem  // Embedded parent class

    EnhancedSessionModeState uint16
    ProcessID                uint32
    // ...
}

func (instance *Msvm_ComputerSystem) GetPropertyProcessID() (value uint32, err error) {
    retValue, err := instance.GetProperty("ProcessID")
    // type conversion...
    return
}
```

### 2. Constructor Variants

Generated classes provide multiple constructors:
- `NewXxxEx1(instance *cim.WmiInstance)` - Wrap existing WMI instance
- `NewXxxEx6(host, namespace, user, pass, domain, query)` - Query and create

### 3. Collection Types

Entities commonly have collection types for batch operations:

```go
// pkg/virtualization/core/storage/disk/virtualharddiskcollection.go
type VirtualHardDiskCollection []*VirtualHardDisk

func (c *VirtualHardDiskCollection) Close() {
    for _, disk := range *c {
        disk.Close()
    }
}
```

### 4. Session Factory Pattern

Sessions are obtained through the session package with caching:

```go
// pkg/base/session/session.go
session, err := session.GetSession(namespaceName, serverName, domain, userName, password)

// Or via host configuration
session, err := session.GetHostSession(namespaceName, whost)
```

### 5. Query Builder

Queries constructed with the query package:

```go
// pkg/base/query/query.go
q := query.NewWmiQuery("Msvm_ComputerSystem")
q.AddFilter("Name", vm.Name)
```

## Code Organization

### Package Structure
- `go/wmi/` - Pure interfaces (no implementation)
- `pkg/wmiinstance/` - COM/OLE implementation of interfaces
- `pkg/base/` - Infrastructure utilities
- `pkg/<domain>/` - High-level domain APIs
- `server<version>/root/<namespace>/` - Generated WMI classes

### File Naming
- `<entity>.go` - Main entity implementation
- `<entity>collection.go` - Collection type
- `<entity>_test.go` - Unit tests
- `constant.go` - Domain constants

## Naming Conventions

### Types
- WMI classes: `Msvm_*`, `CIM_*`, `Win32_*` (matching WMI names)
- High-level wrappers: `VirtualMachine`, `VirtualHardDisk` (Go-idiomatic)
- Collections: `*Collection` suffix

### Methods
- Property getters: `GetProperty<Name>()` (generated), `Get<Name>()` (high-level)
- Property setters: `SetProperty<Name>()` (generated)
- WMI methods: Match WMI method names

### Variables
- WMI hosts: `whost *host.WmiHost`
- Sessions: `session *wmi.WmiSession`
- Instances: `instance *cim.WmiInstance`

## Error Handling Pattern

Errors wrapped with context using `pkg/errors`:

```go
// pkg/errors/errors.go
var (
    NotFound      = errors.New("Not Found")
    Timedout      = errors.New("Timedout")
    InvalidInput  = errors.New("Invalid Input")
    // ...
)

// Usage
if err != nil {
    return errors.Wrapf(err, "Failed to get VM %s", vmName)
}

// Checking
if errors.IsNotFound(err) {
    // handle not found
}
```

## Resource Management Pattern

WMI instances must be closed to release COM resources:

```go
instance, err := GetInstance(...)
if err != nil {
    return err
}
defer instance.Close()  // Always defer Close()

// For collections
instances, err := EnumerateInstances(...)
defer instances.Close()
```

## Async/Job Pattern

Long-running WMI operations return job objects:

```go
result, err := instance.InvokeMethodWithReturn("RequestStateChange", params...)
// result contains job reference for polling
```

## Generated Code Patterns

Generated from WMI MOF with `wmigen`:

1. **Enum Constants**: Separate `*_<EnumName>.go` files
2. **Property Accessors**: Get/Set methods for each property
3. **Inheritance**: Go struct embedding mirrors WMI class hierarchy
4. **Namespace Organization**: Directory structure mirrors WMI namespaces
