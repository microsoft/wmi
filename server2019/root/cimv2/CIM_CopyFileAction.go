// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_CopyFileAction struct
type CIM_CopyFileAction struct {
	CIM_FileAction

	//
	DeleteAfterCopy bool

	//
	Destination string

	//
	Source string
}

// SetDeleteAfterCopy sets the value of DeleteAfterCopy for the instance
func (instance *CIM_CopyFileAction) SetPropertyDeleteAfterCopy(value bool) (err error) {
	return instance.SetProperty("DeleteAfterCopy", value)
}

// GetDeleteAfterCopy gets the value of DeleteAfterCopy for the instance
func (instance *CIM_CopyFileAction) GetPropertyDeleteAfterCopy() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteAfterCopy")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestination sets the value of Destination for the instance
func (instance *CIM_CopyFileAction) SetPropertyDestination(value string) (err error) {
	return instance.SetProperty("Destination", value)
}

// GetDestination gets the value of Destination for the instance
func (instance *CIM_CopyFileAction) GetPropertyDestination() (value string, err error) {
	retValue, err := instance.GetProperty("Destination")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSource sets the value of Source for the instance
func (instance *CIM_CopyFileAction) SetPropertySource(value string) (err error) {
	return instance.SetProperty("Source", value)
}

// GetSource gets the value of Source for the instance
func (instance *CIM_CopyFileAction) GetPropertySource() (value string, err error) {
	retValue, err := instance.GetProperty("Source")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
