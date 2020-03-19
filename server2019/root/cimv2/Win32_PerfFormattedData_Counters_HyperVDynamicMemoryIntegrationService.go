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

// Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService struct
type Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService struct {
	*Win32_PerfFormattedData

	//
	MaximumMemoryMbytes uint64
}

func NewWin32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationServiceEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetMaximumMemoryMbytes sets the value of MaximumMemoryMbytes for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService) SetPropertyMaximumMemoryMbytes(value uint64) (err error) {
	return instance.SetProperty("MaximumMemoryMbytes", value)
}

// GetMaximumMemoryMbytes gets the value of MaximumMemoryMbytes for the instance
func (instance *Win32_PerfFormattedData_Counters_HyperVDynamicMemoryIntegrationService) GetPropertyMaximumMemoryMbytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumMemoryMbytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
