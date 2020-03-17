// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO struct
type Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO struct {
	Win32_PerfRawData

	//
	ReceiveMessageQuotaExceeded uint64

	//
	ReceiveQoSConformantMessagesPersec uint64

	//
	ReceiveQoSExemptMessagesPersec uint64

	//
	ReceiveQoSNonConformantMessagesPersec uint64

	//
	ReceiveQoSTotalMessageDelayTime100ns uint64
}

// SetReceiveMessageQuotaExceeded sets the value of ReceiveMessageQuotaExceeded for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) SetPropertyReceiveMessageQuotaExceeded(value uint64) (err error) {
	return instance.SetProperty("ReceiveMessageQuotaExceeded", value)
}

// GetReceiveMessageQuotaExceeded gets the value of ReceiveMessageQuotaExceeded for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) GetPropertyReceiveMessageQuotaExceeded() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveMessageQuotaExceeded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveQoSConformantMessagesPersec sets the value of ReceiveQoSConformantMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) SetPropertyReceiveQoSConformantMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveQoSConformantMessagesPersec", value)
}

// GetReceiveQoSConformantMessagesPersec gets the value of ReceiveQoSConformantMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) GetPropertyReceiveQoSConformantMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveQoSConformantMessagesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveQoSExemptMessagesPersec sets the value of ReceiveQoSExemptMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) SetPropertyReceiveQoSExemptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveQoSExemptMessagesPersec", value)
}

// GetReceiveQoSExemptMessagesPersec gets the value of ReceiveQoSExemptMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) GetPropertyReceiveQoSExemptMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveQoSExemptMessagesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveQoSNonConformantMessagesPersec sets the value of ReceiveQoSNonConformantMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) SetPropertyReceiveQoSNonConformantMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveQoSNonConformantMessagesPersec", value)
}

// GetReceiveQoSNonConformantMessagesPersec gets the value of ReceiveQoSNonConformantMessagesPersec for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) GetPropertyReceiveQoSNonConformantMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveQoSNonConformantMessagesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveQoSTotalMessageDelayTime100ns sets the value of ReceiveQoSTotalMessageDelayTime100ns for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) SetPropertyReceiveQoSTotalMessageDelayTime100ns(value uint64) (err error) {
	return instance.SetProperty("ReceiveQoSTotalMessageDelayTime100ns", value)
}

// GetReceiveQoSTotalMessageDelayTime100ns gets the value of ReceiveQoSTotalMessageDelayTime100ns for the instance
func (instance *Win32_PerfRawData_VMBusPipeIOPerfProvider_HyperVVMVirtualDevicePipeIO) GetPropertyReceiveQoSTotalMessageDelayTime100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveQoSTotalMessageDelayTime100ns")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
