// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package constant

import "time"

type WMINamespace string

const (
	Virtualization  WMINamespace = "root/virtualization/v2"
	CimV2           WMINamespace = "root/cimv2"
	StadardCimV2    WMINamespace = "root/standardcimv2"
	FailoverCluster WMINamespace = "root/mscluster"
)

const (
	HostName string = "localhost"
)

// Constants for WMI error codes

const (
	Success      int32 = 0
	AsyncJob     int32 = 4096
	SystemInUse  int32 = 32775
	InvalidState int32 = 32773
	Generic      int32 = 32768
)

// Constants for WMI OLE error codes
const (
	WBEM_E_NOT_FOUND = 0x8004100C
)

// Constants for WMI method retries
const (
	WmiMethodMaxRetries = 5
	WmiMethodRetryDelay = 500 * time.Millisecond
)
