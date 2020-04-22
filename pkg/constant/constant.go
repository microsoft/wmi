// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package constant

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
