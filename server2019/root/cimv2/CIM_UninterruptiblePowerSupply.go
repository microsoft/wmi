// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_UninterruptiblePowerSupply struct
type CIM_UninterruptiblePowerSupply struct {
	CIM_PowerSupply

	//
	EstimatedChargeRemaining uint16

	//
	EstimatedRunTime uint32

	//
	RemainingCapacityStatus uint16

	//
	TimeOnBackup uint32
}

// SetEstimatedChargeRemaining sets the value of EstimatedChargeRemaining for the instance
func (instance *CIM_UninterruptiblePowerSupply) SetPropertyEstimatedChargeRemaining(value uint16) (err error) {
	return instance.SetProperty("EstimatedChargeRemaining", value)
}

// GetEstimatedChargeRemaining gets the value of EstimatedChargeRemaining for the instance
func (instance *CIM_UninterruptiblePowerSupply) GetPropertyEstimatedChargeRemaining() (value uint16, err error) {
	retValue, err := instance.GetProperty("EstimatedChargeRemaining")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEstimatedRunTime sets the value of EstimatedRunTime for the instance
func (instance *CIM_UninterruptiblePowerSupply) SetPropertyEstimatedRunTime(value uint32) (err error) {
	return instance.SetProperty("EstimatedRunTime", value)
}

// GetEstimatedRunTime gets the value of EstimatedRunTime for the instance
func (instance *CIM_UninterruptiblePowerSupply) GetPropertyEstimatedRunTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstimatedRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemainingCapacityStatus sets the value of RemainingCapacityStatus for the instance
func (instance *CIM_UninterruptiblePowerSupply) SetPropertyRemainingCapacityStatus(value uint16) (err error) {
	return instance.SetProperty("RemainingCapacityStatus", value)
}

// GetRemainingCapacityStatus gets the value of RemainingCapacityStatus for the instance
func (instance *CIM_UninterruptiblePowerSupply) GetPropertyRemainingCapacityStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemainingCapacityStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOnBackup sets the value of TimeOnBackup for the instance
func (instance *CIM_UninterruptiblePowerSupply) SetPropertyTimeOnBackup(value uint32) (err error) {
	return instance.SetProperty("TimeOnBackup", value)
}

// GetTimeOnBackup gets the value of TimeOnBackup for the instance
func (instance *CIM_UninterruptiblePowerSupply) GetPropertyTimeOnBackup() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeOnBackup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
