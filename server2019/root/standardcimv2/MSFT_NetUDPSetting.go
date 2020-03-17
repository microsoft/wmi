// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetUDPSetting struct
type MSFT_NetUDPSetting struct {
	CIM_PolicyAction

	//
	DynamicPortRangeNumberOfPorts uint16

	//
	DynamicPortRangeStartPort uint16
}

// SetDynamicPortRangeNumberOfPorts sets the value of DynamicPortRangeNumberOfPorts for the instance
func (instance *MSFT_NetUDPSetting) SetPropertyDynamicPortRangeNumberOfPorts(value uint16) (err error) {
	return instance.SetProperty("DynamicPortRangeNumberOfPorts", value)
}

// GetDynamicPortRangeNumberOfPorts gets the value of DynamicPortRangeNumberOfPorts for the instance
func (instance *MSFT_NetUDPSetting) GetPropertyDynamicPortRangeNumberOfPorts() (value uint16, err error) {
	retValue, err := instance.GetProperty("DynamicPortRangeNumberOfPorts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDynamicPortRangeStartPort sets the value of DynamicPortRangeStartPort for the instance
func (instance *MSFT_NetUDPSetting) SetPropertyDynamicPortRangeStartPort(value uint16) (err error) {
	return instance.SetProperty("DynamicPortRangeStartPort", value)
}

// GetDynamicPortRangeStartPort gets the value of DynamicPortRangeStartPort for the instance
func (instance *MSFT_NetUDPSetting) GetPropertyDynamicPortRangeStartPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("DynamicPortRangeStartPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
