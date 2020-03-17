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

// CIM_ToDirectorySpecification struct
type CIM_ToDirectorySpecification struct {
	cim.WmiInstance

	//
	DestinationDirectory CIM_DirectorySpecification

	//
	FileName CIM_CopyFileAction
}

// SetDestinationDirectory sets the value of DestinationDirectory for the instance
func (instance *CIM_ToDirectorySpecification) SetPropertyDestinationDirectory(value CIM_DirectorySpecification) (err error) {
	return instance.SetProperty("DestinationDirectory", value)
}

// GetDestinationDirectory gets the value of DestinationDirectory for the instance
func (instance *CIM_ToDirectorySpecification) GetPropertyDestinationDirectory() (value CIM_DirectorySpecification, err error) {
	retValue, err := instance.GetProperty("DestinationDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DirectorySpecification)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *CIM_ToDirectorySpecification) SetPropertyFileName(value CIM_CopyFileAction) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *CIM_ToDirectorySpecification) GetPropertyFileName() (value CIM_CopyFileAction, err error) {
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
