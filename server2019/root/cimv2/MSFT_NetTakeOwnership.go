// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetTakeOwnership struct
type MSFT_NetTakeOwnership struct {
	MSFT_SCMEventLogEvent

	//
	RegistryKey string
}

// SetRegistryKey sets the value of RegistryKey for the instance
func (instance *MSFT_NetTakeOwnership) SetPropertyRegistryKey(value string) (err error) {
	return instance.SetProperty("RegistryKey", value)
}

// GetRegistryKey gets the value of RegistryKey for the instance
func (instance *MSFT_NetTakeOwnership) GetPropertyRegistryKey() (value string, err error) {
	retValue, err := instance.GetProperty("RegistryKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
