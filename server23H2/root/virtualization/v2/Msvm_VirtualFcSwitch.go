// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualFcSwitch struct
type Msvm_VirtualFcSwitch struct {
	*CIM_ComputerSystem
}

func NewMsvm_VirtualFcSwitchEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualFcSwitch, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualFcSwitch{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMsvm_VirtualFcSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualFcSwitch, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualFcSwitch{
		CIM_ComputerSystem: tmp,
	}
	return
}
