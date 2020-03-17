// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Container struct
type CIM_Container struct {
	CIM_Component

	//
	LocationWithinContainer string
}

// SetLocationWithinContainer sets the value of LocationWithinContainer for the instance
func (instance *CIM_Container) SetPropertyLocationWithinContainer(value string) (err error) {
	return instance.SetProperty("LocationWithinContainer", value)
}

// GetLocationWithinContainer gets the value of LocationWithinContainer for the instance
func (instance *CIM_Container) GetPropertyLocationWithinContainer() (value string, err error) {
	retValue, err := instance.GetProperty("LocationWithinContainer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
