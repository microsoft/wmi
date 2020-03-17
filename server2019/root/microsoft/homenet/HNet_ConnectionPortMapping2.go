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

// HNet_ConnectionPortMapping2 struct
type HNet_ConnectionPortMapping2 struct {
	cim.WmiInstance

	//
	Connection HNet_Connection

	//
	Enabled bool

	//
	NameActive bool

	//
	Protocol HNet_PortMappingProtocol

	//
	TargetIPAddress uint32

	//
	TargetName string

	//
	TargetPort uint16
}

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyConnection(value HNet_Connection) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyConnection() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyEnabled() (value bool, err error) {
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

// SetNameActive sets the value of NameActive for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyNameActive(value bool) (err error) {
	return instance.SetProperty("NameActive", value)
}

// GetNameActive gets the value of NameActive for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyNameActive() (value bool, err error) {
	retValue, err := instance.GetProperty("NameActive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyProtocol(value HNet_PortMappingProtocol) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyProtocol() (value HNet_PortMappingProtocol, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_PortMappingProtocol)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetIPAddress sets the value of TargetIPAddress for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetIPAddress(value uint32) (err error) {
	return instance.SetProperty("TargetIPAddress", value)
}

// GetTargetIPAddress gets the value of TargetIPAddress for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetIPAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetName sets the value of TargetName for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", value)
}

// GetTargetName gets the value of TargetName for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPort sets the value of TargetPort for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetPort(value uint16) (err error) {
	return instance.SetProperty("TargetPort", value)
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
