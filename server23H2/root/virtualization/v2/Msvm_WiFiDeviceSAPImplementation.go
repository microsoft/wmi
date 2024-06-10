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

// Msvm_WiFiDeviceSAPImplementation struct
type Msvm_WiFiDeviceSAPImplementation struct {
	*CIM_DeviceSAPImplementation
}

func NewMsvm_WiFiDeviceSAPImplementationEx1(instance *cim.WmiInstance) (newInstance *Msvm_WiFiDeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiDeviceSAPImplementation{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}

func NewMsvm_WiFiDeviceSAPImplementationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_WiFiDeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiDeviceSAPImplementation{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}
