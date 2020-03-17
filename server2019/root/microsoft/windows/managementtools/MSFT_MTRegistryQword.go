// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTRegistryQword struct
type MSFT_MTRegistryQword struct {
	MSFT_MTRegistryValue

	//
	Data uint64
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryQword) SetPropertyData(value uint64) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryQword) GetPropertyData() (value uint64, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
