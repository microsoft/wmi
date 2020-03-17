// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SomFilterStatus struct
type MSFT_SomFilterStatus struct {
	cim.WmiInstance

	//
	ContainerAvailable bool

	//
	Domain string

	//
	SchemaAvailable bool
}

// SetContainerAvailable sets the value of ContainerAvailable for the instance
func (instance *MSFT_SomFilterStatus) SetPropertyContainerAvailable(value bool) (err error) {
	return instance.SetProperty("ContainerAvailable", value)
}

// GetContainerAvailable gets the value of ContainerAvailable for the instance
func (instance *MSFT_SomFilterStatus) GetPropertyContainerAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("ContainerAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *MSFT_SomFilterStatus) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *MSFT_SomFilterStatus) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSchemaAvailable sets the value of SchemaAvailable for the instance
func (instance *MSFT_SomFilterStatus) SetPropertySchemaAvailable(value bool) (err error) {
	return instance.SetProperty("SchemaAvailable", value)
}

// GetSchemaAvailable gets the value of SchemaAvailable for the instance
func (instance *MSFT_SomFilterStatus) GetPropertySchemaAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("SchemaAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
