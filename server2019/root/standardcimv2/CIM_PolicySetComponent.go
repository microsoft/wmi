// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_PolicySetComponent struct
type CIM_PolicySetComponent struct {
	CIM_PolicyComponent

	//
	Priority uint16
}

// SetPriority sets the value of Priority for the instance
func (instance *CIM_PolicySetComponent) SetPropertyPriority(value uint16) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_PolicySetComponent) GetPropertyPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
