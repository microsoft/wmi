// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VssComponent struct
type Msvm_VssComponent struct {
	*CIM_LogicalDevice
}

func NewMsvm_VssComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_VssComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_VssComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VssComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}
