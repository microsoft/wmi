// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PipelineDefinition struct
type MSFT_PipelineDefinition struct {
	cim.WmiInstance

	//
	InstanceId string

	//
	Name string
}

// SetInstanceId sets the value of InstanceId for the instance
func (instance *MSFT_PipelineDefinition) SetPropertyInstanceId(value string) (err error) {
	return instance.SetProperty("InstanceId", value)
}

// GetInstanceId gets the value of InstanceId for the instance
func (instance *MSFT_PipelineDefinition) GetPropertyInstanceId() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_PipelineDefinition) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_PipelineDefinition) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
