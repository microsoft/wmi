// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ToDirectoryAction struct
type CIM_ToDirectoryAction struct {
	cim.WmiInstance

	//
	DestinationDirectory CIM_DirectoryAction

	//
	FileName CIM_CopyFileAction
}

// SetDestinationDirectory sets the value of DestinationDirectory for the instance
func (instance *CIM_ToDirectoryAction) SetPropertyDestinationDirectory(value CIM_DirectoryAction) (err error) {
	return instance.SetProperty("DestinationDirectory", value)
}

// GetDestinationDirectory gets the value of DestinationDirectory for the instance
func (instance *CIM_ToDirectoryAction) GetPropertyDestinationDirectory() (value CIM_DirectoryAction, err error) {
	retValue, err := instance.GetProperty("DestinationDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DirectoryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *CIM_ToDirectoryAction) SetPropertyFileName(value CIM_CopyFileAction) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *CIM_ToDirectoryAction) GetPropertyFileName() (value CIM_CopyFileAction, err error) {
	retValue, err := instance.GetProperty("FileName")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_CopyFileAction)
	if !ok {
		// TODO: Set an error
	}
	return
}
