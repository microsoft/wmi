// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes struct
type Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes struct {
	Win32_PerfFormattedData

	//
	BytesReadPersec uint64

	//
	BytesWrittenPersec uint64

	//
	ReadsPersec uint64

	//
	WritesPersec uint64
}

// SetBytesReadPersec sets the value of BytesReadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) SetPropertyBytesReadPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReadPersec", value)
}

// GetBytesReadPersec gets the value of BytesReadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) GetPropertyBytesReadPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReadPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesWrittenPersec sets the value of BytesWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) SetPropertyBytesWrittenPersec(value uint64) (err error) {
	return instance.SetProperty("BytesWrittenPersec", value)
}

// GetBytesWrittenPersec gets the value of BytesWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) GetPropertyBytesWrittenPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesWrittenPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadsPersec sets the value of ReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) SetPropertyReadsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsPersec", value)
}

// GetReadsPersec gets the value of ReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) GetPropertyReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWritesPersec sets the value of WritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) SetPropertyWritesPersec(value uint64) (err error) {
	return instance.SetProperty("WritesPersec", value)
}

// GetWritesPersec gets the value of WritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVVirtualMachineBusPipes) GetPropertyWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
