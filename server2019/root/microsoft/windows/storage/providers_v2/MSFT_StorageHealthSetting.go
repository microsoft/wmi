// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageHealthSetting struct
type MSFT_StorageHealthSetting struct {
	cim.WmiInstance

	// Name of the storage health setting
	Name string

	// Value of the storage health setting.
	Value string
}

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageHealthSetting) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageHealthSetting) GetPropertyName() (value string, err error) {
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

// SetValue sets the value of Value for the instance
func (instance *MSFT_StorageHealthSetting) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_StorageHealthSetting) GetPropertyValue() (value string, err error) {
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
