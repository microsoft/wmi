// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetInterfaceFilter struct
type MSFT_NetInterfaceFilter struct {
	CIM_FilterEntryBase

	//
	InterfaceAlias []string
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetInterfaceFilter) SetPropertyInterfaceAlias(value []string) (err error) {
	return instance.SetProperty("InterfaceAlias", value)
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetInterfaceFilter) GetPropertyInterfaceAlias() (value []string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
