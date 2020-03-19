// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_PortMappingProtocol struct
type HNet_PortMappingProtocol struct {
	*cim.WmiInstance

	//
	BuiltIn bool

	//
	Id string

	//
	IPProtocol uint8

	//
	Name string

	//
	Port uint16
}

func NewHNet_PortMappingProtocolEx1(instance *cim.WmiInstance) (newInstance *HNet_PortMappingProtocol, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_PortMappingProtocol{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_PortMappingProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_PortMappingProtocol, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_PortMappingProtocol{
		WmiInstance: tmp,
	}
	return
}

// SetBuiltIn sets the value of BuiltIn for the instance
func (instance *HNet_PortMappingProtocol) SetPropertyBuiltIn(value bool) (err error) {
	return instance.SetProperty("BuiltIn", value)
}

// GetBuiltIn gets the value of BuiltIn for the instance
func (instance *HNet_PortMappingProtocol) GetPropertyBuiltIn() (value bool, err error) {
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

// SetId sets the value of Id for the instance
func (instance *HNet_PortMappingProtocol) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *HNet_PortMappingProtocol) GetPropertyId() (value string, err error) {
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

// SetIPProtocol sets the value of IPProtocol for the instance
func (instance *HNet_PortMappingProtocol) SetPropertyIPProtocol(value uint8) (err error) {
	return instance.SetProperty("IPProtocol", value)
}

// GetIPProtocol gets the value of IPProtocol for the instance
func (instance *HNet_PortMappingProtocol) GetPropertyIPProtocol() (value uint8, err error) {
	retValue, err := instance.GetProperty("IPProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *HNet_PortMappingProtocol) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *HNet_PortMappingProtocol) GetPropertyName() (value string, err error) {
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

// SetPort sets the value of Port for the instance
func (instance *HNet_PortMappingProtocol) SetPropertyPort(value uint16) (err error) {
	return instance.SetProperty("Port", value)
}

// GetPort gets the value of Port for the instance
func (instance *HNet_PortMappingProtocol) GetPropertyPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
