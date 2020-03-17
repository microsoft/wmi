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

// MSFT_Pipeline struct
type MSFT_Pipeline struct {
	cim.WmiInstance

	//
	InstanceId string

	//
	Name string
}

// SetInstanceId sets the value of InstanceId for the instance
func (instance *MSFT_Pipeline) SetPropertyInstanceId(value string) (err error) {
	return instance.SetProperty("InstanceId", value)
}

// GetInstanceId gets the value of InstanceId for the instance
func (instance *MSFT_Pipeline) GetPropertyInstanceId() (value string, err error) {
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
func (instance *MSFT_Pipeline) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_Pipeline) GetPropertyName() (value string, err error) {
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

//

// <param name="pipelineExecutionName" type="string "></param>

// <param name="pipelineExecution" type="MSFT_PipelineExecution "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Pipeline) Execute( /* IN */ pipelineExecutionName string,
	/* OUT */ pipelineExecution MSFT_PipelineExecution) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Execute", pipelineExecutionName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
