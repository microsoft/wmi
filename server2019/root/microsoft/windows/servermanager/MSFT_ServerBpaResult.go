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

// MSFT_ServerBpaResult struct
type MSFT_ServerBpaResult struct {
	cim.WmiInstance

	//
	BpaXPath string

	//
	Values []string
}

// SetBpaXPath sets the value of BpaXPath for the instance
func (instance *MSFT_ServerBpaResult) SetPropertyBpaXPath(value string) (err error) {
	return instance.SetProperty("BpaXPath", value)
}

// GetBpaXPath gets the value of BpaXPath for the instance
func (instance *MSFT_ServerBpaResult) GetPropertyBpaXPath() (value string, err error) {
	retValue, err := instance.GetProperty("BpaXPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValues sets the value of Values for the instance
func (instance *MSFT_ServerBpaResult) SetPropertyValues(value []string) (err error) {
	return instance.SetProperty("Values", value)
}

// GetValues gets the value of Values for the instance
func (instance *MSFT_ServerBpaResult) GetPropertyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("Values")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
