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

// MSFT_Engine struct
type MSFT_Engine struct {
	cim.WmiInstance

	//
	InstanceId string

	//
	Name string
}

// SetInstanceId sets the value of InstanceId for the instance
func (instance *MSFT_Engine) SetPropertyInstanceId(value string) (err error) {
	return instance.SetProperty("InstanceId", value)
}

// GetInstanceId gets the value of InstanceId for the instance
func (instance *MSFT_Engine) GetPropertyInstanceId() (value string, err error) {
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
func (instance *MSFT_Engine) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_Engine) GetPropertyName() (value string, err error) {
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

// <param name="runspaceName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="runspace" type="MSFT_Runspace "></param>
func (instance *MSFT_Engine) CreateDefaultRunspace( /* IN */ runspaceName string,
	/* OUT */ runspace MSFT_Runspace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateDefaultRunspace", runspaceName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
