// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_XHCITransferRing struct
type Win32_PerfRawData_Counters_XHCITransferRing struct {
	Win32_PerfRawData

	//
	BytesPerSec uint32

	//
	FailedTransferCount uint32

	//
	IsochTDFailuresPersec uint32

	//
	IsochTDPersec uint32

	//
	MissedServiceErrorCount uint32

	//
	TransfersPersec uint32

	//
	UnderrunOverruncount uint32
}

// SetBytesPerSec sets the value of BytesPerSec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesPerSec", value)
}

// GetBytesPerSec gets the value of BytesPerSec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedTransferCount sets the value of FailedTransferCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyFailedTransferCount(value uint32) (err error) {
	return instance.SetProperty("FailedTransferCount", value)
}

// GetFailedTransferCount gets the value of FailedTransferCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyFailedTransferCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedTransferCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsochTDFailuresPersec sets the value of IsochTDFailuresPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyIsochTDFailuresPersec(value uint32) (err error) {
	return instance.SetProperty("IsochTDFailuresPersec", value)
}

// GetIsochTDFailuresPersec gets the value of IsochTDFailuresPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyIsochTDFailuresPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsochTDFailuresPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsochTDPersec sets the value of IsochTDPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyIsochTDPersec(value uint32) (err error) {
	return instance.SetProperty("IsochTDPersec", value)
}

// GetIsochTDPersec gets the value of IsochTDPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyIsochTDPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsochTDPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMissedServiceErrorCount sets the value of MissedServiceErrorCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyMissedServiceErrorCount(value uint32) (err error) {
	return instance.SetProperty("MissedServiceErrorCount", value)
}

// GetMissedServiceErrorCount gets the value of MissedServiceErrorCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyMissedServiceErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MissedServiceErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransfersPersec sets the value of TransfersPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyTransfersPersec(value uint32) (err error) {
	return instance.SetProperty("TransfersPersec", value)
}

// GetTransfersPersec gets the value of TransfersPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyTransfersPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransfersPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnderrunOverruncount sets the value of UnderrunOverruncount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) SetPropertyUnderrunOverruncount(value uint32) (err error) {
	return instance.SetProperty("UnderrunOverruncount", value)
}

// GetUnderrunOverruncount gets the value of UnderrunOverruncount for the instance
func (instance *Win32_PerfRawData_Counters_XHCITransferRing) GetPropertyUnderrunOverruncount() (value uint32, err error) {
	retValue, err := instance.GetProperty("UnderrunOverruncount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
