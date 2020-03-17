// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_BIOSElement struct
type CIM_BIOSElement struct {
	CIM_SoftwareElement

	//
	PrimaryBIOS bool
}

// SetPrimaryBIOS sets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) SetPropertyPrimaryBIOS(value bool) (err error) {
	return instance.SetProperty("PrimaryBIOS", value)
}

// GetPrimaryBIOS gets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) GetPropertyPrimaryBIOS() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryBIOS")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
