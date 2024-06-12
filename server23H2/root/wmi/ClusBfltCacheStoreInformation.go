// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// ClusBfltCacheStoreInformation struct
type ClusBfltCacheStoreInformation struct {
	*cim.WmiInstance

	// Device Guid.
	DeviceGuid string

	// Id.
	Id string

	// Key.
	key string

	// Page Size.
	PageSize uint32

	// PathId.
	PathId uint32

	// Status.
	Status uint32

	// Store Size.
	StoreSize uint64
}

func NewClusBfltCacheStoreInformationEx1(instance *cim.WmiInstance) (newInstance *ClusBfltCacheStoreInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltCacheStoreInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltCacheStoreInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltCacheStoreInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltCacheStoreInformation{
		WmiInstance: tmp,
	}
	return
}

// SetDeviceGuid sets the value of DeviceGuid for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertyDeviceGuid(value string) (err error) {
	return instance.SetProperty("DeviceGuid", (value))
}

// GetDeviceGuid gets the value of DeviceGuid for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyDeviceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceGuid")
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

// SetId sets the value of Id for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// Setkey sets the value of key for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
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
func (instance *ClusBfltCacheStoreInformation) SetPropertyPageSize(value uint32) (err error) {
	return instance.SetProperty("PageSize", (value))
}

// GetPageSize gets the value of PageSize for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyPageSize() (value uint32, err error) {
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

// SetPathId sets the value of PathId for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertyPathId(value uint32) (err error) {
	return instance.SetProperty("PathId", (value))
}

// GetPathId gets the value of PathId for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyPathId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PathId")
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

// SetStatus sets the value of Status for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStoreSize sets the value of StoreSize for the instance
func (instance *ClusBfltCacheStoreInformation) SetPropertyStoreSize(value uint64) (err error) {
	return instance.SetProperty("StoreSize", (value))
}

// GetStoreSize gets the value of StoreSize for the instance
func (instance *ClusBfltCacheStoreInformation) GetPropertyStoreSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("StoreSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
