// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_JobSettingData struct
type CIM_JobSettingData struct {
	*CIM_SettingData

	//
	DeleteOnCompletion bool

	//
	OtherRecoveryAction string

	//
	RecoveryAction JobSettingData_RecoveryAction
}

func NewCIM_JobSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_JobSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_JobSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_JobSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_JobSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_JobSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetDeleteOnCompletion sets the value of DeleteOnCompletion for the instance
func (instance *CIM_JobSettingData) SetPropertyDeleteOnCompletion(value bool) (err error) {
	return instance.SetProperty("DeleteOnCompletion", (value))
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *CIM_JobSettingData) GetPropertyDeleteOnCompletion() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteOnCompletion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *CIM_JobSettingData) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", (value))
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *CIM_JobSettingData) GetPropertyOtherRecoveryAction() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecoveryAction")
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

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *CIM_JobSettingData) SetPropertyRecoveryAction(value JobSettingData_RecoveryAction) (err error) {
	return instance.SetProperty("RecoveryAction", (value))
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *CIM_JobSettingData) GetPropertyRecoveryAction() (value JobSettingData_RecoveryAction, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = JobSettingData_RecoveryAction(valuetmp)

	return
}
