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

// Msvm_SerialPortOnSerialController struct
type Msvm_SerialPortOnSerialController struct {
	*CIM_PortOnDevice
}

func NewMsvm_SerialPortOnSerialControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_SerialPortOnSerialController, err error) {
	tmp, err := NewCIM_PortOnDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortOnSerialController{
		CIM_PortOnDevice: tmp,
	}
	return
}

func NewMsvm_SerialPortOnSerialControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SerialPortOnSerialController, err error) {
	tmp, err := NewCIM_PortOnDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortOnSerialController{
		CIM_PortOnDevice: tmp,
	}
	return
}
