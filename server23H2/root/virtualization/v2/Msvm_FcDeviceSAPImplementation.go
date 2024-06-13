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

// Msvm_FcDeviceSAPImplementation struct
type Msvm_FcDeviceSAPImplementation struct {
	*CIM_DeviceSAPImplementation
}

func NewMsvm_FcDeviceSAPImplementationEx1(instance *cim.WmiInstance) (newInstance *Msvm_FcDeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcDeviceSAPImplementation{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}

func NewMsvm_FcDeviceSAPImplementationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FcDeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcDeviceSAPImplementation{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}
