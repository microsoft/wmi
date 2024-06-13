// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_KvpExchangeComponentSettingData struct
type Msvm_KvpExchangeComponentSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	DisableHostKVPItems bool

	//
	EnabledState uint16

	//
	HostExchangeItems []string

	//
	HostOnlyItems []string
}

func NewMsvm_KvpExchangeComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_KvpExchangeComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_KvpExchangeComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_KvpExchangeComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDisableHostKVPItems sets the value of DisableHostKVPItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyDisableHostKVPItems(value bool) (err error) {
	return instance.SetProperty("DisableHostKVPItems", (value))
}

// GetDisableHostKVPItems gets the value of DisableHostKVPItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyDisableHostKVPItems() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableHostKVPItems")
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

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
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

// SetHostExchangeItems sets the value of HostExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyHostExchangeItems(value []string) (err error) {
	return instance.SetProperty("HostExchangeItems", (value))
}

// GetHostExchangeItems gets the value of HostExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyHostExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("HostExchangeItems")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetHostOnlyItems sets the value of HostOnlyItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) SetPropertyHostOnlyItems(value []string) (err error) {
	return instance.SetProperty("HostOnlyItems", (value))
}

// GetHostOnlyItems gets the value of HostOnlyItems for the instance
func (instance *Msvm_KvpExchangeComponentSettingData) GetPropertyHostOnlyItems() (value []string, err error) {
	retValue, err := instance.GetProperty("HostOnlyItems")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
