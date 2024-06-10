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

// Msvm_DiskMergeSettingData struct
type Msvm_DiskMergeSettingData struct {
	*CIM_SettingData

	//
	EnabledState uint32
}

func NewMsvm_DiskMergeSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_DiskMergeSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_DiskMergeSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_DiskMergeSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_DiskMergeSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_DiskMergeSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_DiskMergeSettingData) SetPropertyEnabledState(value uint32) (err error) {
	return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_DiskMergeSettingData) GetPropertyEnabledState() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
