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

// Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic struct
type Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic struct {
	*Win32_PerfRawData

	//
	TotalInboundBytesDropped uint64

	//
	TotalInboundPacketsDropped uint64
}

func NewWin32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTrafficEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTrafficEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetTotalInboundBytesDropped sets the value of TotalInboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic) SetPropertyTotalInboundBytesDropped(value uint64) (err error) {
	return instance.SetProperty("TotalInboundBytesDropped", value)
}

// GetTotalInboundBytesDropped gets the value of TotalInboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic) GetPropertyTotalInboundBytesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInboundBytesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalInboundPacketsDropped sets the value of TotalInboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic) SetPropertyTotalInboundPacketsDropped(value uint64) (err error) {
	return instance.SetProperty("TotalInboundPacketsDropped", value)
}

// GetTotalInboundPacketsDropped gets the value of TotalInboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalInboundNetworkTraffic) GetPropertyTotalInboundPacketsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInboundPacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}