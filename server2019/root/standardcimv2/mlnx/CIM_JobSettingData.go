// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_JobSettingData struct
type CIM_JobSettingData struct {
	CIM_SettingData

	//
	DeleteOnCompletion bool

	//
	OtherRecoveryAction string

	//
	RecoveryAction JobSettingData_RecoveryAction
}

// SetDeleteOnCompletion sets the value of DeleteOnCompletion for the instance
func (instance *CIM_JobSettingData) SetPropertyDeleteOnCompletion(value bool) (err error) {
	return instance.SetProperty("DeleteOnCompletion", value)
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *CIM_JobSettingData) GetPropertyDeleteOnCompletion() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteOnCompletion")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *CIM_JobSettingData) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", value)
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *CIM_JobSettingData) GetPropertyOtherRecoveryAction() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *CIM_JobSettingData) SetPropertyRecoveryAction(value JobSettingData_RecoveryAction) (err error) {
	return instance.SetProperty("RecoveryAction", value)
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *CIM_JobSettingData) GetPropertyRecoveryAction() (value JobSettingData_RecoveryAction, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(JobSettingData_RecoveryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}
