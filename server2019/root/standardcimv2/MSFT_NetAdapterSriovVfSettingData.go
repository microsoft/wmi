// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterSriovVfSettingData struct
type MSFT_NetAdapterSriovVfSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	CurrentMacAddress string

	//
	FunctionID uint16

	//
	PermanentMacAddress string

	//
	SwitchID uint32

	//
	VmFriendlyName string

	//
	VmID string

	//
	VmNicID string

	//
	VPortID []uint32
}

func NewMSFT_NetAdapterSriovVfSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterSriovVfSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovVfSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterSriovVfSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterSriovVfSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovVfSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetCurrentMacAddress sets the value of CurrentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyCurrentMacAddress(value string) (err error) {
	return instance.SetProperty("CurrentMacAddress", (value))
}

// GetCurrentMacAddress gets the value of CurrentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyCurrentMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentMacAddress")
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

// SetFunctionID sets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyFunctionID(value uint16) (err error) {
	return instance.SetProperty("FunctionID", (value))
}

// GetFunctionID gets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyFunctionID() (value uint16, err error) {
	retValue, err := instance.GetProperty("FunctionID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPermanentMacAddress sets the value of PermanentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyPermanentMacAddress(value string) (err error) {
	return instance.SetProperty("PermanentMacAddress", (value))
}

// GetPermanentMacAddress gets the value of PermanentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyPermanentMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PermanentMacAddress")
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

// SetSwitchID sets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertySwitchID(value uint32) (err error) {
	return instance.SetProperty("SwitchID", (value))
}

// GetSwitchID gets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertySwitchID() (value uint32, err error) {
	retValue, err := instance.GetProperty("SwitchID")
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

// SetVmFriendlyName sets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmFriendlyName(value string) (err error) {
	return instance.SetProperty("VmFriendlyName", (value))
}

// GetVmFriendlyName gets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("VmFriendlyName")
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

// SetVmID sets the value of VmID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmID(value string) (err error) {
	return instance.SetProperty("VmID", (value))
}

// GetVmID gets the value of VmID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmID() (value string, err error) {
	retValue, err := instance.GetProperty("VmID")
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

// SetVmNicID sets the value of VmNicID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmNicID(value string) (err error) {
	return instance.SetProperty("VmNicID", (value))
}

// GetVmNicID gets the value of VmNicID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmNicID() (value string, err error) {
	retValue, err := instance.GetProperty("VmNicID")
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

// SetVPortID sets the value of VPortID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVPortID(value []uint32) (err error) {
	return instance.SetProperty("VPortID", (value))
}

// GetVPortID gets the value of VPortID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVPortID() (value []uint32, err error) {
	retValue, err := instance.GetProperty("VPortID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}
