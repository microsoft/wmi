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

// CIM_StorageDefect struct
type CIM_StorageDefect struct {
	cim.WmiInstance

	//
	Error CIM_StorageError

	//
	Extent CIM_StorageExtent
}

// SetError sets the value of Error for the instance
func (instance *CIM_StorageDefect) SetPropertyError(value CIM_StorageError) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *CIM_StorageDefect) GetPropertyError() (value CIM_StorageError, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StorageError)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtent sets the value of Extent for the instance
func (instance *CIM_StorageDefect) SetPropertyExtent(value CIM_StorageExtent) (err error) {
	return instance.SetProperty("Extent", value)
}

// GetExtent gets the value of Extent for the instance
func (instance *CIM_StorageDefect) GetPropertyExtent() (value CIM_StorageExtent, err error) {
	retValue, err := instance.GetProperty("Extent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StorageExtent)
	if !ok {
		// TODO: Set an error
	}
	return
}
