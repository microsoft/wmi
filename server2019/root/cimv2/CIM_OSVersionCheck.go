// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_OSVersionCheck struct
type CIM_OSVersionCheck struct {
	CIM_Check

	//
	MaximumVersion string

	//
	MinimumVersion string
}

// SetMaximumVersion sets the value of MaximumVersion for the instance
func (instance *CIM_OSVersionCheck) SetPropertyMaximumVersion(value string) (err error) {
	return instance.SetProperty("MaximumVersion", value)
}

// GetMaximumVersion gets the value of MaximumVersion for the instance
func (instance *CIM_OSVersionCheck) GetPropertyMaximumVersion() (value string, err error) {
	retValue, err := instance.GetProperty("MaximumVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumVersion sets the value of MinimumVersion for the instance
func (instance *CIM_OSVersionCheck) SetPropertyMinimumVersion(value string) (err error) {
	return instance.SetProperty("MinimumVersion", value)
}

// GetMinimumVersion gets the value of MinimumVersion for the instance
func (instance *CIM_OSVersionCheck) GetPropertyMinimumVersion() (value string, err error) {
	retValue, err := instance.GetProperty("MinimumVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
