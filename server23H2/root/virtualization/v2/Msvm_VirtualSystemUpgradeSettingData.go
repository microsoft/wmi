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

// Msvm_VirtualSystemUpgradeSettingData struct
type Msvm_VirtualSystemUpgradeSettingData struct {
	*CIM_SettingData

	//
	TargetVersion string
}

func NewMsvm_VirtualSystemUpgradeSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemUpgradeSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemUpgradeSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemUpgradeSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemUpgradeSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemUpgradeSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetTargetVersion sets the value of TargetVersion for the instance
func (instance *Msvm_VirtualSystemUpgradeSettingData) SetPropertyTargetVersion(value string) (err error) {
	return instance.SetProperty("TargetVersion", (value))
}

// GetTargetVersion gets the value of TargetVersion for the instance
func (instance *Msvm_VirtualSystemUpgradeSettingData) GetPropertyTargetVersion() (value string, err error) {
	retValue, err := instance.GetProperty("TargetVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
