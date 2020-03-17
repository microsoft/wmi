// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcKeyValuePair struct
type MSFT_OdbcKeyValuePair struct {
	cim.WmiInstance

	//
	key string

	//
	ParentKey string

	//
	Value string
}

// Setkey sets the value of key for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentKey sets the value of ParentKey for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertyParentKey(value string) (err error) {
	return instance.SetProperty("ParentKey", value)
}

// GetParentKey gets the value of ParentKey for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertyParentKey() (value string, err error) {
	retValue, err := instance.GetProperty("ParentKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue sets the value of Value for the instance
func (instance *MSFT_OdbcKeyValuePair) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_OdbcKeyValuePair) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
