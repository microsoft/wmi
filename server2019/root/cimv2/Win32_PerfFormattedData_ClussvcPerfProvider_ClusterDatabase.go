// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase struct {
	Win32_PerfFormattedData

	//
	Flushes uint64

	//
	FlushesPersec uint64
}

// SetFlushes sets the value of Flushes for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase) SetPropertyFlushes(value uint64) (err error) {
	return instance.SetProperty("Flushes", value)
}

// GetFlushes gets the value of Flushes for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase) GetPropertyFlushes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Flushes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlushesPersec sets the value of FlushesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase) SetPropertyFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("FlushesPersec", value)
}

// GetFlushesPersec gets the value of FlushesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterDatabase) GetPropertyFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
