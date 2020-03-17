// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTRegistryDword struct
type MSFT_MTRegistryDword struct {
	MSFT_MTRegistryValue

	//
	Data uint32
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryDword) SetPropertyData(value uint32) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryDword) GetPropertyData() (value uint32, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
