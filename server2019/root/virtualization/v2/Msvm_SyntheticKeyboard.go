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

// Msvm_SyntheticKeyboard struct
type Msvm_SyntheticKeyboard struct {
	*CIM_UserDevice
}

func NewMsvm_SyntheticKeyboardEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticKeyboard, err error) {
	tmp, err := NewCIM_UserDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticKeyboard{
		CIM_UserDevice: tmp,
	}
	return
}

func NewMsvm_SyntheticKeyboardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticKeyboard, err error) {
	tmp, err := NewCIM_UserDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticKeyboard{
		CIM_UserDevice: tmp,
	}
	return
}

func (instance *Msvm_SyntheticKeyboard) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_SyntheticKeyboard) GetRelatedResourceAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourceAllocationSettingData")
}

func (instance *Msvm_SyntheticKeyboard) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
