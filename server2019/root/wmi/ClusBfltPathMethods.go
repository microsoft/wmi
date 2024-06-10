// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClusBfltPathMethods struct
type ClusBfltPathMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewClusBfltPathMethodsEx1(instance *cim.WmiInstance) (newInstance *ClusBfltPathMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltPathMethods{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltPathMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltPathMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltPathMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ClusBfltPathMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ClusBfltPathMethods) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ClusBfltPathMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ClusBfltPathMethods) GetPropertyInstanceName() (value string, err error) {
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

// Reinitializes disk partition with GPT, trims SSD disk

// <param name="Flags" type="uint32 ">Flags</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) ReInitializeDisk( /* IN */ PathId uint32,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("ReInitializeDisk", PathId, Flags)
	if err != nil {
		return
	}
	return

}

// Creates Cache Store on SSD

// <param name="PageSize" type="uint32 ">Page Size</param>
// <param name="PathId" type="uint32 ">Path Id</param>
// <param name="ReservePercentage" type="uint32 ">Reserve Percentage</param>
// <param name="UnusedSize" type="uint64 ">Unused Size</param>
func (instance *ClusBfltPathMethods) CreateSsdCacheStore( /* IN */ PathId uint32,
	/* IN */ ReservePercentage uint32,
	/* IN */ UnusedSize uint64,
	/* IN */ PageSize uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("CreateSsdCacheStore", PathId, ReservePercentage, UnusedSize, PageSize)
	if err != nil {
		return
	}
	return

}

// Prepares HDD Disk

// <param name="Options" type="uint32 ">Options</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) PrepareHddForCache( /* IN */ PathId uint32,
	/* IN */ Options uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("PrepareHddForCache", PathId, Options)
	if err != nil {
		return
	}
	return

}

// Binds Hdd to Cache Store

// <param name="Attributes" type="uint32 ">Attributes</param>
// <param name="CacheStoreId" type="string ">CacheStoreId</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) BindHddToCacheStore( /* IN */ PathId uint32,
	/* IN */ Attributes uint32,
	/* IN */ CacheStoreId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("BindHddToCacheStore", PathId, Attributes, CacheStoreId)
	if err != nil {
		return
	}
	return

}

// Unbinds Hdd from Cache Store

// <param name="Flags" type="uint32 ">Flags</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) UnBindHdd( /* IN */ PathId uint32,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("UnBindHdd", PathId, Flags)
	if err != nil {
		return
	}
	return

}

// Queries Hdd Binding

// <param name="PathId" type="uint32 ">Path Id</param>

// <param name="BindingInfo" type="ClusBfltHddBindingInformation ">BindingInfo</param>
func (instance *ClusBfltPathMethods) QueryHddBinding( /* IN */ PathId uint32,
	/* OUT */ BindingInfo ClusBfltHddBindingInformation) (err error) {
	_, err = instance.InvokeMethod("QueryHddBinding", PathId)
	if err != nil {
		return
	}
	return

}

// Queries Cache Stores

// <param name="PathId" type="uint32 ">Path Id</param>

// <param name="CacheStores" type="ClusBfltCacheStoresInformation ">CacheStores</param>
func (instance *ClusBfltPathMethods) QueryCacheStores( /* IN */ PathId uint32,
	/* OUT */ CacheStores ClusBfltCacheStoresInformation) (err error) {
	_, err = instance.InvokeMethod("QueryCacheStores", PathId)
	if err != nil {
		return
	}
	return

}

// Queries Cache Store Binding Records

// <param name="CacheStoreId" type="string ">Cache Store Id</param>
// <param name="PathId" type="uint32 ">Path Id</param>

// <param name="BindingRecords" type="ClusBfltSsdBindingRecords ">BindingRecords</param>
func (instance *ClusBfltPathMethods) QuerySsdBindingRecords( /* IN */ PathId uint32,
	/* IN */ CacheStoreId string,
	/* OUT */ BindingRecords ClusBfltSsdBindingRecords) (err error) {
	_, err = instance.InvokeMethod("QuerySsdBindingRecords", PathId, CacheStoreId)
	if err != nil {
		return
	}
	return

}

// Disabled Cache Store Binding

// <param name="BindingId" type="string ">Binding Id</param>
// <param name="CacheStoreId" type="string ">Cache Store Id</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) DisableSsdBinding( /* IN */ PathId uint32,
	/* IN */ CacheStoreId string,
	/* IN */ BindingId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableSsdBinding", PathId, CacheStoreId, BindingId)
	if err != nil {
		return
	}
	return

}

// Gets Path Id of a Given Type by Device Guid

// <param name="Attributes" type="uint32 ">Attributes</param>
// <param name="AttributesMask" type="uint32 ">AttributesMask</param>
// <param name="DeviceGuid" type="string ">Device Id</param>
// <param name="PathType" type="uint32 ">PathType</param>
// <param name="Timeout" type="uint32 ">Timeout</param>

// <param name="PathId" type="uint32 ">PathId</param>
func (instance *ClusBfltPathMethods) GetPathIdByDeviceGuid( /* IN */ DeviceGuid string,
	/* IN */ PathType uint32,
	/* IN */ Attributes uint32,
	/* IN */ AttributesMask uint32,
	/* IN */ Timeout uint32,
	/* OUT */ PathId uint32) (err error) {
	_, err = instance.InvokeMethod("GetPathIdByDeviceGuid", DeviceGuid, PathType, Attributes, AttributesMask, Timeout)
	if err != nil {
		return
	}
	return

}

// Trims the disk

// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) TrimDisk( /* IN */ PathId uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("TrimDisk", PathId)
	if err != nil {
		return
	}
	return

}

// Disabled Cache Store Binding

// <param name="Attributes" type="uint32 ">Attributes</param>
// <param name="AttributesMask" type="uint32 ">AttributesMask</param>
// <param name="BindingId" type="string ">Binding Id</param>
// <param name="CacheStoreId" type="string ">Cache Store Id</param>
// <param name="PathId" type="uint32 ">Path Id</param>
func (instance *ClusBfltPathMethods) SetSsdBindingAttributes( /* IN */ PathId uint32,
	/* IN */ Attributes uint32,
	/* IN */ AttributesMask uint32,
	/* IN */ CacheStoreId string,
	/* IN */ BindingId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetSsdBindingAttributes", PathId, Attributes, AttributesMask, CacheStoreId, BindingId)
	if err != nil {
		return
	}
	return

}
