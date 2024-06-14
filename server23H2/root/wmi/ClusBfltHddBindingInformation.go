// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClusBfltHddBindingInformation struct
type ClusBfltHddBindingInformation struct {
	*cim.WmiInstance

	// Binding Id.
	BindingId string

	// Device Id.
	CacheStoreDeviceId string

	// Id.
	CacheStoreId string

	// Key.
	CacheStoreKey string

	// Page Size.
	PageSize uint32
}

func NewClusBfltHddBindingInformationEx1(instance *cim.WmiInstance) (newInstance *ClusBfltHddBindingInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltHddBindingInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltHddBindingInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltHddBindingInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltHddBindingInformation{
		WmiInstance: tmp,
	}
	return
}

// SetBindingId sets the value of BindingId for the instance
func (instance *ClusBfltHddBindingInformation) SetPropertyBindingId(value string) (err error) {
	return instance.SetProperty("BindingId", (value))
}

// GetBindingId gets the value of BindingId for the instance
func (instance *ClusBfltHddBindingInformation) GetPropertyBindingId() (value string, err error) {
	retValue, err := instance.GetProperty("BindingId")
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

// SetCacheStoreDeviceId sets the value of CacheStoreDeviceId for the instance
func (instance *ClusBfltHddBindingInformation) SetPropertyCacheStoreDeviceId(value string) (err error) {
	return instance.SetProperty("CacheStoreDeviceId", (value))
}

// GetCacheStoreDeviceId gets the value of CacheStoreDeviceId for the instance
func (instance *ClusBfltHddBindingInformation) GetPropertyCacheStoreDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("CacheStoreDeviceId")
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

// SetCacheStoreId sets the value of CacheStoreId for the instance
func (instance *ClusBfltHddBindingInformation) SetPropertyCacheStoreId(value string) (err error) {
	return instance.SetProperty("CacheStoreId", (value))
}

// GetCacheStoreId gets the value of CacheStoreId for the instance
func (instance *ClusBfltHddBindingInformation) GetPropertyCacheStoreId() (value string, err error) {
	retValue, err := instance.GetProperty("CacheStoreId")
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

// SetCacheStoreKey sets the value of CacheStoreKey for the instance
func (instance *ClusBfltHddBindingInformation) SetPropertyCacheStoreKey(value string) (err error) {
	return instance.SetProperty("CacheStoreKey", (value))
}

// GetCacheStoreKey gets the value of CacheStoreKey for the instance
func (instance *ClusBfltHddBindingInformation) GetPropertyCacheStoreKey() (value string, err error) {
	retValue, err := instance.GetProperty("CacheStoreKey")
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

// SetPageSize sets the value of PageSize for the instance
func (instance *ClusBfltHddBindingInformation) SetPropertyPageSize(value uint32) (err error) {
	return instance.SetProperty("PageSize", (value))
}

// GetPageSize gets the value of PageSize for the instance
func (instance *ClusBfltHddBindingInformation) GetPropertyPageSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
