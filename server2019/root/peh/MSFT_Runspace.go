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

// MSFT_Runspace struct
type MSFT_Runspace struct {
	cim.WmiInstance

	//
	InstanceId string

	//
	Name string
}

// SetInstanceId sets the value of InstanceId for the instance
func (instance *MSFT_Runspace) SetPropertyInstanceId(value string) (err error) {
	return instance.SetProperty("InstanceId", value)
}

// GetInstanceId gets the value of InstanceId for the instance
func (instance *MSFT_Runspace) GetPropertyInstanceId() (value string, err error) {
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
func (instance *MSFT_Runspace) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_Runspace) GetPropertyName() (value string, err error) {
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

// <param name="args" type="interface{} "></param>
// <param name="pipelineDefinitionInstanceId" type="string "></param>
// <param name="pipelineName" type="string "></param>

// <param name="pipeline" type="MSFT_Pipeline "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Runspace) CreatePipeline( /* IN */ pipelineDefinitionInstanceId string,
	/* IN */ pipelineName string,
	/* IN */ args interface{},
	/* OUT */ pipeline MSFT_Pipeline) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreatePipeline", pipelineDefinitionInstanceId, pipelineName, args)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="modules" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Runspace) ImportModule( /* IN */ modules []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ImportModule", modules)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
