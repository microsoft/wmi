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

// Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode struct
type Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode struct {
	*Win32_PerfFormattedData

	//
	PageCount uint64

	//
	ProcessorCount uint64
}

func NewWin32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNodeEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetPageCount sets the value of PageCount for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode) SetPropertyPageCount(value uint64) (err error) {
	return instance.SetProperty("PageCount", value)
}

// GetPageCount gets the value of PageCount for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode) GetPropertyPageCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorCount sets the value of ProcessorCount for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode) SetPropertyProcessorCount(value uint64) (err error) {
	return instance.SetProperty("ProcessorCount", value)
}

// GetProcessorCount gets the value of ProcessorCount for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidNumaNode) GetPropertyProcessorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
