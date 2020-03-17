// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000 struct
type Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000 struct {
	Win32_PerfRawData

	//
	BytesReceived uint64

	//
	BytesSent uint64

	//
	ConnectionsEstablished uint32

	//
	DatagramsReceived uint32

	//
	DatagramsSent uint32

	//
	HttpWebRequestsAbortedPerSec uint32

	//
	HttpWebRequestsAverageLifetime uint64

	//
	HttpWebRequestsAverageLifetime_Base uint32

	//
	HttpWebRequestsAverageQueueTime uint64

	//
	HttpWebRequestsAverageQueueTime_Base uint32

	//
	HttpWebRequestsCreatedPerSec uint32

	//
	HttpWebRequestsFailedPerSec uint32

	//
	HttpWebRequestsQueuedPerSec uint32
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", value)
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionsEstablished sets the value of ConnectionsEstablished for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyConnectionsEstablished(value uint32) (err error) {
	return instance.SetProperty("ConnectionsEstablished", value)
}

// GetConnectionsEstablished gets the value of ConnectionsEstablished for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyConnectionsEstablished() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionsEstablished")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceived sets the value of DatagramsReceived for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyDatagramsReceived(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceived", value)
}

// GetDatagramsReceived gets the value of DatagramsReceived for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyDatagramsReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsSent sets the value of DatagramsSent for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyDatagramsSent(value uint32) (err error) {
	return instance.SetProperty("DatagramsSent", value)
}

// GetDatagramsSent gets the value of DatagramsSent for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyDatagramsSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsAbortedPerSec sets the value of HttpWebRequestsAbortedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsAbortedPerSec(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsAbortedPerSec", value)
}

// GetHttpWebRequestsAbortedPerSec gets the value of HttpWebRequestsAbortedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsAbortedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsAbortedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsAverageLifetime sets the value of HttpWebRequestsAverageLifetime for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsAverageLifetime(value uint64) (err error) {
	return instance.SetProperty("HttpWebRequestsAverageLifetime", value)
}

// GetHttpWebRequestsAverageLifetime gets the value of HttpWebRequestsAverageLifetime for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsAverageLifetime() (value uint64, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsAverageLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsAverageLifetime_Base sets the value of HttpWebRequestsAverageLifetime_Base for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsAverageLifetime_Base(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsAverageLifetime_Base", value)
}

// GetHttpWebRequestsAverageLifetime_Base gets the value of HttpWebRequestsAverageLifetime_Base for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsAverageLifetime_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsAverageLifetime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsAverageQueueTime sets the value of HttpWebRequestsAverageQueueTime for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsAverageQueueTime(value uint64) (err error) {
	return instance.SetProperty("HttpWebRequestsAverageQueueTime", value)
}

// GetHttpWebRequestsAverageQueueTime gets the value of HttpWebRequestsAverageQueueTime for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsAverageQueueTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsAverageQueueTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsAverageQueueTime_Base sets the value of HttpWebRequestsAverageQueueTime_Base for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsAverageQueueTime_Base(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsAverageQueueTime_Base", value)
}

// GetHttpWebRequestsAverageQueueTime_Base gets the value of HttpWebRequestsAverageQueueTime_Base for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsAverageQueueTime_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsAverageQueueTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsCreatedPerSec sets the value of HttpWebRequestsCreatedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsCreatedPerSec(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsCreatedPerSec", value)
}

// GetHttpWebRequestsCreatedPerSec gets the value of HttpWebRequestsCreatedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsCreatedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsCreatedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsFailedPerSec sets the value of HttpWebRequestsFailedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsFailedPerSec(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsFailedPerSec", value)
}

// GetHttpWebRequestsFailedPerSec gets the value of HttpWebRequestsFailedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsFailedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsFailedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpWebRequestsQueuedPerSec sets the value of HttpWebRequestsQueuedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) SetPropertyHttpWebRequestsQueuedPerSec(value uint32) (err error) {
	return instance.SetProperty("HttpWebRequestsQueuedPerSec", value)
}

// GetHttpWebRequestsQueuedPerSec gets the value of HttpWebRequestsQueuedPerSec for the instance
func (instance *Win32_PerfRawData_NETCLRNetworking4000_NETCLRNetworking4000) GetPropertyHttpWebRequestsQueuedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpWebRequestsQueuedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
