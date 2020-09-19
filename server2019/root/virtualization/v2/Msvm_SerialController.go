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

// Msvm_SerialController struct
type Msvm_SerialController struct {
	*CIM_SerialController
}

func NewMsvm_SerialControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_SerialController, err error) {
	tmp, err := NewCIM_SerialControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialController{
		CIM_SerialController: tmp,
	}
	return
}

func NewMsvm_SerialControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SerialController, err error) {
	tmp, err := NewCIM_SerialControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialController{
		CIM_SerialController: tmp,
	}
	return
}

func (instance *Msvm_SerialController) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_SerialController) GetRelatedSerialPort() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_SerialPort")
}

func (instance *Msvm_SerialController) GetRelatedResourceAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourceAllocationSettingData")
}
