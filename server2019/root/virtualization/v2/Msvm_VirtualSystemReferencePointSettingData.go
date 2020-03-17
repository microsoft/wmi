// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemReferencePointSettingData struct
type Msvm_VirtualSystemReferencePointSettingData struct {
	CIM_SettingData

	//
	ConsistencyLevel uint8
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePointSettingData) SetPropertyConsistencyLevel(value uint8) (err error) {
	return instance.SetProperty("ConsistencyLevel", value)
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePointSettingData) GetPropertyConsistencyLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
