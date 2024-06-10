// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Wdac
//
// ////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_OdbcKeyValuePair struct
type MSFT_OdbcKeyValuePair struct {
	*cim.WmiInstance

	//
	key string

	//
	ParentKey string

	//
	Value string
}

func NewMSFT_OdbcKeyValuePairEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcKeyValuePair, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcKeyValuePair{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OdbcKeyValuePairEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OdbcKeyValuePair, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcKeyValuePair{
		WmiInstance: tmp,
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertykey() (value string, err error) {
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

// SetParentKey sets the value of ParentKey for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertyParentKey(value string) (err error) {
	return instance.SetProperty("ParentKey", (value))
}

// GetParentKey gets the value of ParentKey for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertyParentKey() (value string, err error) {
	retValue, err := instance.GetProperty("ParentKey")
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

// SetValue sets the value of Value for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
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
