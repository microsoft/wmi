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

// CIM_DependencyContext struct
type CIM_DependencyContext struct {
	cim.WmiInstance

	//
	Context CIM_Configuration

	//
	Dependency CIM_Dependency
}

// SetContext sets the value of Context for the instance
func (instance *CIM_DependencyContext) SetPropertyContext(value CIM_Configuration) (err error) {
	return instance.SetProperty("Context", value)
}

// GetContext gets the value of Context for the instance
func (instance *CIM_DependencyContext) GetPropertyContext() (value CIM_Configuration, err error) {
	retValue, err := instance.GetProperty("Context")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Configuration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependency sets the value of Dependency for the instance
func (instance *CIM_DependencyContext) SetPropertyDependency(value CIM_Dependency) (err error) {
	return instance.SetProperty("Dependency", value)
}

// GetDependency gets the value of Dependency for the instance
func (instance *CIM_DependencyContext) GetPropertyDependency() (value CIM_Dependency, err error) {
	retValue, err := instance.GetProperty("Dependency")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Dependency)
	if !ok {
		// TODO: Set an error
	}
	return
}
