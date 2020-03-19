// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SwitchPort struct
type CIM_SwitchPort struct {
	*CIM_ProtocolEndpoint

	// Numeric identifier for a switch port.
	PortNumber uint16
}

func NewCIM_SwitchPortEx1(instance *cim.WmiInstance) (newInstance *CIM_SwitchPort, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchPort{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_SwitchPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SwitchPort, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchPort{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *CIM_SwitchPort) SetPropertyPortNumber(value uint16) (err error) {
	return instance.SetProperty("PortNumber", value)
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *CIM_SwitchPort) GetPropertyPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
