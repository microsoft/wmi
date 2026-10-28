// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// LENOVO_GAMEZONE_DATA struct
type LENOVO_GAMEZONE_DATA struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewLENOVO_GAMEZONE_DATAEx1(instance *cim.WmiInstance) (newInstance *LENOVO_GAMEZONE_DATA, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &LENOVO_GAMEZONE_DATA{
		WmiInstance: tmp,
	}
	return
}

func NewLENOVO_GAMEZONE_DATAEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LENOVO_GAMEZONE_DATA, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LENOVO_GAMEZONE_DATA{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *LENOVO_GAMEZONE_DATA) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *LENOVO_GAMEZONE_DATA) GetPropertyActive() (value bool, err error) {
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
func (instance *LENOVO_GAMEZONE_DATA) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *LENOVO_GAMEZONE_DATA) GetPropertyInstanceName() (value string, err error) {
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

// Get AslCode Version

// <param name="Data" type="uint32 ">AslCode version</param>
func (instance *LENOVO_GAMEZONE_DATA) GetVersion( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("GetVersion")
	if err != nil {
		return
	}
	return

}

// Get Product information

// <param name="Data" type="uint32 ">Get Product information</param>
func (instance *LENOVO_GAMEZONE_DATA) GetProductInfo( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("GetProductInfo")
	if err != nil {
		return
	}
	return

}

// Smart Fan feature

// <param name="Data" type="uint32 ">Smart Fan feature</param>
func (instance *LENOVO_GAMEZONE_DATA) IsSupportSmartFan( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("IsSupportSmartFan")
	if err != nil {
		return
	}
	return

}

// Get Support Thermal Mode

// <param name="Data" type="uint32 ">Get Support Thermal Mode</param>
func (instance *LENOVO_GAMEZONE_DATA) GetSupportThermalMode( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("GetSupportThermalMode")
	if err != nil {
		return
	}
	return

}

// Set Smart Fan Mode

// <param name="Data" type="uint32 ">Set Smart Fan Mode</param>
func (instance *LENOVO_GAMEZONE_DATA) SetSmartFanMode( /* IN */ Data uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetSmartFanMode", Data)
	if err != nil {
		return
	}
	return

}

// Get Smart Fan Mode

// <param name="Data" type="uint32 ">Get Smart Fan Mode</param>
func (instance *LENOVO_GAMEZONE_DATA) GetSmartFanMode( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("GetSmartFanMode")
	if err != nil {
		return
	}
	return

}

// Cooler Lighting feature

// <param name="Data" type="uint32 ">Cooler Lighting feature</param>
func (instance *LENOVO_GAMEZONE_DATA) IsSupportCoolerLighting( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("IsSupportCoolerLighting")
	if err != nil {
		return
	}
	return

}

// Get CoolerLighting status

// <param name="Data" type="uint32 ">Get CoolerLighting status</param>
func (instance *LENOVO_GAMEZONE_DATA) GetCoolerLighting( /* OUT */ Data uint32) (err error) {
	_, err = instance.InvokeMethod("GetCoolerLighting")
	if err != nil {
		return
	}
	return

}

// Set CoolerLighting status

// <param name="Data" type="uint32 ">Set CoolerLighting status</param>
func (instance *LENOVO_GAMEZONE_DATA) SetCoolerLighting( /* IN */ Data uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetCoolerLighting", Data)
	if err != nil {
		return
	}
	return

}
