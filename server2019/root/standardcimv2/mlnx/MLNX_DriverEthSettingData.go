// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DriverEthSettingData struct
type MLNX_DriverEthSettingData struct {
	*MLNX_DriverSettingData

	//
	ModeFlags uint32

	//
	NdkDebugFlags uint32

	//
	NdkDebugLevel uint32
}

func NewMLNX_DriverEthSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverEthSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverEthSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

func NewMLNX_DriverEthSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverEthSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverEthSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

// SetModeFlags sets the value of ModeFlags for the instance
func (instance *MLNX_DriverEthSettingData) SetPropertyModeFlags(value uint32) (err error) {
	return instance.SetProperty("ModeFlags", value)
}

// GetModeFlags gets the value of ModeFlags for the instance
func (instance *MLNX_DriverEthSettingData) GetPropertyModeFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugFlags sets the value of NdkDebugFlags for the instance
func (instance *MLNX_DriverEthSettingData) SetPropertyNdkDebugFlags(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags", value)
}

// GetNdkDebugFlags gets the value of NdkDebugFlags for the instance
func (instance *MLNX_DriverEthSettingData) GetPropertyNdkDebugFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugLevel sets the value of NdkDebugLevel for the instance
func (instance *MLNX_DriverEthSettingData) SetPropertyNdkDebugLevel(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel", value)
}

// GetNdkDebugLevel gets the value of NdkDebugLevel for the instance
func (instance *MLNX_DriverEthSettingData) GetPropertyNdkDebugLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ModeFlags" type="uint32 "></param>
// <param name="NdkDebugFlags" type="uint32 "></param>
// <param name="NdkDebugLevel" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverEthSettingData) SetValue( /* IN */ ModeFlags uint32,
	/* IN */ NdkDebugFlags uint32,
	/* IN */ NdkDebugLevel uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", ModeFlags, NdkDebugFlags, NdkDebugLevel)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverEthSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
