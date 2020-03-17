// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40 struct
type Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40 struct {
	Win32_PerfRawData

	//
	CacheEntries uint32

	//
	CacheHitRatio uint32

	//
	CacheHitRatio_Base uint32

	//
	CacheHits uint32

	//
	CacheMisses uint32

	//
	CacheTrims uint32

	//
	CacheTurnoverRate uint32
}

// SetCacheEntries sets the value of CacheEntries for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheEntries(value uint32) (err error) {
	return instance.SetProperty("CacheEntries", value)
}

// GetCacheEntries gets the value of CacheEntries for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheEntries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheHitRatio sets the value of CacheHitRatio for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheHitRatio(value uint32) (err error) {
	return instance.SetProperty("CacheHitRatio", value)
}

// GetCacheHitRatio gets the value of CacheHitRatio for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheHitRatio() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheHitRatio")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheHitRatio_Base sets the value of CacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheHitRatio_Base(value uint32) (err error) {
	return instance.SetProperty("CacheHitRatio_Base", value)
}

// GetCacheHitRatio_Base gets the value of CacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheHitRatio_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheHitRatio_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheHits sets the value of CacheHits for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheHits(value uint32) (err error) {
	return instance.SetProperty("CacheHits", value)
}

// GetCacheHits gets the value of CacheHits for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheHits")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheMisses sets the value of CacheMisses for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheMisses(value uint32) (err error) {
	return instance.SetProperty("CacheMisses", value)
}

// GetCacheMisses gets the value of CacheMisses for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheMisses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheTrims sets the value of CacheTrims for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheTrims(value uint32) (err error) {
	return instance.SetProperty("CacheTrims", value)
}

// GetCacheTrims gets the value of CacheTrims for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheTrims() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheTrims")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheTurnoverRate sets the value of CacheTurnoverRate for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) SetPropertyCacheTurnoverRate(value uint32) (err error) {
	return instance.SetProperty("CacheTurnoverRate", value)
}

// GetCacheTurnoverRate gets the value of CacheTurnoverRate for the instance
func (instance *Win32_PerfRawData_NETMemoryCache40_NETMemoryCache40) GetPropertyCacheTurnoverRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheTurnoverRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
