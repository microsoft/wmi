// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTRegistryMultiString struct
type MSFT_MTRegistryMultiString struct {
	MSFT_MTRegistryValue

	//
	Data []string
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryMultiString) SetPropertyData(value []string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryMultiString) GetPropertyData() (value []string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
