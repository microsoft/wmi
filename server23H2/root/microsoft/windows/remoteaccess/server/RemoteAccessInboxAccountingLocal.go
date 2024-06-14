// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessInboxAccountingLocal struct
type RemoteAccessInboxAccountingLocal struct {
	*cim.WmiInstance

	//
	AccountingStatus bool

	//
	StoreFirstRecordDate string

	//
	StoreFreeBytes uint64

	//
	StoreLastRecordDate string

	//
	StoreLimit uint32

	//
	StoreLimitUnits uint32

	//
	StoreUsedBytes uint64
}

func NewRemoteAccessInboxAccountingLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessInboxAccountingLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessInboxAccountingLocal{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessInboxAccountingLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessInboxAccountingLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessInboxAccountingLocal{
		WmiInstance: tmp,
	}
	return
}

// SetAccountingStatus sets the value of AccountingStatus for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyAccountingStatus(value bool) (err error) {
	return instance.SetProperty("AccountingStatus", (value))
}

// GetAccountingStatus gets the value of AccountingStatus for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyAccountingStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("AccountingStatus")
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

// SetStoreFirstRecordDate sets the value of StoreFirstRecordDate for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreFirstRecordDate(value string) (err error) {
	return instance.SetProperty("StoreFirstRecordDate", (value))
}

// GetStoreFirstRecordDate gets the value of StoreFirstRecordDate for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreFirstRecordDate() (value string, err error) {
	retValue, err := instance.GetProperty("StoreFirstRecordDate")
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

// SetStoreFreeBytes sets the value of StoreFreeBytes for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreFreeBytes(value uint64) (err error) {
	return instance.SetProperty("StoreFreeBytes", (value))
}

// GetStoreFreeBytes gets the value of StoreFreeBytes for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreFreeBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("StoreFreeBytes")
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

// SetStoreLastRecordDate sets the value of StoreLastRecordDate for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreLastRecordDate(value string) (err error) {
	return instance.SetProperty("StoreLastRecordDate", (value))
}

// GetStoreLastRecordDate gets the value of StoreLastRecordDate for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreLastRecordDate() (value string, err error) {
	retValue, err := instance.GetProperty("StoreLastRecordDate")
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

// SetStoreLimit sets the value of StoreLimit for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreLimit(value uint32) (err error) {
	return instance.SetProperty("StoreLimit", (value))
}

// GetStoreLimit gets the value of StoreLimit for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("StoreLimit")
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

// SetStoreLimitUnits sets the value of StoreLimitUnits for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreLimitUnits(value uint32) (err error) {
	return instance.SetProperty("StoreLimitUnits", (value))
}

// GetStoreLimitUnits gets the value of StoreLimitUnits for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreLimitUnits() (value uint32, err error) {
	retValue, err := instance.GetProperty("StoreLimitUnits")
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

// SetStoreUsedBytes sets the value of StoreUsedBytes for the instance
func (instance *RemoteAccessInboxAccountingLocal) SetPropertyStoreUsedBytes(value uint64) (err error) {
	return instance.SetProperty("StoreUsedBytes", (value))
}

// GetStoreUsedBytes gets the value of StoreUsedBytes for the instance
func (instance *RemoteAccessInboxAccountingLocal) GetPropertyStoreUsedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("StoreUsedBytes")
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
