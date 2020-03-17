// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_HeartbeatComponentSettingData struct
type Msvm_HeartbeatComponentSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	EnabledState uint16

	//
	ErrorThreshold uint32

	//
	Interval uint32

	//
	Latency uint32
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
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

// SetErrorThreshold sets the value of ErrorThreshold for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyErrorThreshold(value uint32) (err error) {
	return instance.SetProperty("ErrorThreshold", value)
}

// GetErrorThreshold gets the value of ErrorThreshold for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyErrorThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterval sets the value of Interval for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyInterval(value uint32) (err error) {
	return instance.SetProperty("Interval", value)
}

// GetInterval gets the value of Interval for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("Interval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLatency sets the value of Latency for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", value)
}

// GetLatency gets the value of Latency for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
