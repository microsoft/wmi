// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition struct
type Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition struct {
	Win32_PerfFormattedData

	//
	AddressSpaces uint64

	//
	AttachedDevices uint64

	//
	DepositedPages uint64

	//
	DeviceDMAErrors uint64

	//
	DeviceInterruptErrors uint64

	//
	DeviceInterruptMappings uint64

	//
	DeviceInterruptThrottleEvents uint64

	//
	GPAPages uint64

	//
	GPASpaceModificationsPersec uint64

	//
	IOTLBFlushCost uint64

	//
	IOTLBFlushesPersec uint64

	//
	NestedTLBFreeListSize uint64

	//
	NestedTLBSize uint64

	//
	NestedTLBTrimmedPagesPersec uint64

	//
	Pages1Persec uint64

	//
	Pages2Persec uint64

	//
	RecommendedNestedTLBSize uint64

	//
	RecommendedVirtualTLBSize uint64

	//
	SkippedTimerTicks uint64

	//
	Value1Gdevicepages uint64

	//
	Value1GGPApages uint64

	//
	Value2Mdevicepages uint64

	//
	Value2MGPApages uint64

	//
	Value4Kdevicepages uint64

	//
	Value4KGPApages uint64

	//
	VirtualProcessors uint64

	//
	VirtualTLBFlushEntiresPersec uint64

	//
	VirtualTLBPages uint64
}

// SetAddressSpaces sets the value of AddressSpaces for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyAddressSpaces(value uint64) (err error) {
	return instance.SetProperty("AddressSpaces", value)
}

// GetAddressSpaces gets the value of AddressSpaces for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyAddressSpaces() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddressSpaces")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttachedDevices sets the value of AttachedDevices for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyAttachedDevices(value uint64) (err error) {
	return instance.SetProperty("AttachedDevices", value)
}

// GetAttachedDevices gets the value of AttachedDevices for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyAttachedDevices() (value uint64, err error) {
	retValue, err := instance.GetProperty("AttachedDevices")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDepositedPages sets the value of DepositedPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyDepositedPages(value uint64) (err error) {
	return instance.SetProperty("DepositedPages", value)
}

// GetDepositedPages gets the value of DepositedPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyDepositedPages() (value uint64, err error) {
	retValue, err := instance.GetProperty("DepositedPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceDMAErrors sets the value of DeviceDMAErrors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyDeviceDMAErrors(value uint64) (err error) {
	return instance.SetProperty("DeviceDMAErrors", value)
}

// GetDeviceDMAErrors gets the value of DeviceDMAErrors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyDeviceDMAErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceDMAErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceInterruptErrors sets the value of DeviceInterruptErrors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyDeviceInterruptErrors(value uint64) (err error) {
	return instance.SetProperty("DeviceInterruptErrors", value)
}

// GetDeviceInterruptErrors gets the value of DeviceInterruptErrors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyDeviceInterruptErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceInterruptErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceInterruptMappings sets the value of DeviceInterruptMappings for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyDeviceInterruptMappings(value uint64) (err error) {
	return instance.SetProperty("DeviceInterruptMappings", value)
}

// GetDeviceInterruptMappings gets the value of DeviceInterruptMappings for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyDeviceInterruptMappings() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceInterruptMappings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceInterruptThrottleEvents sets the value of DeviceInterruptThrottleEvents for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyDeviceInterruptThrottleEvents(value uint64) (err error) {
	return instance.SetProperty("DeviceInterruptThrottleEvents", value)
}

// GetDeviceInterruptThrottleEvents gets the value of DeviceInterruptThrottleEvents for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyDeviceInterruptThrottleEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceInterruptThrottleEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGPAPages sets the value of GPAPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyGPAPages(value uint64) (err error) {
	return instance.SetProperty("GPAPages", value)
}

// GetGPAPages gets the value of GPAPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyGPAPages() (value uint64, err error) {
	retValue, err := instance.GetProperty("GPAPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGPASpaceModificationsPersec sets the value of GPASpaceModificationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyGPASpaceModificationsPersec(value uint64) (err error) {
	return instance.SetProperty("GPASpaceModificationsPersec", value)
}

// GetGPASpaceModificationsPersec gets the value of GPASpaceModificationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyGPASpaceModificationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GPASpaceModificationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOTLBFlushCost sets the value of IOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyIOTLBFlushCost(value uint64) (err error) {
	return instance.SetProperty("IOTLBFlushCost", value)
}

// GetIOTLBFlushCost gets the value of IOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyIOTLBFlushCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOTLBFlushCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOTLBFlushesPersec sets the value of IOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyIOTLBFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("IOTLBFlushesPersec", value)
}

// GetIOTLBFlushesPersec gets the value of IOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyIOTLBFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOTLBFlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNestedTLBFreeListSize sets the value of NestedTLBFreeListSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyNestedTLBFreeListSize(value uint64) (err error) {
	return instance.SetProperty("NestedTLBFreeListSize", value)
}

// GetNestedTLBFreeListSize gets the value of NestedTLBFreeListSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyNestedTLBFreeListSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedTLBFreeListSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNestedTLBSize sets the value of NestedTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyNestedTLBSize(value uint64) (err error) {
	return instance.SetProperty("NestedTLBSize", value)
}

// GetNestedTLBSize gets the value of NestedTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyNestedTLBSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedTLBSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNestedTLBTrimmedPagesPersec sets the value of NestedTLBTrimmedPagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyNestedTLBTrimmedPagesPersec(value uint64) (err error) {
	return instance.SetProperty("NestedTLBTrimmedPagesPersec", value)
}

// GetNestedTLBTrimmedPagesPersec gets the value of NestedTLBTrimmedPagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyNestedTLBTrimmedPagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedTLBTrimmedPagesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPages1Persec sets the value of Pages1Persec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyPages1Persec(value uint64) (err error) {
	return instance.SetProperty("Pages1Persec", value)
}

// GetPages1Persec gets the value of Pages1Persec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyPages1Persec() (value uint64, err error) {
	retValue, err := instance.GetProperty("Pages1Persec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPages2Persec sets the value of Pages2Persec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyPages2Persec(value uint64) (err error) {
	return instance.SetProperty("Pages2Persec", value)
}

// GetPages2Persec gets the value of Pages2Persec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyPages2Persec() (value uint64, err error) {
	retValue, err := instance.GetProperty("Pages2Persec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecommendedNestedTLBSize sets the value of RecommendedNestedTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyRecommendedNestedTLBSize(value uint64) (err error) {
	return instance.SetProperty("RecommendedNestedTLBSize", value)
}

// GetRecommendedNestedTLBSize gets the value of RecommendedNestedTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyRecommendedNestedTLBSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecommendedNestedTLBSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecommendedVirtualTLBSize sets the value of RecommendedVirtualTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyRecommendedVirtualTLBSize(value uint64) (err error) {
	return instance.SetProperty("RecommendedVirtualTLBSize", value)
}

// GetRecommendedVirtualTLBSize gets the value of RecommendedVirtualTLBSize for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyRecommendedVirtualTLBSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecommendedVirtualTLBSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSkippedTimerTicks sets the value of SkippedTimerTicks for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertySkippedTimerTicks(value uint64) (err error) {
	return instance.SetProperty("SkippedTimerTicks", value)
}

// GetSkippedTimerTicks gets the value of SkippedTimerTicks for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertySkippedTimerTicks() (value uint64, err error) {
	retValue, err := instance.GetProperty("SkippedTimerTicks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue1Gdevicepages sets the value of Value1Gdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue1Gdevicepages(value uint64) (err error) {
	return instance.SetProperty("Value1Gdevicepages", value)
}

// GetValue1Gdevicepages gets the value of Value1Gdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue1Gdevicepages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value1Gdevicepages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue1GGPApages sets the value of Value1GGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue1GGPApages(value uint64) (err error) {
	return instance.SetProperty("Value1GGPApages", value)
}

// GetValue1GGPApages gets the value of Value1GGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue1GGPApages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value1GGPApages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue2Mdevicepages sets the value of Value2Mdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue2Mdevicepages(value uint64) (err error) {
	return instance.SetProperty("Value2Mdevicepages", value)
}

// GetValue2Mdevicepages gets the value of Value2Mdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue2Mdevicepages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value2Mdevicepages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue2MGPApages sets the value of Value2MGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue2MGPApages(value uint64) (err error) {
	return instance.SetProperty("Value2MGPApages", value)
}

// GetValue2MGPApages gets the value of Value2MGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue2MGPApages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value2MGPApages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue4Kdevicepages sets the value of Value4Kdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue4Kdevicepages(value uint64) (err error) {
	return instance.SetProperty("Value4Kdevicepages", value)
}

// GetValue4Kdevicepages gets the value of Value4Kdevicepages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue4Kdevicepages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value4Kdevicepages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue4KGPApages sets the value of Value4KGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyValue4KGPApages(value uint64) (err error) {
	return instance.SetProperty("Value4KGPApages", value)
}

// GetValue4KGPApages gets the value of Value4KGPApages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyValue4KGPApages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value4KGPApages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualProcessors sets the value of VirtualProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyVirtualProcessors(value uint64) (err error) {
	return instance.SetProperty("VirtualProcessors", value)
}

// GetVirtualProcessors gets the value of VirtualProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyVirtualProcessors() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualTLBFlushEntiresPersec sets the value of VirtualTLBFlushEntiresPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyVirtualTLBFlushEntiresPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualTLBFlushEntiresPersec", value)
}

// GetVirtualTLBFlushEntiresPersec gets the value of VirtualTLBFlushEntiresPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyVirtualTLBFlushEntiresPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualTLBFlushEntiresPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualTLBPages sets the value of VirtualTLBPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) SetPropertyVirtualTLBPages(value uint64) (err error) {
	return instance.SetProperty("VirtualTLBPages", value)
}

// GetVirtualTLBPages gets the value of VirtualTLBPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorPartition) GetPropertyVirtualTLBPages() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualTLBPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
