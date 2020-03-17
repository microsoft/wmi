// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_EventTracingforWindows struct
type Win32_PerfRawData_Counters_EventTracingforWindows struct {
	Win32_PerfRawData

	//
	TotalMemoryUsageNonPagedPool uint32

	//
	TotalMemoryUsagePagedPool uint32

	//
	TotalNumberofActiveSessions uint32

	//
	TotalNumberofDistinctDisabledProviders uint32

	//
	TotalNumberofDistinctEnabledProviders uint32

	//
	TotalNumberofDistinctPreEnabledProviders uint32
}

// SetTotalMemoryUsageNonPagedPool sets the value of TotalMemoryUsageNonPagedPool for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalMemoryUsageNonPagedPool(value uint32) (err error) {
	return instance.SetProperty("TotalMemoryUsageNonPagedPool", value)
}

// GetTotalMemoryUsageNonPagedPool gets the value of TotalMemoryUsageNonPagedPool for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalMemoryUsageNonPagedPool() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMemoryUsageNonPagedPool")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalMemoryUsagePagedPool sets the value of TotalMemoryUsagePagedPool for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalMemoryUsagePagedPool(value uint32) (err error) {
	return instance.SetProperty("TotalMemoryUsagePagedPool", value)
}

// GetTotalMemoryUsagePagedPool gets the value of TotalMemoryUsagePagedPool for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalMemoryUsagePagedPool() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMemoryUsagePagedPool")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofActiveSessions sets the value of TotalNumberofActiveSessions for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalNumberofActiveSessions(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofActiveSessions", value)
}

// GetTotalNumberofActiveSessions gets the value of TotalNumberofActiveSessions for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalNumberofActiveSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofActiveSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofDistinctDisabledProviders sets the value of TotalNumberofDistinctDisabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalNumberofDistinctDisabledProviders(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofDistinctDisabledProviders", value)
}

// GetTotalNumberofDistinctDisabledProviders gets the value of TotalNumberofDistinctDisabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalNumberofDistinctDisabledProviders() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofDistinctDisabledProviders")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofDistinctEnabledProviders sets the value of TotalNumberofDistinctEnabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalNumberofDistinctEnabledProviders(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofDistinctEnabledProviders", value)
}

// GetTotalNumberofDistinctEnabledProviders gets the value of TotalNumberofDistinctEnabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalNumberofDistinctEnabledProviders() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofDistinctEnabledProviders")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofDistinctPreEnabledProviders sets the value of TotalNumberofDistinctPreEnabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) SetPropertyTotalNumberofDistinctPreEnabledProviders(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofDistinctPreEnabledProviders", value)
}

// GetTotalNumberofDistinctPreEnabledProviders gets the value of TotalNumberofDistinctPreEnabledProviders for the instance
func (instance *Win32_PerfRawData_Counters_EventTracingforWindows) GetPropertyTotalNumberofDistinctPreEnabledProviders() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofDistinctPreEnabledProviders")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
