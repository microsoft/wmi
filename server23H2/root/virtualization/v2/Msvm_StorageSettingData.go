// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_StorageSettingData struct
type Msvm_StorageSettingData struct {
	*Msvm_SystemComponentSettingData

	//
	DisableInterruptBatching bool

	//
	ThreadCountPerChannel StorageSettingData_ThreadCountPerChannel

	//
	VirtualProcessorsPerChannel uint16
}

func NewMsvm_StorageSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_StorageSettingData, err error) {
	tmp, err := NewMsvm_SystemComponentSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageSettingData{
		Msvm_SystemComponentSettingData: tmp,
	}
	return
}

func NewMsvm_StorageSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_StorageSettingData, err error) {
	tmp, err := NewMsvm_SystemComponentSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageSettingData{
		Msvm_SystemComponentSettingData: tmp,
	}
	return
}

// SetDisableInterruptBatching sets the value of DisableInterruptBatching for the instance
func (instance *Msvm_StorageSettingData) SetPropertyDisableInterruptBatching(value bool) (err error) {
	return instance.SetProperty("DisableInterruptBatching", (value))
}

// GetDisableInterruptBatching gets the value of DisableInterruptBatching for the instance
func (instance *Msvm_StorageSettingData) GetPropertyDisableInterruptBatching() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableInterruptBatching")
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

// SetThreadCountPerChannel sets the value of ThreadCountPerChannel for the instance
func (instance *Msvm_StorageSettingData) SetPropertyThreadCountPerChannel(value StorageSettingData_ThreadCountPerChannel) (err error) {
	return instance.SetProperty("ThreadCountPerChannel", (value))
}

// GetThreadCountPerChannel gets the value of ThreadCountPerChannel for the instance
func (instance *Msvm_StorageSettingData) GetPropertyThreadCountPerChannel() (value StorageSettingData_ThreadCountPerChannel, err error) {
	retValue, err := instance.GetProperty("ThreadCountPerChannel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = StorageSettingData_ThreadCountPerChannel(valuetmp)

	return
}

// SetVirtualProcessorsPerChannel sets the value of VirtualProcessorsPerChannel for the instance
func (instance *Msvm_StorageSettingData) SetPropertyVirtualProcessorsPerChannel(value uint16) (err error) {
	return instance.SetProperty("VirtualProcessorsPerChannel", (value))
}

// GetVirtualProcessorsPerChannel gets the value of VirtualProcessorsPerChannel for the instance
func (instance *Msvm_StorageSettingData) GetPropertyVirtualProcessorsPerChannel() (value uint16, err error) {
	retValue, err := instance.GetProperty("VirtualProcessorsPerChannel")
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
