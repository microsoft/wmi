// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DCOMApplication struct
type Win32_DCOMApplication struct {
	Win32_COMApplication

	//
	AppID string
}

// SetAppID sets the value of AppID for the instance
func (instance *Win32_DCOMApplication) SetPropertyAppID(value string) (err error) {
	return instance.SetProperty("AppID", value)
}

// GetAppID gets the value of AppID for the instance
func (instance *Win32_DCOMApplication) GetPropertyAppID() (value string, err error) {
	retValue, err := instance.GetProperty("AppID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
