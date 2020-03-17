// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerPlan struct
type Win32_PowerPlan struct {
	CIM_SettingData

	//
	IsActive bool
}

// SetIsActive sets the value of IsActive for the instance
func (instance *Win32_PowerPlan) SetPropertyIsActive(value bool) (err error) {
	return instance.SetProperty("IsActive", value)
}

// GetIsActive gets the value of IsActive for the instance
func (instance *Win32_PowerPlan) GetPropertyIsActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IsActive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *Win32_PowerPlan) Activate() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Activate")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
