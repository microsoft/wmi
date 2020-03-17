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

// MSFT_OdbcDriver struct
type MSFT_OdbcDriver struct {
	cim.WmiInstance

	//
	KeyValuePair []MSFT_OdbcKeyValuePair

	//
	Name string

	//
	Platform string
}

// SetKeyValuePair sets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDriver) SetPropertyKeyValuePair(value []MSFT_OdbcKeyValuePair) (err error) {
	return instance.SetProperty("KeyValuePair", value)
}

// GetKeyValuePair gets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDriver) GetPropertyKeyValuePair() (value []MSFT_OdbcKeyValuePair, err error) {
	retValue, err := instance.GetProperty("KeyValuePair")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_OdbcKeyValuePair)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_OdbcDriver) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_OdbcDriver) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_OdbcDriver) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", value)
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcDriver) GetPropertyPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("Platform")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
