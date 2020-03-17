// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_ApplicationProtocol struct
type HNet_ApplicationProtocol struct {
	cim.WmiInstance

	//
	BuiltIn bool

	//
	Enabled bool

	//
	Id string

	//
	Name string

	//
	OutgoingIPProtocol uint8

	//
	OutgoingPort uint16

	//
	ResponseArray []HNet_ResponseRange

	//
	ResponseCount uint16
}

// SetBuiltIn sets the value of BuiltIn for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyBuiltIn(value bool) (err error) {
	return instance.SetProperty("BuiltIn", value)
}

// GetBuiltIn gets the value of BuiltIn for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyBuiltIn() (value bool, err error) {
	retValue, err := instance.GetProperty("BuiltIn")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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
func (instance *HNet_ApplicationProtocol) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyName() (value string, err error) {
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

// SetOutgoingIPProtocol sets the value of OutgoingIPProtocol for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyOutgoingIPProtocol(value uint8) (err error) {
	return instance.SetProperty("OutgoingIPProtocol", value)
}

// GetOutgoingIPProtocol gets the value of OutgoingIPProtocol for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyOutgoingIPProtocol() (value uint8, err error) {
	retValue, err := instance.GetProperty("OutgoingIPProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutgoingPort sets the value of OutgoingPort for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyOutgoingPort(value uint16) (err error) {
	return instance.SetProperty("OutgoingPort", value)
}

// GetOutgoingPort gets the value of OutgoingPort for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyOutgoingPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("OutgoingPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseArray sets the value of ResponseArray for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyResponseArray(value []HNet_ResponseRange) (err error) {
	return instance.SetProperty("ResponseArray", value)
}

// GetResponseArray gets the value of ResponseArray for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyResponseArray() (value []HNet_ResponseRange, err error) {
	retValue, err := instance.GetProperty("ResponseArray")
	if err != nil {
		return
	}
	value, ok := retValue.([]HNet_ResponseRange)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseCount sets the value of ResponseCount for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyResponseCount(value uint16) (err error) {
	return instance.SetProperty("ResponseCount", value)
}

// GetResponseCount gets the value of ResponseCount for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyResponseCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResponseCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
