// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated struct
type Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated struct {
	Win32_PerfFormattedData

	//
	ReadBytesPersec uint64

	//
	ReadSectorsPersec uint64

	//
	WriteBytesPersec uint64

	//
	WrittenSectorsPersec uint64
}

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", value)
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadSectorsPersec sets the value of ReadSectorsPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) SetPropertyReadSectorsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadSectorsPersec", value)
}

// GetReadSectorsPersec gets the value of ReadSectorsPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) GetPropertyReadSectorsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadSectorsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", value)
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWrittenSectorsPersec sets the value of WrittenSectorsPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) SetPropertyWrittenSectorsPersec(value uint64) (err error) {
	return instance.SetProperty("WrittenSectorsPersec", value)
}

// GetWrittenSectorsPersec gets the value of WrittenSectorsPersec for the instance
func (instance *Win32_PerfFormattedData_IdePerfProvider_HyperVVirtualIDEControllerEmulated) GetPropertyWrittenSectorsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WrittenSectorsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
