// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessInboxAccounting struct
type RemoteAccessInboxAccounting struct {
	*cim.WmiInstance

	//
	InboxAccountingStatus string

	//
	StoreFirstRecordDate string

	//
	StoreFreeBytes uint64

	//
	StoreFreeBytesInPercentage float32

	//
	StoreLastRecordDate string

	//
	StoreLimit string

	//
	StoreUsedBytes uint64

	//
	StoreUsedBytesInPercentage float32
}

func NewRemoteAccessInboxAccountingEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessInboxAccounting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessInboxAccounting{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessInboxAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessInboxAccounting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessInboxAccounting{
		WmiInstance: tmp,
	}
	return
}

// SetInboxAccountingStatus sets the value of InboxAccountingStatus for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyInboxAccountingStatus(value string) (err error) {
	return instance.SetProperty("InboxAccountingStatus", (value))
}

// GetInboxAccountingStatus gets the value of InboxAccountingStatus for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyInboxAccountingStatus() (value string, err error) {
	retValue, err := instance.GetProperty("InboxAccountingStatus")
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

// SetStoreFirstRecordDate sets the value of StoreFirstRecordDate for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreFirstRecordDate(value string) (err error) {
	return instance.SetProperty("StoreFirstRecordDate", (value))
}

// GetStoreFirstRecordDate gets the value of StoreFirstRecordDate for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreFirstRecordDate() (value string, err error) {
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
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreFreeBytes(value uint64) (err error) {
	return instance.SetProperty("StoreFreeBytes", (value))
}

// GetStoreFreeBytes gets the value of StoreFreeBytes for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreFreeBytes() (value uint64, err error) {
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

// SetStoreFreeBytesInPercentage sets the value of StoreFreeBytesInPercentage for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreFreeBytesInPercentage(value float32) (err error) {
	return instance.SetProperty("StoreFreeBytesInPercentage", (value))
}

// GetStoreFreeBytesInPercentage gets the value of StoreFreeBytesInPercentage for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreFreeBytesInPercentage() (value float32, err error) {
	retValue, err := instance.GetProperty("StoreFreeBytesInPercentage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float32(valuetmp)

	return
}

// SetStoreLastRecordDate sets the value of StoreLastRecordDate for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreLastRecordDate(value string) (err error) {
	return instance.SetProperty("StoreLastRecordDate", (value))
}

// GetStoreLastRecordDate gets the value of StoreLastRecordDate for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreLastRecordDate() (value string, err error) {
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
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreLimit(value string) (err error) {
	return instance.SetProperty("StoreLimit", (value))
}

// GetStoreLimit gets the value of StoreLimit for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreLimit() (value string, err error) {
	retValue, err := instance.GetProperty("StoreLimit")
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

// SetStoreUsedBytes sets the value of StoreUsedBytes for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreUsedBytes(value uint64) (err error) {
	return instance.SetProperty("StoreUsedBytes", (value))
}

// GetStoreUsedBytes gets the value of StoreUsedBytes for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreUsedBytes() (value uint64, err error) {
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

// SetStoreUsedBytesInPercentage sets the value of StoreUsedBytesInPercentage for the instance
func (instance *RemoteAccessInboxAccounting) SetPropertyStoreUsedBytesInPercentage(value float32) (err error) {
	return instance.SetProperty("StoreUsedBytesInPercentage", (value))
}

// GetStoreUsedBytesInPercentage gets the value of StoreUsedBytesInPercentage for the instance
func (instance *RemoteAccessInboxAccounting) GetPropertyStoreUsedBytesInPercentage() (value float32, err error) {
	retValue, err := instance.GetProperty("StoreUsedBytesInPercentage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float32(valuetmp)

	return
}
