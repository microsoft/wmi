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

// Win32_PerfRawData_Counters_WFP struct
type Win32_PerfRawData_Counters_WFP struct {
	*Win32_PerfRawData

	//
	ProviderCount uint32
}

func NewWin32_PerfRawData_Counters_WFPEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_WFP, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_WFP{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_WFPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_WFP, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_WFP{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetProviderCount sets the value of ProviderCount for the instance
func (instance *Win32_PerfRawData_Counters_WFP) SetPropertyProviderCount(value uint32) (err error) {
	return instance.SetProperty("ProviderCount", value)
}

// GetProviderCount gets the value of ProviderCount for the instance
func (instance *Win32_PerfRawData_Counters_WFP) GetPropertyProviderCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProviderCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
