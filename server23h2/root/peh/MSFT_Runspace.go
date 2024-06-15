// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_Runspace struct
type MSFT_Runspace struct {
	*cim.WmiInstance

	//
	InstanceId string

	//
	Name string
}

func NewMSFT_RunspaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_Runspace, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_Runspace{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_RunspaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Runspace, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Runspace{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceId sets the value of InstanceId for the instance
func (instance *MSFT_Runspace) SetPropertyInstanceId(value string) (err error) {
	return instance.SetProperty("InstanceId", (value))
}

// GetInstanceId gets the value of InstanceId for the instance
func (instance *MSFT_Runspace) GetPropertyInstanceId() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_Runspace) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_Runspace) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

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
