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

// ClusBfltCacheStoresInformation struct
type ClusBfltCacheStoresInformation struct {
	*cim.WmiInstance

	//
	Active bool

	// Cache Store Info.
	CacheStoreInfo []ClusBfltCacheStoreInformation

	//
	InstanceName string

	// Number of Cache Stores.
	NumberOfCacheStores uint32
}

func NewClusBfltCacheStoresInformationEx1(instance *cim.WmiInstance) (newInstance *ClusBfltCacheStoresInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltCacheStoresInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltCacheStoresInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltCacheStoresInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltCacheStoresInformation{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ClusBfltCacheStoresInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ClusBfltCacheStoresInformation) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetCacheStoreInfo sets the value of CacheStoreInfo for the instance
func (instance *ClusBfltCacheStoresInformation) SetPropertyCacheStoreInfo(value []ClusBfltCacheStoreInformation) (err error) {
	return instance.SetProperty("CacheStoreInfo", (value))
}

// GetCacheStoreInfo gets the value of CacheStoreInfo for the instance
func (instance *ClusBfltCacheStoresInformation) GetPropertyCacheStoreInfo() (value []ClusBfltCacheStoreInformation, err error) {
	retValue, err := instance.GetProperty("CacheStoreInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ClusBfltCacheStoreInformation)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ClusBfltCacheStoreInformation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ClusBfltCacheStoreInformation(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ClusBfltCacheStoresInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ClusBfltCacheStoresInformation) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetNumberOfCacheStores sets the value of NumberOfCacheStores for the instance
func (instance *ClusBfltCacheStoresInformation) SetPropertyNumberOfCacheStores(value uint32) (err error) {
	return instance.SetProperty("NumberOfCacheStores", (value))
}

// GetNumberOfCacheStores gets the value of NumberOfCacheStores for the instance
func (instance *ClusBfltCacheStoresInformation) GetPropertyNumberOfCacheStores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfCacheStores")
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
