// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor struct
type Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor struct {
	Win32_PerfFormattedData

	//
	NumberofTransmitCompletesPersec uint64

	//
	NumberofVMQs uint64

	//
	PacketsfromExternalPersec uint64

	//
	PacketsfromInternalPersec uint64
}

// SetNumberofTransmitCompletesPersec sets the value of NumberofTransmitCompletesPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) SetPropertyNumberofTransmitCompletesPersec(value uint64) (err error) {
	return instance.SetProperty("NumberofTransmitCompletesPersec", value)
}

// GetNumberofTransmitCompletesPersec gets the value of NumberofTransmitCompletesPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) GetPropertyNumberofTransmitCompletesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberofTransmitCompletesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofVMQs sets the value of NumberofVMQs for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) SetPropertyNumberofVMQs(value uint64) (err error) {
	return instance.SetProperty("NumberofVMQs", value)
}

// GetNumberofVMQs gets the value of NumberofVMQs for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) GetPropertyNumberofVMQs() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberofVMQs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsfromExternalPersec sets the value of PacketsfromExternalPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) SetPropertyPacketsfromExternalPersec(value uint64) (err error) {
	return instance.SetProperty("PacketsfromExternalPersec", value)
}

// GetPacketsfromExternalPersec gets the value of PacketsfromExternalPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) GetPropertyPacketsfromExternalPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsfromExternalPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsfromInternalPersec sets the value of PacketsfromInternalPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) SetPropertyPacketsfromInternalPersec(value uint64) (err error) {
	return instance.SetProperty("PacketsfromInternalPersec", value)
}

// GetPacketsfromInternalPersec gets the value of PacketsfromInternalPersec for the instance
func (instance *Win32_PerfFormattedData_NvspSwitchProcStats_HyperVVirtualSwitchProcessor) GetPropertyPacketsfromInternalPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsfromInternalPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
