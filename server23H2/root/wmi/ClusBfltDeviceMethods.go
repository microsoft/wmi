// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClusBfltDeviceMethods struct
type ClusBfltDeviceMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewClusBfltDeviceMethodsEx1(instance *cim.WmiInstance) (newInstance *ClusBfltDeviceMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltDeviceMethods{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltDeviceMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltDeviceMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltDeviceMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ClusBfltDeviceMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ClusBfltDeviceMethods) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ClusBfltDeviceMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ClusBfltDeviceMethods) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// Gets Device Attributes

// <param name="DeviceGuid" type="string ">Device Id</param>

// <param name="Attributes" type="uint32 ">Attributes</param>
func (instance *ClusBfltDeviceMethods) GetDeviceAttributes( /* IN */ DeviceGuid string,
	/* OUT */ Attributes uint32) (err error) {
	_, err = instance.InvokeMethod("GetDeviceAttributes", DeviceGuid)
	if err != nil {
		return
	}
	return

}

// Sets Device Attributes

// <param name="Attributes" type="uint32 ">Attributes</param>
// <param name="AttributesMask" type="uint32 ">AttributesMask</param>
// <param name="DeviceGuid" type="string ">Device Id</param>
func (instance *ClusBfltDeviceMethods) SetDeviceAttributes( /* IN */ DeviceGuid string,
	/* IN */ Attributes uint32,
	/* IN */ AttributesMask uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetDeviceAttributes", DeviceGuid, Attributes, AttributesMask)
	if err != nil {
		return
	}
	return

}

// Pauses Device IOs

// <param name="DeviceGuid" type="string ">Device Id</param>
// <param name="TimeMs" type="uint32 ">Time</param>
func (instance *ClusBfltDeviceMethods) PauseDeviceIOs( /* IN */ DeviceGuid string,
	/* IN */ TimeMs uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("PauseDeviceIOs", DeviceGuid, TimeMs)
	if err != nil {
		return
	}
	return

}

// Resumes Device IOs

// <param name="DeviceGuid" type="string ">Device Id</param>
func (instance *ClusBfltDeviceMethods) ResumeDeviceIOs( /* IN */ DeviceGuid string) (err error) {
	_, err = instance.InvokeMethodWithReturn("ResumeDeviceIOs", DeviceGuid)
	if err != nil {
		return
	}
	return

}

// Refresh Reg Params

// <param name="fRebootRequired" type="bool ">fRebootRequired</param>
func (instance *ClusBfltDeviceMethods) RefreshRegParams( /* OUT */ fRebootRequired bool) (err error) {
	_, err = instance.InvokeMethod("RefreshRegParams")
	if err != nil {
		return
	}
	return

}
