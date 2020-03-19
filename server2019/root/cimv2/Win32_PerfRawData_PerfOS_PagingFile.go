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

// Win32_PerfRawData_PerfOS_PagingFile struct
type Win32_PerfRawData_PerfOS_PagingFile struct {
	*Win32_PerfRawData

	//
	PercentUsage uint32

	//
	PercentUsage_Base uint32

	//
	PercentUsagePeak uint32

	//
	PercentUsagePeak_Base uint32
}

func NewWin32_PerfRawData_PerfOS_PagingFileEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_PerfOS_PagingFile, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_PerfOS_PagingFile{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_PerfOS_PagingFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_PerfOS_PagingFile, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_PerfOS_PagingFile{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetPercentUsage sets the value of PercentUsage for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) SetPropertyPercentUsage(value uint32) (err error) {
	return instance.SetProperty("PercentUsage", value)
}

// GetPercentUsage gets the value of PercentUsage for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) GetPropertyPercentUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUsage_Base sets the value of PercentUsage_Base for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) SetPropertyPercentUsage_Base(value uint32) (err error) {
	return instance.SetProperty("PercentUsage_Base", value)
}

// GetPercentUsage_Base gets the value of PercentUsage_Base for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) GetPropertyPercentUsage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsage_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUsagePeak sets the value of PercentUsagePeak for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) SetPropertyPercentUsagePeak(value uint32) (err error) {
	return instance.SetProperty("PercentUsagePeak", value)
}

// GetPercentUsagePeak gets the value of PercentUsagePeak for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) GetPropertyPercentUsagePeak() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsagePeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUsagePeak_Base sets the value of PercentUsagePeak_Base for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) SetPropertyPercentUsagePeak_Base(value uint32) (err error) {
	return instance.SetProperty("PercentUsagePeak_Base", value)
}

// GetPercentUsagePeak_Base gets the value of PercentUsagePeak_Base for the instance
func (instance *Win32_PerfRawData_PerfOS_PagingFile) GetPropertyPercentUsagePeak_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsagePeak_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
