// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_TermService_TerminalServicesSession struct
type Win32_PerfFormattedData_TermService_TerminalServicesSession struct {
	Win32_PerfFormattedData

	//
	HandleCount uint32

	//
	PageFaultsPersec uint32

	//
	PageFileBytes uint64

	//
	PageFileBytesPeak uint64

	//
	PercentPrivilegedTime uint64

	//
	PercentProcessorTime uint64

	//
	PercentUserTime uint64

	//
	PoolNonpagedBytes uint32

	//
	PoolPagedBytes uint32

	//
	PrivateBytes uint64

	//
	ThreadCount uint32

	//
	VirtualBytes uint64

	//
	VirtualBytesPeak uint64

	//
	WorkingSet uint64

	//
	WorkingSetPeak uint64
}

// SetHandleCount sets the value of HandleCount for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyHandleCount(value uint32) (err error) {
	return instance.SetProperty("HandleCount", value)
}

// GetHandleCount gets the value of HandleCount for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyHandleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandleCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPageFaultsPersec sets the value of PageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPageFaultsPersec(value uint32) (err error) {
	return instance.SetProperty("PageFaultsPersec", value)
}

// GetPageFaultsPersec gets the value of PageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPageFaultsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageFaultsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPageFileBytes sets the value of PageFileBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPageFileBytes(value uint64) (err error) {
	return instance.SetProperty("PageFileBytes", value)
}

// GetPageFileBytes gets the value of PageFileBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPageFileBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFileBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPageFileBytesPeak sets the value of PageFileBytesPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPageFileBytesPeak(value uint64) (err error) {
	return instance.SetProperty("PageFileBytesPeak", value)
}

// GetPageFileBytesPeak gets the value of PageFileBytesPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPageFileBytesPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFileBytesPeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentPrivilegedTime sets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPercentPrivilegedTime(value uint64) (err error) {
	return instance.SetProperty("PercentPrivilegedTime", value)
}

// GetPercentPrivilegedTime gets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPercentPrivilegedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPrivilegedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentProcessorTime sets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPercentProcessorTime(value uint64) (err error) {
	return instance.SetProperty("PercentProcessorTime", value)
}

// GetPercentProcessorTime gets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPercentProcessorTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentProcessorTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUserTime sets the value of PercentUserTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPercentUserTime(value uint64) (err error) {
	return instance.SetProperty("PercentUserTime", value)
}

// GetPercentUserTime gets the value of PercentUserTime for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPercentUserTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentUserTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPoolNonpagedBytes sets the value of PoolNonpagedBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPoolNonpagedBytes(value uint32) (err error) {
	return instance.SetProperty("PoolNonpagedBytes", value)
}

// GetPoolNonpagedBytes gets the value of PoolNonpagedBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPoolNonpagedBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolNonpagedBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPoolPagedBytes sets the value of PoolPagedBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPoolPagedBytes(value uint32) (err error) {
	return instance.SetProperty("PoolPagedBytes", value)
}

// GetPoolPagedBytes gets the value of PoolPagedBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPoolPagedBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolPagedBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivateBytes sets the value of PrivateBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyPrivateBytes(value uint64) (err error) {
	return instance.SetProperty("PrivateBytes", value)
}

// GetPrivateBytes gets the value of PrivateBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyPrivateBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PrivateBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadCount sets the value of ThreadCount for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyThreadCount(value uint32) (err error) {
	return instance.SetProperty("ThreadCount", value)
}

// GetThreadCount gets the value of ThreadCount for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyThreadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualBytes sets the value of VirtualBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyVirtualBytes(value uint64) (err error) {
	return instance.SetProperty("VirtualBytes", value)
}

// GetVirtualBytes gets the value of VirtualBytes for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyVirtualBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualBytesPeak sets the value of VirtualBytesPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyVirtualBytesPeak(value uint64) (err error) {
	return instance.SetProperty("VirtualBytesPeak", value)
}

// GetVirtualBytesPeak gets the value of VirtualBytesPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyVirtualBytesPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualBytesPeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkingSet sets the value of WorkingSet for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyWorkingSet(value uint64) (err error) {
	return instance.SetProperty("WorkingSet", value)
}

// GetWorkingSet gets the value of WorkingSet for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyWorkingSet() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSet")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkingSetPeak sets the value of WorkingSetPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) SetPropertyWorkingSetPeak(value uint64) (err error) {
	return instance.SetProperty("WorkingSetPeak", value)
}

// GetWorkingSetPeak gets the value of WorkingSetPeak for the instance
func (instance *Win32_PerfFormattedData_TermService_TerminalServicesSession) GetPropertyWorkingSetPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSetPeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
