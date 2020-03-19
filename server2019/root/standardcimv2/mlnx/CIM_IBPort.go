// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_IBPort struct
type CIM_IBPort struct {
	*CIM_NetworkPort

	//
	LIDMask uint8

	//
	LinkSpeedActive uint8

	//
	LinkWidthActive uint16
}

func NewCIM_IBPortEx1(instance *cim.WmiInstance) (newInstance *CIM_IBPort, err error) {
	tmp, err := NewCIM_NetworkPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IBPort{
		CIM_NetworkPort: tmp,
	}
	return
}

func NewCIM_IBPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IBPort, err error) {
	tmp, err := NewCIM_NetworkPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IBPort{
		CIM_NetworkPort: tmp,
	}
	return
}

// SetLIDMask sets the value of LIDMask for the instance
func (instance *CIM_IBPort) SetPropertyLIDMask(value uint8) (err error) {
	return instance.SetProperty("LIDMask", value)
}

// GetLIDMask gets the value of LIDMask for the instance
func (instance *CIM_IBPort) GetPropertyLIDMask() (value uint8, err error) {
	retValue, err := instance.GetProperty("LIDMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkSpeedActive sets the value of LinkSpeedActive for the instance
func (instance *CIM_IBPort) SetPropertyLinkSpeedActive(value uint8) (err error) {
	return instance.SetProperty("LinkSpeedActive", value)
}

// GetLinkSpeedActive gets the value of LinkSpeedActive for the instance
func (instance *CIM_IBPort) GetPropertyLinkSpeedActive() (value uint8, err error) {
	retValue, err := instance.GetProperty("LinkSpeedActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkWidthActive sets the value of LinkWidthActive for the instance
func (instance *CIM_IBPort) SetPropertyLinkWidthActive(value uint16) (err error) {
	return instance.SetProperty("LinkWidthActive", value)
}

// GetLinkWidthActive gets the value of LinkWidthActive for the instance
func (instance *CIM_IBPort) GetPropertyLinkWidthActive() (value uint16, err error) {
	retValue, err := instance.GetProperty("LinkWidthActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
