// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSNetworkAdapterListSetting struct
type Win32_TSNetworkAdapterListSetting struct {
	Win32_TerminalSetting

	//
	NetworkAdapterID string

	//
	NetworkAdapterIP string
}

// SetNetworkAdapterID sets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterListSetting) SetPropertyNetworkAdapterID(value string) (err error) {
	return instance.SetProperty("NetworkAdapterID", value)
}

// GetNetworkAdapterID gets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterListSetting) GetPropertyNetworkAdapterID() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterIP sets the value of NetworkAdapterIP for the instance
func (instance *Win32_TSNetworkAdapterListSetting) SetPropertyNetworkAdapterIP(value string) (err error) {
	return instance.SetProperty("NetworkAdapterIP", value)
}

// GetNetworkAdapterIP gets the value of NetworkAdapterIP for the instance
func (instance *Win32_TSNetworkAdapterListSetting) GetPropertyNetworkAdapterIP() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
