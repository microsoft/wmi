// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes struct
type Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes struct {
	Win32_PerfRawData

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
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) SetPropertyBytesReadPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReadPersec", value)
}

// GetBytesReadPersec gets the value of BytesReadPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) GetPropertyBytesReadPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) SetPropertyBytesWrittenPersec(value uint64) (err error) {
	return instance.SetProperty("BytesWrittenPersec", value)
}

// GetBytesWrittenPersec gets the value of BytesWrittenPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) GetPropertyBytesWrittenPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) SetPropertyReadsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsPersec", value)
}

// GetReadsPersec gets the value of ReadsPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) GetPropertyReadsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) SetPropertyWritesPersec(value uint64) (err error) {
	return instance.SetProperty("WritesPersec", value)
}

// GetWritesPersec gets the value of WritesPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBusProviderPipes) GetPropertyWritesPersec() (value uint64, err error) {
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
