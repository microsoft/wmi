// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_SMBServer struct
type Win32_PerfRawData_Counters_SMBServer struct {
	Win32_PerfRawData

	//
	ReadBytesPersec uint64

	//
	ReadRequestsPersec uint64

	//
	ReceiveBytesPersec uint64

	//
	SendBytesPersec uint64

	//
	WriteBytesPersec uint64

	//
	WriteRequestsPersec uint64
}

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", value)
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertyReadBytesPersec() (value uint64, err error) {
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

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertyReadRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadRequestsPersec", value)
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertyReadRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveBytesPersec sets the value of ReceiveBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertyReceiveBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveBytesPersec", value)
}

// GetReceiveBytesPersec gets the value of ReceiveBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertyReceiveBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendBytesPersec sets the value of SendBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertySendBytesPersec(value uint64) (err error) {
	return instance.SetProperty("SendBytesPersec", value)
}

// GetSendBytesPersec gets the value of SendBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertySendBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendBytesPersec")
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
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", value)
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertyWriteBytesPersec() (value uint64, err error) {
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

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) SetPropertyWriteRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("WriteRequestsPersec", value)
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServer) GetPropertyWriteRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
