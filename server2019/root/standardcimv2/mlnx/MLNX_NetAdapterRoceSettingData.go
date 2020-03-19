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

// MLNX_NetAdapterRoceSettingData struct
type MLNX_NetAdapterRoceSettingData struct {
	*MLNX_NetAdapterSettingData

	//
	Enabled bool

	//
	PortNumber uint16

	//
	RoceMode string
}

func NewMLNX_NetAdapterRoceSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterRoceSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterRoceSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

func NewMLNX_NetAdapterRoceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterRoceSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterRoceSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MLNX_NetAdapterRoceSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MLNX_NetAdapterRoceSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *MLNX_NetAdapterRoceSettingData) SetPropertyPortNumber(value uint16) (err error) {
	return instance.SetProperty("PortNumber", value)
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *MLNX_NetAdapterRoceSettingData) GetPropertyPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoceMode sets the value of RoceMode for the instance
func (instance *MLNX_NetAdapterRoceSettingData) SetPropertyRoceMode(value string) (err error) {
	return instance.SetProperty("RoceMode", value)
}

// GetRoceMode gets the value of RoceMode for the instance
func (instance *MLNX_NetAdapterRoceSettingData) GetPropertyRoceMode() (value string, err error) {
	retValue, err := instance.GetProperty("RoceMode")
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

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapterRoceSettingData) Disable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapterRoceSettingData) Enable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}
