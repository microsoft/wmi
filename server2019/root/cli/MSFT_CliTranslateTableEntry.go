// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CliTranslateTableEntry struct
type MSFT_CliTranslateTableEntry struct {
	cim.WmiInstance

	//
	FromValue string

	//
	ToValue string
}

// SetFromValue sets the value of FromValue for the instance
func (instance *MSFT_CliTranslateTableEntry) SetPropertyFromValue(value string) (err error) {
	return instance.SetProperty("FromValue", value)
}

// GetFromValue gets the value of FromValue for the instance
func (instance *MSFT_CliTranslateTableEntry) GetPropertyFromValue() (value string, err error) {
	retValue, err := instance.GetProperty("FromValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetToValue sets the value of ToValue for the instance
func (instance *MSFT_CliTranslateTableEntry) SetPropertyToValue(value string) (err error) {
	return instance.SetProperty("ToValue", value)
}

// GetToValue gets the value of ToValue for the instance
func (instance *MSFT_CliTranslateTableEntry) GetPropertyToValue() (value string, err error) {
	retValue, err := instance.GetProperty("ToValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
