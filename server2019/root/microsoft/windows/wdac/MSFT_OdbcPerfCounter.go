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

// MSFT_OdbcPerfCounter struct
type MSFT_OdbcPerfCounter struct {
	cim.WmiInstance

	//
	IsEnabled bool

	//
	Platform string
}

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *MSFT_OdbcPerfCounter) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", value)
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *MSFT_OdbcPerfCounter) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_OdbcPerfCounter) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", value)
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcPerfCounter) GetPropertyPlatform() (value string, err error) {
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
