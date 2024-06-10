// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SerialPort struct
type Msvm_SerialPort struct {
	*CIM_LogicalPort
}

func NewMsvm_SerialPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_SerialPort, err error) {
	tmp, err := NewCIM_LogicalPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPort{
		CIM_LogicalPort: tmp,
	}
	return
}

func NewMsvm_SerialPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SerialPort, err error) {
	tmp, err := NewCIM_LogicalPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPort{
		CIM_LogicalPort: tmp,
	}
	return
}

func (instance *Msvm_SerialPort) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_SerialPort) GetRelatedSerialController() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SerialController")
}

func (instance *Msvm_SerialPort) GetRelatedSerialPortSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SerialPortSettingData")
}

func (instance *Msvm_SerialPort) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
