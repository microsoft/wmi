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

// CIM_FromDirectoryAction struct
type CIM_FromDirectoryAction struct {
	cim.WmiInstance

	//
	FileName CIM_FileAction

	//
	SourceDirectory CIM_DirectoryAction
}

// SetFileName sets the value of FileName for the instance
func (instance *CIM_FromDirectoryAction) SetPropertyFileName(value CIM_FileAction) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *CIM_FromDirectoryAction) GetPropertyFileName() (value CIM_FileAction, err error) {
	retValue, err := instance.GetProperty("FileName")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_FileAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceDirectory sets the value of SourceDirectory for the instance
func (instance *CIM_FromDirectoryAction) SetPropertySourceDirectory(value CIM_DirectoryAction) (err error) {
	return instance.SetProperty("SourceDirectory", value)
}

// GetSourceDirectory gets the value of SourceDirectory for the instance
func (instance *CIM_FromDirectoryAction) GetPropertySourceDirectory() (value CIM_DirectoryAction, err error) {
	retValue, err := instance.GetProperty("SourceDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DirectoryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}
