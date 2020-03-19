// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SystemDevice struct
type Msvm_SystemDevice struct {
	*CIM_SystemDevice
}

func NewMsvm_SystemDeviceEx1(instance *cim.WmiInstance) (newInstance *Msvm_SystemDevice, err error) {
	tmp, err := NewCIM_SystemDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemDevice{
		CIM_SystemDevice: tmp,
	}
	return
}

func NewMsvm_SystemDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SystemDevice, err error) {
	tmp, err := NewCIM_SystemDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemDevice{
		CIM_SystemDevice: tmp,
	}
	return
}
