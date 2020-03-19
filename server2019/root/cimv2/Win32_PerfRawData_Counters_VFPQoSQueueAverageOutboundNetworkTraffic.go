// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic struct
type Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic struct {
	*Win32_PerfRawData

	//
	AverageOutboundBytesAllowedThroughtheQueue uint64

	//
	AverageOutboundBytesDropped uint64

	//
	AverageOutboundBytesEnteringtheQueue uint64

	//
	AverageOutboundBytesQueuedduetoBacklog uint64

	//
	AverageOutboundBytesQueuedduetoInsufficientTokens uint64

	//
	AverageOutboundBytesResumed uint64

	//
	AverageOutboundPacketsAllowedThroughtheQueue uint64

	//
	AverageOutboundPacketsDropped uint64

	//
	AverageOutboundPacketsEnteringtheQueue uint64

	//
	AverageOutboundPacketsQueuedduetoBacklog uint64

	//
	AverageOutboundPacketsQueuedduetoInsufficientTokens uint64

	//
	AverageOutboundPacketsResumed uint64
}

func NewWin32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTrafficEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTrafficEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAverageOutboundBytesAllowedThroughtheQueue sets the value of AverageOutboundBytesAllowedThroughtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesAllowedThroughtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesAllowedThroughtheQueue", value)
}

// GetAverageOutboundBytesAllowedThroughtheQueue gets the value of AverageOutboundBytesAllowedThroughtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesAllowedThroughtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesAllowedThroughtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundBytesDropped sets the value of AverageOutboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesDropped(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesDropped", value)
}

// GetAverageOutboundBytesDropped gets the value of AverageOutboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundBytesEnteringtheQueue sets the value of AverageOutboundBytesEnteringtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesEnteringtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesEnteringtheQueue", value)
}

// GetAverageOutboundBytesEnteringtheQueue gets the value of AverageOutboundBytesEnteringtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesEnteringtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesEnteringtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundBytesQueuedduetoBacklog sets the value of AverageOutboundBytesQueuedduetoBacklog for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesQueuedduetoBacklog(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesQueuedduetoBacklog", value)
}

// GetAverageOutboundBytesQueuedduetoBacklog gets the value of AverageOutboundBytesQueuedduetoBacklog for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesQueuedduetoBacklog() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesQueuedduetoBacklog")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundBytesQueuedduetoInsufficientTokens sets the value of AverageOutboundBytesQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesQueuedduetoInsufficientTokens(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesQueuedduetoInsufficientTokens", value)
}

// GetAverageOutboundBytesQueuedduetoInsufficientTokens gets the value of AverageOutboundBytesQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesQueuedduetoInsufficientTokens() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesQueuedduetoInsufficientTokens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundBytesResumed sets the value of AverageOutboundBytesResumed for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundBytesResumed(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundBytesResumed", value)
}

// GetAverageOutboundBytesResumed gets the value of AverageOutboundBytesResumed for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundBytesResumed() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundBytesResumed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsAllowedThroughtheQueue sets the value of AverageOutboundPacketsAllowedThroughtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsAllowedThroughtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsAllowedThroughtheQueue", value)
}

// GetAverageOutboundPacketsAllowedThroughtheQueue gets the value of AverageOutboundPacketsAllowedThroughtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsAllowedThroughtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsAllowedThroughtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsDropped sets the value of AverageOutboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsDropped(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsDropped", value)
}

// GetAverageOutboundPacketsDropped gets the value of AverageOutboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsEnteringtheQueue sets the value of AverageOutboundPacketsEnteringtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsEnteringtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsEnteringtheQueue", value)
}

// GetAverageOutboundPacketsEnteringtheQueue gets the value of AverageOutboundPacketsEnteringtheQueue for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsEnteringtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsEnteringtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsQueuedduetoBacklog sets the value of AverageOutboundPacketsQueuedduetoBacklog for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsQueuedduetoBacklog(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsQueuedduetoBacklog", value)
}

// GetAverageOutboundPacketsQueuedduetoBacklog gets the value of AverageOutboundPacketsQueuedduetoBacklog for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsQueuedduetoBacklog() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsQueuedduetoBacklog")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsQueuedduetoInsufficientTokens sets the value of AverageOutboundPacketsQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsQueuedduetoInsufficientTokens(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsQueuedduetoInsufficientTokens", value)
}

// GetAverageOutboundPacketsQueuedduetoInsufficientTokens gets the value of AverageOutboundPacketsQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsQueuedduetoInsufficientTokens() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsQueuedduetoInsufficientTokens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageOutboundPacketsResumed sets the value of AverageOutboundPacketsResumed for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) SetPropertyAverageOutboundPacketsResumed(value uint64) (err error) {
	return instance.SetProperty("AverageOutboundPacketsResumed", value)
}

// GetAverageOutboundPacketsResumed gets the value of AverageOutboundPacketsResumed for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueAverageOutboundNetworkTraffic) GetPropertyAverageOutboundPacketsResumed() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageOutboundPacketsResumed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
