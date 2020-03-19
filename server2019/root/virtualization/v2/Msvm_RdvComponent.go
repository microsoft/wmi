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

// Msvm_RdvComponent struct
type Msvm_RdvComponent struct {
	*CIM_LogicalDevice
}

func NewMsvm_RdvComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_RdvComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_RdvComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_RdvComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_RdvComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_RdvComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_RdvComponent) EnableEndPoints() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableEndPoints")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

func (instance *Msvm_RdvComponent) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_RdvComponent) GetRelatedRdvComponentSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_RdvComponentSettingData")
}
