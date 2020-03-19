// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterBindingSettingData struct
type MSFT_NetAdapterBindingSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	BindName string

	//
	Characteristics uint32

	//
	ComponentClassGuid string

	//
	ComponentClassName string

	//
	ComponentID string

	//
	DisplayName string

	//
	Enabled bool
}

func NewMSFT_NetAdapterBindingSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterBindingSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterBindingSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterBindingSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterBindingSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterBindingSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetBindName sets the value of BindName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyBindName(value string) (err error) {
	return instance.SetProperty("BindName", value)
}

// GetBindName gets the value of BindName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyBindName() (value string, err error) {
	retValue, err := instance.GetProperty("BindName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyCharacteristics(value uint32) (err error) {
	return instance.SetProperty("Characteristics", value)
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyCharacteristics() (value uint32, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponentClassGuid sets the value of ComponentClassGuid for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyComponentClassGuid(value string) (err error) {
	return instance.SetProperty("ComponentClassGuid", value)
}

// GetComponentClassGuid gets the value of ComponentClassGuid for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyComponentClassGuid() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentClassGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponentClassName sets the value of ComponentClassName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyComponentClassName(value string) (err error) {
	return instance.SetProperty("ComponentClassName", value)
}

// GetComponentClassName gets the value of ComponentClassName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyComponentClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponentID sets the value of ComponentID for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyComponentID(value string) (err error) {
	return instance.SetProperty("ComponentID", value)
}

// GetComponentID gets the value of ComponentID for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyComponentID() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterBindingSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterBindingSettingData) GetPropertyEnabled() (value bool, err error) {
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

//

// <param name="cmdletOutput" type="MSFT_NetAdapterBindingSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterBindingSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterBindingSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterBindingSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterBindingSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterBindingSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
