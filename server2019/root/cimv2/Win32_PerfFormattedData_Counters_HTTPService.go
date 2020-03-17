// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_HTTPService struct
type Win32_PerfFormattedData_Counters_HTTPService struct {
	Win32_PerfFormattedData

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
