// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_DriverIbSettingData struct
type MLNX_DriverIbSettingData struct {
	*MLNX_DriverSettingData

	//
	DebugFlags uint32

	//
	IbalDebugFlags uint32

	//
	IbalDebugLevel uint32
}

func NewMLNX_DriverIbSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverIbSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIbSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

func NewMLNX_DriverIbSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverIbSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIbSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

// SetDebugFlags sets the value of DebugFlags for the instance
func (instance *MLNX_DriverIbSettingData) SetPropertyDebugFlags(value uint32) (err error) {
	return instance.SetProperty("DebugFlags", (value))
}

// GetDebugFlags gets the value of DebugFlags for the instance
func (instance *MLNX_DriverIbSettingData) GetPropertyDebugFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIbalDebugFlags sets the value of IbalDebugFlags for the instance
func (instance *MLNX_DriverIbSettingData) SetPropertyIbalDebugFlags(value uint32) (err error) {
	return instance.SetProperty("IbalDebugFlags", (value))
}

// GetIbalDebugFlags gets the value of IbalDebugFlags for the instance
func (instance *MLNX_DriverIbSettingData) GetPropertyIbalDebugFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIbalDebugLevel sets the value of IbalDebugLevel for the instance
func (instance *MLNX_DriverIbSettingData) SetPropertyIbalDebugLevel(value uint32) (err error) {
	return instance.SetProperty("IbalDebugLevel", (value))
}

// GetIbalDebugLevel gets the value of IbalDebugLevel for the instance
func (instance *MLNX_DriverIbSettingData) GetPropertyIbalDebugLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="DebugFlags" type="uint32 "></param>
// <param name="IbalDebugFlags" type="uint32 "></param>
// <param name="IbalDebugLevel" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverIbSettingData) SetValue( /* IN */ DebugFlags uint32,
	/* IN */ IbalDebugFlags uint32,
	/* IN */ IbalDebugLevel uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", DebugFlags, IbalDebugFlags, IbalDebugLevel)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverIbSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
