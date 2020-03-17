// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerEventLog struct
type MSFT_ServerEventLog struct {
	cim.WmiInstance

	//
	LocalizedName string

	//
	Name string
}

// SetLocalizedName sets the value of LocalizedName for the instance
func (instance *MSFT_ServerEventLog) SetPropertyLocalizedName(value string) (err error) {
	return instance.SetProperty("LocalizedName", value)
}

// GetLocalizedName gets the value of LocalizedName for the instance
func (instance *MSFT_ServerEventLog) GetPropertyLocalizedName() (value string, err error) {
	retValue, err := instance.GetProperty("LocalizedName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerEventLog) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerEventLog) GetPropertyName() (value string, err error) {
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
