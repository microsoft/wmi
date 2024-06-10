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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VssComponentSettingData struct
type Msvm_VssComponentSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	EnabledState uint16
}

func NewMsvm_VssComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VssComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_VssComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VssComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_VssComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_VssComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
