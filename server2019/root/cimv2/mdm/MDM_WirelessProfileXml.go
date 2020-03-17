// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_WirelessProfileXml struct
type MDM_WirelessProfileXml struct {
	cim.WmiInstance

	//
	Name string

	//
	ProfileXml string
}

// SetName sets the value of Name for the instance
func (instance *MDM_WirelessProfileXml) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MDM_WirelessProfileXml) GetPropertyName() (value string, err error) {
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

// SetProfileXml sets the value of ProfileXml for the instance
func (instance *MDM_WirelessProfileXml) SetPropertyProfileXml(value string) (err error) {
	return instance.SetProperty("ProfileXml", value)
}

// GetProfileXml gets the value of ProfileXml for the instance
func (instance *MDM_WirelessProfileXml) GetPropertyProfileXml() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileXml")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
