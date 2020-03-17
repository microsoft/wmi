// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_DiskMergeSettingData struct
type Msvm_DiskMergeSettingData struct {
	CIM_SettingData

	//
	EnabledState uint32
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_DiskMergeSettingData) SetPropertyEnabledState(value uint32) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_DiskMergeSettingData) GetPropertyEnabledState() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
