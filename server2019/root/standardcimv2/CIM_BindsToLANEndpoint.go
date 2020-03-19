// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BindsToLANEndpoint struct
type CIM_BindsToLANEndpoint struct {
	*CIM_BindsTo

	//
	FrameType uint16
}

func NewCIM_BindsToLANEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_BindsToLANEndpoint, err error) {
	tmp, err := NewCIM_BindsToEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BindsToLANEndpoint{
		CIM_BindsTo: tmp,
	}
	return
}

func NewCIM_BindsToLANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BindsToLANEndpoint, err error) {
	tmp, err := NewCIM_BindsToEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BindsToLANEndpoint{
		CIM_BindsTo: tmp,
	}
	return
}

// SetFrameType sets the value of FrameType for the instance
func (instance *CIM_BindsToLANEndpoint) SetPropertyFrameType(value uint16) (err error) {
	return instance.SetProperty("FrameType", value)
}

// GetFrameType gets the value of FrameType for the instance
func (instance *CIM_BindsToLANEndpoint) GetPropertyFrameType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FrameType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
