// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Perf struct
type Win32_Perf struct {
	CIM_StatisticalInformation

	//
	Frequency_Object uint64

	//
	Frequency_PerfTime uint64

	//
	Frequency_Sys100NS uint64

	//
	Timestamp_Object uint64

	//
	Timestamp_PerfTime uint64

	//
	Timestamp_Sys100NS uint64
}

// SetFrequency_Object sets the value of Frequency_Object for the instance
func (instance *Win32_Perf) SetPropertyFrequency_Object(value uint64) (err error) {
	return instance.SetProperty("Frequency_Object", value)
}

// GetFrequency_Object gets the value of Frequency_Object for the instance
func (instance *Win32_Perf) GetPropertyFrequency_Object() (value uint64, err error) {
	retValue, err := instance.GetProperty("Frequency_Object")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrequency_PerfTime sets the value of Frequency_PerfTime for the instance
func (instance *Win32_Perf) SetPropertyFrequency_PerfTime(value uint64) (err error) {
	return instance.SetProperty("Frequency_PerfTime", value)
}

// GetFrequency_PerfTime gets the value of Frequency_PerfTime for the instance
func (instance *Win32_Perf) GetPropertyFrequency_PerfTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Frequency_PerfTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrequency_Sys100NS sets the value of Frequency_Sys100NS for the instance
func (instance *Win32_Perf) SetPropertyFrequency_Sys100NS(value uint64) (err error) {
	return instance.SetProperty("Frequency_Sys100NS", value)
}

// GetFrequency_Sys100NS gets the value of Frequency_Sys100NS for the instance
func (instance *Win32_Perf) GetPropertyFrequency_Sys100NS() (value uint64, err error) {
	retValue, err := instance.GetProperty("Frequency_Sys100NS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp_Object sets the value of Timestamp_Object for the instance
func (instance *Win32_Perf) SetPropertyTimestamp_Object(value uint64) (err error) {
	return instance.SetProperty("Timestamp_Object", value)
}

// GetTimestamp_Object gets the value of Timestamp_Object for the instance
func (instance *Win32_Perf) GetPropertyTimestamp_Object() (value uint64, err error) {
	retValue, err := instance.GetProperty("Timestamp_Object")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp_PerfTime sets the value of Timestamp_PerfTime for the instance
func (instance *Win32_Perf) SetPropertyTimestamp_PerfTime(value uint64) (err error) {
	return instance.SetProperty("Timestamp_PerfTime", value)
}

// GetTimestamp_PerfTime gets the value of Timestamp_PerfTime for the instance
func (instance *Win32_Perf) GetPropertyTimestamp_PerfTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Timestamp_PerfTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp_Sys100NS sets the value of Timestamp_Sys100NS for the instance
func (instance *Win32_Perf) SetPropertyTimestamp_Sys100NS(value uint64) (err error) {
	return instance.SetProperty("Timestamp_Sys100NS", value)
}

// GetTimestamp_Sys100NS gets the value of Timestamp_Sys100NS for the instance
func (instance *Win32_Perf) GetPropertyTimestamp_Sys100NS() (value uint64, err error) {
	retValue, err := instance.GetProperty("Timestamp_Sys100NS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
