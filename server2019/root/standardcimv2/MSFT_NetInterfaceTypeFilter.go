// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetInterfaceTypeFilter struct
type MSFT_NetInterfaceTypeFilter struct {
	CIM_FilterEntryBase

	//
	InterfaceType uint32
}

// SetInterfaceType sets the value of InterfaceType for the instance
func (instance *MSFT_NetInterfaceTypeFilter) SetPropertyInterfaceType(value uint32) (err error) {
	return instance.SetProperty("InterfaceType", value)
}

// GetInterfaceType gets the value of InterfaceType for the instance
func (instance *MSFT_NetInterfaceTypeFilter) GetPropertyInterfaceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
