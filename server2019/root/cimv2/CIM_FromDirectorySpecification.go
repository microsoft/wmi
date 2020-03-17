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

// CIM_FromDirectorySpecification struct
type CIM_FromDirectorySpecification struct {
	cim.WmiInstance

	//
	FileName CIM_FileAction

	//
	SourceDirectory CIM_DirectorySpecification
}

// SetFileName sets the value of FileName for the instance
func (instance *CIM_FromDirectorySpecification) SetPropertyFileName(value CIM_FileAction) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *CIM_FromDirectorySpecification) GetPropertyFileName() (value CIM_FileAction, err error) {
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
func (instance *CIM_FromDirectorySpecification) SetPropertySourceDirectory(value CIM_DirectorySpecification) (err error) {
	return instance.SetProperty("SourceDirectory", value)
}

// GetSourceDirectory gets the value of SourceDirectory for the instance
func (instance *CIM_FromDirectorySpecification) GetPropertySourceDirectory() (value CIM_DirectorySpecification, err error) {
	retValue, err := instance.GetProperty("SourceDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DirectorySpecification)
	if !ok {
		// TODO: Set an error
	}
	return
}
