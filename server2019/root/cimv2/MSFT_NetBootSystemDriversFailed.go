// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetBootSystemDriversFailed struct
type MSFT_NetBootSystemDriversFailed struct {
	MSFT_SCMEventLogEvent

	//
	DriverList string
}

// SetDriverList sets the value of DriverList for the instance
func (instance *MSFT_NetBootSystemDriversFailed) SetPropertyDriverList(value string) (err error) {
	return instance.SetProperty("DriverList", value)
}

// GetDriverList gets the value of DriverList for the instance
func (instance *MSFT_NetBootSystemDriversFailed) GetPropertyDriverList() (value string, err error) {
	retValue, err := instance.GetProperty("DriverList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
