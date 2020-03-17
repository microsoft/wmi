// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages struct
type Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages struct {
	Win32_PerfRawData

	//
	MessageQueueLength uint64

	//
	UnacknowledgedMessages uint64
}

// SetMessageQueueLength sets the value of MessageQueueLength for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages) SetPropertyMessageQueueLength(value uint64) (err error) {
	return instance.SetProperty("MessageQueueLength", value)
}

// GetMessageQueueLength gets the value of MessageQueueLength for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages) GetPropertyMessageQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnacknowledgedMessages sets the value of UnacknowledgedMessages for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages) SetPropertyUnacknowledgedMessages(value uint64) (err error) {
	return instance.SetProperty("UnacknowledgedMessages", value)
}

// GetUnacknowledgedMessages gets the value of UnacknowledgedMessages for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages) GetPropertyUnacknowledgedMessages() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnacknowledgedMessages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
