// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM struct
type Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM struct {
	Win32_PerfRawData

	//
	AddedMemory uint64

	//
	AveragePressure uint32

	//
	CurrentPressure uint32

	//
	GuestAvailableMemory uint32

	//
	GuestVisiblePhysicalMemory uint32

	//
	MaximumPressure uint32

	//
	MemoryAddOperations uint64

	//
	MemoryRemoveOperations uint64

	//
	MinimumPressure uint32

	//
	PhysicalMemory uint32

	//
	RemovedMemory uint64

	//
	SmartPagingWorkingSetSize uint32
}

// SetAddedMemory sets the value of AddedMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyAddedMemory(value uint64) (err error) {
	return instance.SetProperty("AddedMemory", value)
}

// GetAddedMemory gets the value of AddedMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyAddedMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddedMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAveragePressure sets the value of AveragePressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyAveragePressure(value uint32) (err error) {
	return instance.SetProperty("AveragePressure", value)
}

// GetAveragePressure gets the value of AveragePressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyAveragePressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("AveragePressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentPressure sets the value of CurrentPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyCurrentPressure(value uint32) (err error) {
	return instance.SetProperty("CurrentPressure", value)
}

// GetCurrentPressure gets the value of CurrentPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyCurrentPressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentPressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuestAvailableMemory sets the value of GuestAvailableMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyGuestAvailableMemory(value uint32) (err error) {
	return instance.SetProperty("GuestAvailableMemory", value)
}

// GetGuestAvailableMemory gets the value of GuestAvailableMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyGuestAvailableMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("GuestAvailableMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuestVisiblePhysicalMemory sets the value of GuestVisiblePhysicalMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyGuestVisiblePhysicalMemory(value uint32) (err error) {
	return instance.SetProperty("GuestVisiblePhysicalMemory", value)
}

// GetGuestVisiblePhysicalMemory gets the value of GuestVisiblePhysicalMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyGuestVisiblePhysicalMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("GuestVisiblePhysicalMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumPressure sets the value of MaximumPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyMaximumPressure(value uint32) (err error) {
	return instance.SetProperty("MaximumPressure", value)
}

// GetMaximumPressure gets the value of MaximumPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyMaximumPressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumPressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryAddOperations sets the value of MemoryAddOperations for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyMemoryAddOperations(value uint64) (err error) {
	return instance.SetProperty("MemoryAddOperations", value)
}

// GetMemoryAddOperations gets the value of MemoryAddOperations for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyMemoryAddOperations() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryAddOperations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryRemoveOperations sets the value of MemoryRemoveOperations for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyMemoryRemoveOperations(value uint64) (err error) {
	return instance.SetProperty("MemoryRemoveOperations", value)
}

// GetMemoryRemoveOperations gets the value of MemoryRemoveOperations for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyMemoryRemoveOperations() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryRemoveOperations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumPressure sets the value of MinimumPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyMinimumPressure(value uint32) (err error) {
	return instance.SetProperty("MinimumPressure", value)
}

// GetMinimumPressure gets the value of MinimumPressure for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyMinimumPressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumPressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalMemory sets the value of PhysicalMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyPhysicalMemory(value uint32) (err error) {
	return instance.SetProperty("PhysicalMemory", value)
}

// GetPhysicalMemory gets the value of PhysicalMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyPhysicalMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovedMemory sets the value of RemovedMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertyRemovedMemory(value uint64) (err error) {
	return instance.SetProperty("RemovedMemory", value)
}

// GetRemovedMemory gets the value of RemovedMemory for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertyRemovedMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("RemovedMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmartPagingWorkingSetSize sets the value of SmartPagingWorkingSetSize for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) SetPropertySmartPagingWorkingSetSize(value uint32) (err error) {
	return instance.SetProperty("SmartPagingWorkingSetSize", value)
}

// GetSmartPagingWorkingSetSize gets the value of SmartPagingWorkingSetSize for the instance
func (instance *Win32_PerfRawData_BalancerStats_HyperVDynamicMemoryVM) GetPropertySmartPagingWorkingSetSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("SmartPagingWorkingSetSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
