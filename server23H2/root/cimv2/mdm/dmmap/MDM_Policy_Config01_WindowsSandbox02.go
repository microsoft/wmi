// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_WindowsSandbox02 struct
type MDM_Policy_Config01_WindowsSandbox02 struct {
	*cim.WmiInstance

	//
	AllowAudioInput int32

	//
	AllowClipboardRedirection int32

	//
	AllowNetworking int32

	//
	AllowPrinterRedirection int32

	//
	AllowVGPU int32

	//
	AllowVideoInput int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_WindowsSandbox02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_WindowsSandbox02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsSandbox02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_WindowsSandbox02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_WindowsSandbox02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsSandbox02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAudioInput sets the value of AllowAudioInput for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowAudioInput(value int32) (err error) {
	return instance.SetProperty("AllowAudioInput", (value))
}

// GetAllowAudioInput gets the value of AllowAudioInput for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowAudioInput() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAudioInput")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowClipboardRedirection sets the value of AllowClipboardRedirection for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowClipboardRedirection(value int32) (err error) {
	return instance.SetProperty("AllowClipboardRedirection", (value))
}

// GetAllowClipboardRedirection gets the value of AllowClipboardRedirection for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowClipboardRedirection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowClipboardRedirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowNetworking sets the value of AllowNetworking for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowNetworking(value int32) (err error) {
	return instance.SetProperty("AllowNetworking", (value))
}

// GetAllowNetworking gets the value of AllowNetworking for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowNetworking() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowNetworking")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowPrinterRedirection sets the value of AllowPrinterRedirection for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowPrinterRedirection(value int32) (err error) {
	return instance.SetProperty("AllowPrinterRedirection", (value))
}

// GetAllowPrinterRedirection gets the value of AllowPrinterRedirection for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowPrinterRedirection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPrinterRedirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowVGPU sets the value of AllowVGPU for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowVGPU(value int32) (err error) {
	return instance.SetProperty("AllowVGPU", (value))
}

// GetAllowVGPU gets the value of AllowVGPU for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowVGPU() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVGPU")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowVideoInput sets the value of AllowVideoInput for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyAllowVideoInput(value int32) (err error) {
	return instance.SetProperty("AllowVideoInput", (value))
}

// GetAllowVideoInput gets the value of AllowVideoInput for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyAllowVideoInput() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVideoInput")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WindowsSandbox02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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
