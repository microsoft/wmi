// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_KvpExchangeComponentSettingData struct
type Msvm_KvpExchangeComponentSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	DisableHostKVPItems bool

	//
	EnabledState uint16

	//
	HostExchangeItems []string

	//
	HostOnlyItems []string
}

// SetDisableHostKVPItems sets the value of DisableHostKVPItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyDisableHostKVPItems(value bool) (err error) {
	return instance.SetProperty("DisableHostKVPItems", value)
}

// GetDisableHostKVPItems gets the value of DisableHostKVPItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyDisableHostKVPItems() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableHostKVPItems")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostExchangeItems sets the value of HostExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyHostExchangeItems(value []string) (err error) {
	return instance.SetProperty("HostExchangeItems", value)
}

// GetHostExchangeItems gets the value of HostExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyHostExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("HostExchangeItems")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostOnlyItems sets the value of HostOnlyItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyHostOnlyItems(value []string) (err error) {
	return instance.SetProperty("HostOnlyItems", value)
}

// GetHostOnlyItems gets the value of HostOnlyItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyHostOnlyItems() (value []string, err error) {
	retValue, err := instance.GetProperty("HostOnlyItems")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
