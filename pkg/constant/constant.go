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

const (
	// Timeout for WMI operations
	WmiVmCreateTimeout = 60 * time.Minute
	WmiVmDeleteTimeout = 60 * time.Minute
)
