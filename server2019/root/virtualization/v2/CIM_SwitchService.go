// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SwitchService struct
type CIM_SwitchService struct {
	*CIM_ForwardingService

	// Address used by this SwitchService when it must be uniquely identified. For an ethernet bridge, the MAC Address serves as the BridgeAddress. When concatenated with a SpanningTreeService Priority, a unique bridge identifier results. The MAC address is formatted as twelve hexadecimal digits (e.g., "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order according to RFC 2469. In other scenarios, like Ipv6, the address is formatted as "ffff:ffff:ffff:ffff".
	BridgeAddress string

	// BridgeAddressType defines the type of addressing scheme used for this Bridge and its BridgeAddress property.
	BridgeAddressType SwitchService_BridgeAddressType

	// Indicates what type of switching service can be performed.
	BridgeType SwitchService_BridgeType

	// The number of switch ports controlled by this switching service.
	NumPorts uint16
}

func NewCIM_SwitchServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_SwitchService, err error) {
	tmp, err := NewCIM_ForwardingServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchService{
		CIM_ForwardingService: tmp,
	}
	return
}

func NewCIM_SwitchServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SwitchService, err error) {
	tmp, err := NewCIM_ForwardingServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchService{
		CIM_ForwardingService: tmp,
	}
	return
}

// SetBridgeAddress sets the value of BridgeAddress for the instance
func (instance *CIM_SwitchService) SetPropertyBridgeAddress(value string) (err error) {
	return instance.SetProperty("BridgeAddress", value)
}

// GetBridgeAddress gets the value of BridgeAddress for the instance
func (instance *CIM_SwitchService) GetPropertyBridgeAddress() (value string, err error) {
	retValue, err := instance.GetProperty("BridgeAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBridgeAddressType sets the value of BridgeAddressType for the instance
func (instance *CIM_SwitchService) SetPropertyBridgeAddressType(value SwitchService_BridgeAddressType) (err error) {
	return instance.SetProperty("BridgeAddressType", value)
}

// GetBridgeAddressType gets the value of BridgeAddressType for the instance
func (instance *CIM_SwitchService) GetPropertyBridgeAddressType() (value SwitchService_BridgeAddressType, err error) {
	retValue, err := instance.GetProperty("BridgeAddressType")
	if err != nil {
		return
	}
	value, ok := retValue.(SwitchService_BridgeAddressType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBridgeType sets the value of BridgeType for the instance
func (instance *CIM_SwitchService) SetPropertyBridgeType(value SwitchService_BridgeType) (err error) {
	return instance.SetProperty("BridgeType", value)
}

// GetBridgeType gets the value of BridgeType for the instance
func (instance *CIM_SwitchService) GetPropertyBridgeType() (value SwitchService_BridgeType, err error) {
	retValue, err := instance.GetProperty("BridgeType")
	if err != nil {
		return
	}
	value, ok := retValue.(SwitchService_BridgeType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumPorts sets the value of NumPorts for the instance
func (instance *CIM_SwitchService) SetPropertyNumPorts(value uint16) (err error) {
	return instance.SetProperty("NumPorts", value)
}

// GetNumPorts gets the value of NumPorts for the instance
func (instance *CIM_SwitchService) GetPropertyNumPorts() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumPorts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
