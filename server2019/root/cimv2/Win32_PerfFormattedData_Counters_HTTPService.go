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

// Win32_PerfFormattedData_Counters_HTTPService struct
type Win32_PerfFormattedData_Counters_HTTPService struct {
	*Win32_PerfFormattedData

	//
	CurrentUrisCached uint32

	//
	TotalFlushedUris uint32

	//
	TotalUrisCached uint32

	//
	UriCacheFlushes uint32

	//
	UriCacheHits uint32

	//
	UriCacheMisses uint32
}

func NewWin32_PerfFormattedData_Counters_HTTPServiceEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_HTTPService, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_HTTPService{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_HTTPServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_HTTPService, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_HTTPService{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCurrentUrisCached sets the value of CurrentUrisCached for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyCurrentUrisCached(value uint32) (err error) {
	return instance.SetProperty("CurrentUrisCached", value)
}

// GetCurrentUrisCached gets the value of CurrentUrisCached for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyCurrentUrisCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentUrisCached")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalFlushedUris sets the value of TotalFlushedUris for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyTotalFlushedUris(value uint32) (err error) {
	return instance.SetProperty("TotalFlushedUris", value)
}

// GetTotalFlushedUris gets the value of TotalFlushedUris for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyTotalFlushedUris() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalFlushedUris")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalUrisCached sets the value of TotalUrisCached for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyTotalUrisCached(value uint32) (err error) {
	return instance.SetProperty("TotalUrisCached", value)
}

// GetTotalUrisCached gets the value of TotalUrisCached for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyTotalUrisCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalUrisCached")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUriCacheFlushes sets the value of UriCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyUriCacheFlushes(value uint32) (err error) {
	return instance.SetProperty("UriCacheFlushes", value)
}

// GetUriCacheFlushes gets the value of UriCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyUriCacheFlushes() (value uint32, err error) {
	retValue, err := instance.GetProperty("UriCacheFlushes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUriCacheHits sets the value of UriCacheHits for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyUriCacheHits(value uint32) (err error) {
	return instance.SetProperty("UriCacheHits", value)
}

// GetUriCacheHits gets the value of UriCacheHits for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyUriCacheHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("UriCacheHits")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUriCacheMisses sets the value of UriCacheMisses for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) SetPropertyUriCacheMisses(value uint32) (err error) {
	return instance.SetProperty("UriCacheMisses", value)
}

// GetUriCacheMisses gets the value of UriCacheMisses for the instance
func (instance *Win32_PerfFormattedData_Counters_HTTPService) GetPropertyUriCacheMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("UriCacheMisses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}