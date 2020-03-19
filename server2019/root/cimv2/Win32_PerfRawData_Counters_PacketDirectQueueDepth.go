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

// Win32_PerfRawData_Counters_PacketDirectQueueDepth struct
type Win32_PerfRawData_Counters_PacketDirectQueueDepth struct {
	*Win32_PerfRawData

	//
	AverageQueueDepth uint32

	//
	PercentAverageQueueUtilization uint32
}

func NewWin32_PerfRawData_Counters_PacketDirectQueueDepthEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_PacketDirectQueueDepth, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_PacketDirectQueueDepth{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_PacketDirectQueueDepthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_PacketDirectQueueDepth, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_PacketDirectQueueDepth{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAverageQueueDepth sets the value of AverageQueueDepth for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) SetPropertyAverageQueueDepth(value uint32) (err error) {
	return instance.SetProperty("AverageQueueDepth", value)
}

// GetAverageQueueDepth gets the value of AverageQueueDepth for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) GetPropertyAverageQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("AverageQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentAverageQueueUtilization sets the value of PercentAverageQueueUtilization for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) SetPropertyPercentAverageQueueUtilization(value uint32) (err error) {
	return instance.SetProperty("PercentAverageQueueUtilization", value)
}

// GetPercentAverageQueueUtilization gets the value of PercentAverageQueueUtilization for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) GetPropertyPercentAverageQueueUtilization() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentAverageQueueUtilization")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
