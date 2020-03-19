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

// Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2 struct
type Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2 struct {
	*Win32_PerfRawData

	//
	UnacknowledgedMessageCount uint64
}

func NewWin32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetUnacknowledgedMessageCount sets the value of UnacknowledgedMessageCount for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2) SetPropertyUnacknowledgedMessageCount(value uint64) (err error) {
	return instance.SetProperty("UnacknowledgedMessageCount", value)
}

// GetUnacknowledgedMessageCount gets the value of UnacknowledgedMessageCount for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2) GetPropertyUnacknowledgedMessageCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnacknowledgedMessageCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
