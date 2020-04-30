// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_DirectorySpecification struct
type Win32_DirectorySpecification struct {
	*CIM_DirectorySpecification

	//
	DefaultDir string

	//
	Directory string
}

func NewWin32_DirectorySpecificationEx1(instance *cim.WmiInstance) (newInstance *Win32_DirectorySpecification, err error) {
	tmp, err := NewCIM_DirectorySpecificationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_DirectorySpecification{
		CIM_DirectorySpecification: tmp,
	}
	return
}

func NewWin32_DirectorySpecificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_DirectorySpecification, err error) {
	tmp, err := NewCIM_DirectorySpecificationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_DirectorySpecification{
		CIM_DirectorySpecification: tmp,
	}
	return
}

// SetDefaultDir sets the value of DefaultDir for the instance
func (instance *Win32_DirectorySpecification) SetPropertyDefaultDir(value string) (err error) {
	return instance.SetProperty("DefaultDir", value)
}

// GetDefaultDir gets the value of DefaultDir for the instance
func (instance *Win32_DirectorySpecification) GetPropertyDefaultDir() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultDir")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectory sets the value of Directory for the instance
func (instance *Win32_DirectorySpecification) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", value)
}

// GetDirectory gets the value of Directory for the instance
func (instance *Win32_DirectorySpecification) GetPropertyDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("Directory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}