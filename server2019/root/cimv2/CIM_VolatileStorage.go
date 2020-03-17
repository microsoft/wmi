// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_VolatileStorage struct
type CIM_VolatileStorage struct {
	CIM_Memory

	//
	Cacheable bool

	//
	CacheType uint16
}

// SetCacheable sets the value of Cacheable for the instance
func (instance *CIM_VolatileStorage) SetPropertyCacheable(value bool) (err error) {
	return instance.SetProperty("Cacheable", value)
}

// GetCacheable gets the value of Cacheable for the instance
func (instance *CIM_VolatileStorage) GetPropertyCacheable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cacheable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheType sets the value of CacheType for the instance
func (instance *CIM_VolatileStorage) SetPropertyCacheType(value uint16) (err error) {
	return instance.SetProperty("CacheType", value)
}

// GetCacheType gets the value of CacheType for the instance
func (instance *CIM_VolatileStorage) GetPropertyCacheType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CacheType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
