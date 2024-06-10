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

// MLNX_PCIDeviceSriovSettingData struct
type MLNX_PCIDeviceSriovSettingData struct {
	*MLNX_PCIDeviceSettingData

	//
	SriovPort1NumVFs uint32

	//
	SriovPort2NumVFs uint32

	//
	SriovPortMode uint32
}

func NewMLNX_PCIDeviceSriovSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_PCIDeviceSriovSettingData, err error) {
	tmp, err := NewMLNX_PCIDeviceSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceSriovSettingData{
		MLNX_PCIDeviceSettingData: tmp,
	}
	return
}

func NewMLNX_PCIDeviceSriovSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_PCIDeviceSriovSettingData, err error) {
	tmp, err := NewMLNX_PCIDeviceSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceSriovSettingData{
		MLNX_PCIDeviceSettingData: tmp,
	}
	return
}

// SetSriovPort1NumVFs sets the value of SriovPort1NumVFs for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) SetPropertySriovPort1NumVFs(value uint32) (err error) {
	return instance.SetProperty("SriovPort1NumVFs", (value))
}

// GetSriovPort1NumVFs gets the value of SriovPort1NumVFs for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) GetPropertySriovPort1NumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPort1NumVFs")
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

// SetSriovPort2NumVFs sets the value of SriovPort2NumVFs for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) SetPropertySriovPort2NumVFs(value uint32) (err error) {
	return instance.SetProperty("SriovPort2NumVFs", (value))
}

// GetSriovPort2NumVFs gets the value of SriovPort2NumVFs for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) GetPropertySriovPort2NumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPort2NumVFs")
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

// SetSriovPortMode sets the value of SriovPortMode for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) SetPropertySriovPortMode(value uint32) (err error) {
	return instance.SetProperty("SriovPortMode", (value))
}

// GetSriovPortMode gets the value of SriovPortMode for the instance
func (instance *MLNX_PCIDeviceSriovSettingData) GetPropertySriovPortMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPortMode")
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

// <param name="SriovPort1NumVFs" type="uint32 "></param>
// <param name="SriovPort2NumVFs" type="uint32 "></param>
// <param name="SriovPortMode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDeviceSriovSettingData) SetValue( /* IN */ SriovPortMode uint32,
	/* IN */ SriovPort1NumVFs uint32,
	/* IN */ SriovPort2NumVFs uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", SriovPortMode, SriovPort1NumVFs, SriovPort2NumVFs)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDeviceSriovSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
