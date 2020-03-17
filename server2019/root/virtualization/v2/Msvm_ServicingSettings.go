// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ServicingSettings struct
type Msvm_ServicingSettings struct {
	cim.WmiInstance

	//
	Version string
}

// SetVersion sets the value of Version for the instance
func (instance *Msvm_ServicingSettings) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_ServicingSettings) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
