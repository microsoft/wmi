// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Battery struct
type Win32_Battery struct {
	CIM_Battery

	//
	BatteryRechargeTime uint32

	//
	ExpectedBatteryLife uint32
}

// SetBatteryRechargeTime sets the value of BatteryRechargeTime for the instance
func (instance *Win32_Battery) SetPropertyBatteryRechargeTime(value uint32) (err error) {
	return instance.SetProperty("BatteryRechargeTime", value)
}

// GetBatteryRechargeTime gets the value of BatteryRechargeTime for the instance
func (instance *Win32_Battery) GetPropertyBatteryRechargeTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("BatteryRechargeTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpectedBatteryLife sets the value of ExpectedBatteryLife for the instance
func (instance *Win32_Battery) SetPropertyExpectedBatteryLife(value uint32) (err error) {
	return instance.SetProperty("ExpectedBatteryLife", value)
}

// GetExpectedBatteryLife gets the value of ExpectedBatteryLife for the instance
func (instance *Win32_Battery) GetPropertyExpectedBatteryLife() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExpectedBatteryLife")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
