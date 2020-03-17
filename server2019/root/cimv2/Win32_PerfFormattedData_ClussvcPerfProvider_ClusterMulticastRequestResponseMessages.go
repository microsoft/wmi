// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages struct {
	Win32_PerfFormattedData

	//
	MessagesOutstanding uint64

	//
	MessagesSent uint64

	//
	MessagesSentPersec uint64
}

// SetMessagesOutstanding sets the value of MessagesOutstanding for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesOutstanding(value uint64) (err error) {
	return instance.SetProperty("MessagesOutstanding", value)
}

// GetMessagesOutstanding gets the value of MessagesOutstanding for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesOutstanding() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesOutstanding")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessagesSent sets the value of MessagesSent for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesSent(value uint64) (err error) {
	return instance.SetProperty("MessagesSent", value)
}

// GetMessagesSent gets the value of MessagesSent for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessagesSentPersec sets the value of MessagesSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesSentPersec(value uint64) (err error) {
	return instance.SetProperty("MessagesSentPersec", value)
}

// GetMessagesSentPersec gets the value of MessagesSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
