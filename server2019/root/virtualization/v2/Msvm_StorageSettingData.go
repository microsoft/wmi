// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_StorageSettingData struct
type Msvm_StorageSettingData struct {
	Msvm_SystemComponentSettingData

	//
	DisableInterruptBatching bool

	//
	ThreadCountPerChannel StorageSettingData_ThreadCountPerChannel

	//
	VirtualProcessorsPerChannel uint16
}

// SetDisableInterruptBatching sets the value of DisableInterruptBatching for the instance
func (instance *Msvm_StorageSettingData) SetPropertyDisableInterruptBatching(value bool) (err error) {
	return instance.SetProperty("DisableInterruptBatching", value)
}

// GetDisableInterruptBatching gets the value of DisableInterruptBatching for the instance
func (instance *Msvm_StorageSettingData) GetPropertyDisableInterruptBatching() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableInterruptBatching")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadCountPerChannel sets the value of ThreadCountPerChannel for the instance
func (instance *Msvm_StorageSettingData) SetPropertyThreadCountPerChannel(value StorageSettingData_ThreadCountPerChannel) (err error) {
	return instance.SetProperty("ThreadCountPerChannel", value)
}

// GetThreadCountPerChannel gets the value of ThreadCountPerChannel for the instance
func (instance *Msvm_StorageSettingData) GetPropertyThreadCountPerChannel() (value StorageSettingData_ThreadCountPerChannel, err error) {
	retValue, err := instance.GetProperty("ThreadCountPerChannel")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageSettingData_ThreadCountPerChannel)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualProcessorsPerChannel sets the value of VirtualProcessorsPerChannel for the instance
func (instance *Msvm_StorageSettingData) SetPropertyVirtualProcessorsPerChannel(value uint16) (err error) {
	return instance.SetProperty("VirtualProcessorsPerChannel", value)
}

// GetVirtualProcessorsPerChannel gets the value of VirtualProcessorsPerChannel for the instance
func (instance *Msvm_StorageSettingData) GetPropertyVirtualProcessorsPerChannel() (value uint16, err error) {
	retValue, err := instance.GetProperty("VirtualProcessorsPerChannel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
