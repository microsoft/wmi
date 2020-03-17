// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_ExtensionEventSource struct
type RSOP_ExtensionEventSource struct {
	cim.WmiInstance

	//
	eventLogName string

	//
	eventLogSource string

	//
	id string
}

// SeteventLogName sets the value of eventLogName for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyeventLogName(value string) (err error) {
	return instance.SetProperty("eventLogName", value)
}

// GeteventLogName gets the value of eventLogName for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyeventLogName() (value string, err error) {
	retValue, err := instance.GetProperty("eventLogName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeteventLogSource sets the value of eventLogSource for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyeventLogSource(value string) (err error) {
	return instance.SetProperty("eventLogSource", value)
}

// GeteventLogSource gets the value of eventLogSource for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyeventLogSource() (value string, err error) {
	retValue, err := instance.GetProperty("eventLogSource")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setid sets the value of id for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
