// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_InstalledOS struct
type CIM_InstalledOS struct {
	CIM_SystemComponent

	//
	PrimaryOS bool
}

// SetPrimaryOS sets the value of PrimaryOS for the instance
func (instance *CIM_InstalledOS) SetPropertyPrimaryOS(value bool) (err error) {
	return instance.SetProperty("PrimaryOS", value)
}

// GetPrimaryOS gets the value of PrimaryOS for the instance
func (instance *CIM_InstalledOS) GetPropertyPrimaryOS() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryOS")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
