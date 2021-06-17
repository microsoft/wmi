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

type SecureBootTemplateGUID string

const (
	WindowsSecureBootTemplateGUID SecureBootTemplateGUID = "1734c6e8-3154-4dda-ba5f-a874cc483422"
	LinuxSecureBootTemplateGUID   SecureBootTemplateGUID = "272e7447-90a4-4563-a4b9-8e4ab00526ce"
	OSSSecureBootTemplateGUID     SecureBootTemplateGUID = "4292ae2b-ee2c-42b5-a969-dd8f8689f6f3"
)
