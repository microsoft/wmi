// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_RedundancyGroup struct
type CIM_RedundancyGroup struct {
	CIM_LogicalElement

	//
	CreationClassName string

	//
	RedundancyStatus uint16
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_RedundancyGroup) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_RedundancyGroup) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRedundancyStatus sets the value of RedundancyStatus for the instance
func (instance *CIM_RedundancyGroup) SetPropertyRedundancyStatus(value uint16) (err error) {
	return instance.SetProperty("RedundancyStatus", value)
}

// GetRedundancyStatus gets the value of RedundancyStatus for the instance
func (instance *CIM_RedundancyGroup) GetPropertyRedundancyStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("RedundancyStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
