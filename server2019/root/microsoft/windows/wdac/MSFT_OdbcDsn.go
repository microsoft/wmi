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

// MSFT_OdbcDsn struct
type MSFT_OdbcDsn struct {
	cim.WmiInstance

	//
	DriverName string

	//
	DsnType string

	//
	KeyValuePair []MSFT_OdbcKeyValuePair

	//
	Name string

	//
	Platform string
}

// SetDriverName sets the value of DriverName for the instance
func (instance *MSFT_OdbcDsn) SetPropertyDriverName(value string) (err error) {
	return instance.SetProperty("DriverName", value)
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSFT_OdbcDsn) GetPropertyDriverName() (value string, err error) {
	retValue, err := instance.GetProperty("DriverName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDsnType sets the value of DsnType for the instance
func (instance *MSFT_OdbcDsn) SetPropertyDsnType(value string) (err error) {
	return instance.SetProperty("DsnType", value)
}

// GetDsnType gets the value of DsnType for the instance
func (instance *MSFT_OdbcDsn) GetPropertyDsnType() (value string, err error) {
	retValue, err := instance.GetProperty("DsnType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyValuePair sets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDsn) SetPropertyKeyValuePair(value []MSFT_OdbcKeyValuePair) (err error) {
	return instance.SetProperty("KeyValuePair", value)
}

// GetKeyValuePair gets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDsn) GetPropertyKeyValuePair() (value []MSFT_OdbcKeyValuePair, err error) {
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
func (instance *MSFT_OdbcDsn) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_OdbcDsn) GetPropertyName() (value string, err error) {
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
func (instance *MSFT_OdbcDsn) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", value)
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcDsn) GetPropertyPlatform() (value string, err error) {
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
