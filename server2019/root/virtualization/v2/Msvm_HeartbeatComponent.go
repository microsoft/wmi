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

// Msvm_HeartbeatComponent struct
type Msvm_HeartbeatComponent struct {
	*CIM_LogicalDevice
}

func NewMsvm_HeartbeatComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_HeartbeatComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_HeartbeatComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_HeartbeatComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_HeartbeatComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_HeartbeatComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}
