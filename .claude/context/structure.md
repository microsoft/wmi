# Structure

> Last updated: 2026-01-11

## Directory Layout

```
wmi/
├── go/
│   └── wmi/                    # Interface definitions
│       ├── Class.go            # WMI class interface
│       ├── Credentials.go      # Credential handling
│       ├── Instance.go         # WMI instance interface
│       ├── InstanceManager.go  # Instance lifecycle
│       ├── MethodDeclaration.go
│       ├── MethodParameter.go
│       ├── MethodResult.go
│       ├── Property.go
│       ├── Qualifier.go
│       ├── Query.go
│       └── Session.go          # WMI session interface
│
├── pkg/
│   ├── base/                   # Core infrastructure
│   │   ├── credential/         # WMI credentials
│   │   ├── event/              # Event handling
│   │   ├── host/               # Remote host config
│   │   ├── instance/           # Instance management
│   │   ├── monitor/            # Event monitoring
│   │   ├── query/              # Query building
│   │   └── session/            # Session pooling
│   │
│   ├── cluster/                # Failover Clustering APIs
│   │   ├── compute/
│   │   │   ├── affinityrule/   # Placement rules
│   │   │   ├── cluster/        # Cluster management
│   │   │   ├── clusternode/    # Node operations
│   │   │   ├── groupset/       # Availability sets
│   │   │   ├── internal/       # Internal helpers
│   │   │   ├── resource/       # Cluster resources
│   │   │   └── resourcegroup/  # Resource groups
│   │   ├── constant/
│   │   ├── network/
│   │   │   └── clusternetwork/ # Cluster networking
│   │   ├── service/            # Cluster service
│   │   └── storage/
│   │       └── clustersharedvolume/
│   │
│   ├── constant/               # Global constants
│   │
│   ├── errors/                 # Error definitions
│   │
│   ├── hardware/               # Hardware APIs
│   │   ├── host/               # Computer system info
│   │   └── network/
│   │       ├── netadapter/
│   │       ├── netroute/
│   │       └── service/
│   │
│   ├── operatingsystem/        # OS APIs
│   │   └── core/
│   │       ├── process/
│   │       └── service/
│   │
│   ├── virtualization/         # Hyper-V APIs
│   │   ├── constant/
│   │   ├── core/
│   │   │   ├── gpu/            # GPU passthrough
│   │   │   ├── host/
│   │   │   ├── job/            # Async job handling
│   │   │   ├── memory/
│   │   │   ├── pcie/
│   │   │   ├── pnp/
│   │   │   ├── processor/
│   │   │   ├── resource/
│   │   │   ├── security/
│   │   │   ├── service/
│   │   │   ├── storage/
│   │   │   │   ├── controller/ # SCSI/IDE controllers
│   │   │   │   ├── disk/       # Virtual hard disks
│   │   │   │   ├── drive/      # DVD drives
│   │   │   │   └── service/
│   │   │   └── virtualsystem/  # VM management
│   │   └── network/            # Virtual networking
│   │
│   └── wmiinstance/            # WMI implementation
│       ├── WmiClass.go
│       ├── WmiInstance.go      # Core instance wrapper
│       ├── WmiMethod.go
│       ├── WmiProperty.go
│       ├── WmiSession.go
│       ├── WmiSessionManager.go
│       └── variant.go          # VARIANT handling
│
├── server2019/                 # Generated classes (Server 2019)
│   └── root/
│       ├── accesslogging/
│       ├── cimv2/              # Standard WMI classes
│       │   ├── mdm/
│       │   ├── power/
│       │   ├── security/
│       │   └── terminalservices/
│       ├── hypervcluster/
│       ├── microsoft/
│       │   └── windows/
│       │       ├── cluster/
│       │       ├── storage/
│       │       └── ...
│       ├── standardcimv2/
│       └── virtualization/
│           └── v2/             # Hyper-V classes (~800 files)
│
├── server23h2/                 # Generated classes (Server 23H2)
│   └── root/                   # Same structure as server2019
│
├── .github/                    # GitHub workflows
├── .gdn/                       # Guardian config
├── .pipelines/                 # Azure Pipelines
│
├── go.mod                      # Go module definition
├── go.sum
├── Makefile                    # Build commands
├── README.md
├── LICENSE                     # MIT License
├── SECURITY.md
└── azure-pipelines.yml
```

## Key Files

| File | Purpose |
|------|---------|
| `go.mod` | Module: `github.com/microsoft/wmi`, Go 1.24.3 |
| `Makefile` | Build: `make library`, `make test`, `make format` |
| `pkg/wmiinstance/WmiInstance.go` | Core WMI instance implementation |
| `pkg/wmiinstance/WmiSession.go` | Session connection/query logic |
| `pkg/base/session/session.go` | Session pool management |
| `pkg/errors/errors.go` | Standard error definitions |

## Module Relationships

```
Application
    │
    ├──► pkg/virtualization  ─┐
    ├──► pkg/cluster         ─┼──► server2019/ or server23h2/
    ├──► pkg/hardware        ─┤        │
    └──► pkg/operatingsystem─┘        │
                                       ▼
                               pkg/wmiinstance
                                       │
                                       ▼
                                  pkg/base/
                                (session, query, host)
                                       │
                                       ▼
                                   go-ole
```

## Generated Code Volume

| Directory | Go Files | Description |
|-----------|----------|-------------|
| server2019/ | ~9,945 | Server 2019 WMI classes |
| server23h2/ | ~12,367 | Server 23H2 WMI classes |
| pkg/ | ~150 | High-level APIs |
| go/wmi/ | 11 | Interface definitions |
