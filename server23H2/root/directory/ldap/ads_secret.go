// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_secret struct
type ads_secret struct {
	*ds_leaf

	//
	DS_currentValue Uint8Array

	//
	DS_lastSetTime int64

	//
	DS_priorSetTime int64

	//
	DS_priorValue Uint8Array
}

func Newads_secretEx1(instance *cim.WmiInstance) (newInstance *ads_secret, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_secret{
		ds_leaf: tmp,
	}
	return
}

func Newads_secretEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_secret, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_secret{
		ds_leaf: tmp,
	}
	return
}

// SetDS_currentValue sets the value of DS_currentValue for the instance
func (instance *ads_secret) SetPropertyDS_currentValue(value Uint8Array) (err error) {
	return instance.SetProperty("DS_currentValue", (value))
}

// GetDS_currentValue gets the value of DS_currentValue for the instance
func (instance *ads_secret) GetPropertyDS_currentValue() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_currentValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_lastSetTime sets the value of DS_lastSetTime for the instance
func (instance *ads_secret) SetPropertyDS_lastSetTime(value int64) (err error) {
	return instance.SetProperty("DS_lastSetTime", (value))
}

// GetDS_lastSetTime gets the value of DS_lastSetTime for the instance
func (instance *ads_secret) GetPropertyDS_lastSetTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastSetTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_priorSetTime sets the value of DS_priorSetTime for the instance
func (instance *ads_secret) SetPropertyDS_priorSetTime(value int64) (err error) {
	return instance.SetProperty("DS_priorSetTime", (value))
}

// GetDS_priorSetTime gets the value of DS_priorSetTime for the instance
func (instance *ads_secret) GetPropertyDS_priorSetTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_priorSetTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_priorValue sets the value of DS_priorValue for the instance
func (instance *ads_secret) SetPropertyDS_priorValue(value Uint8Array) (err error) {
	return instance.SetProperty("DS_priorValue", (value))
}

// GetDS_priorValue gets the value of DS_priorValue for the instance
func (instance *ads_secret) GetPropertyDS_priorValue() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_priorValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}
