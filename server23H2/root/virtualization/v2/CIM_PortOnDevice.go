// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PortOnDevice struct
type CIM_PortOnDevice struct {
	*CIM_HostedDependency
}

func NewCIM_PortOnDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_PortOnDevice, err error) {
	tmp, err := NewCIM_HostedDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PortOnDevice{
		CIM_HostedDependency: tmp,
	}
	return
}

func NewCIM_PortOnDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PortOnDevice, err error) {
	tmp, err := NewCIM_HostedDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PortOnDevice{
		CIM_HostedDependency: tmp,
	}
	return
}
