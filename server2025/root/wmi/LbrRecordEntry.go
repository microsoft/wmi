// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// LbrRecordEntry struct
type LbrRecordEntry struct {
	*cim.WmiInstance

	//
	FromAddress uint64

	//
	Reserved uint64

	//
	ToAddress uint64
}

func NewLbrRecordEntryEx1(instance *cim.WmiInstance) (newInstance *LbrRecordEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &LbrRecordEntry{
		WmiInstance: tmp,
	}
	return
}

func NewLbrRecordEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LbrRecordEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LbrRecordEntry{
		WmiInstance: tmp,
	}
	return
}

// SetFromAddress sets the value of FromAddress for the instance
func (instance *LbrRecordEntry) SetPropertyFromAddress(value uint64) (err error) {
	return instance.SetProperty("FromAddress", (value))
}

// GetFromAddress gets the value of FromAddress for the instance
func (instance *LbrRecordEntry) GetPropertyFromAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("FromAddress")
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

// SetReserved sets the value of Reserved for the instance
func (instance *LbrRecordEntry) SetPropertyReserved(value uint64) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *LbrRecordEntry) GetPropertyReserved() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetToAddress sets the value of ToAddress for the instance
func (instance *LbrRecordEntry) SetPropertyToAddress(value uint64) (err error) {
	return instance.SetProperty("ToAddress", (value))
}

// GetToAddress gets the value of ToAddress for the instance
func (instance *LbrRecordEntry) GetPropertyToAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("ToAddress")
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
