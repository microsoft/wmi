// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CacheSection struct
type CacheSection struct {
	*ConfigurationSection

	//
	DisableExpiration bool

	//
	DisableMemoryCollection bool

	//
	PercentagePhysicalMemoryUsedLimit int32

	//
	PrivateBytesLimit string

	//
	PrivateBytesPollTime string
}

func NewCacheSectionEx1(instance *cim.WmiInstance) (newInstance *CacheSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CacheSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewCacheSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CacheSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CacheSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDisableExpiration sets the value of DisableExpiration for the instance
func (instance *CacheSection) SetPropertyDisableExpiration(value bool) (err error) {
	return instance.SetProperty("DisableExpiration", (value))
}

// GetDisableExpiration gets the value of DisableExpiration for the instance
func (instance *CacheSection) GetPropertyDisableExpiration() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableExpiration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDisableMemoryCollection sets the value of DisableMemoryCollection for the instance
func (instance *CacheSection) SetPropertyDisableMemoryCollection(value bool) (err error) {
	return instance.SetProperty("DisableMemoryCollection", (value))
}

// GetDisableMemoryCollection gets the value of DisableMemoryCollection for the instance
func (instance *CacheSection) GetPropertyDisableMemoryCollection() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableMemoryCollection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetPercentagePhysicalMemoryUsedLimit sets the value of PercentagePhysicalMemoryUsedLimit for the instance
func (instance *CacheSection) SetPropertyPercentagePhysicalMemoryUsedLimit(value int32) (err error) {
	return instance.SetProperty("PercentagePhysicalMemoryUsedLimit", (value))
}

// GetPercentagePhysicalMemoryUsedLimit gets the value of PercentagePhysicalMemoryUsedLimit for the instance
func (instance *CacheSection) GetPropertyPercentagePhysicalMemoryUsedLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("PercentagePhysicalMemoryUsedLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetPrivateBytesLimit sets the value of PrivateBytesLimit for the instance
func (instance *CacheSection) SetPropertyPrivateBytesLimit(value string) (err error) {
	return instance.SetProperty("PrivateBytesLimit", (value))
}

// GetPrivateBytesLimit gets the value of PrivateBytesLimit for the instance
func (instance *CacheSection) GetPropertyPrivateBytesLimit() (value string, err error) {
	retValue, err := instance.GetProperty("PrivateBytesLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPrivateBytesPollTime sets the value of PrivateBytesPollTime for the instance
func (instance *CacheSection) SetPropertyPrivateBytesPollTime(value string) (err error) {
	return instance.SetProperty("PrivateBytesPollTime", (value))
}

// GetPrivateBytesPollTime gets the value of PrivateBytesPollTime for the instance
func (instance *CacheSection) GetPropertyPrivateBytesPollTime() (value string, err error) {
	retValue, err := instance.GetProperty("PrivateBytesPollTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
