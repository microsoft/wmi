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

// CIM_IPProtocolEndpoint struct
type CIM_IPProtocolEndpoint struct {
	*CIM_ProtocolEndpoint

	//
	Address string

	//
	AddressOrigin uint16

	//
	AddressType uint16

	//
	IPv4Address string

	//
	IPv6Address string

	//
	IPVersionSupport uint16

	//
	PrefixLength uint8

	//
	SubnetMask string
}

func NewCIM_IPProtocolEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_IPProtocolEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IPProtocolEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_IPProtocolEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IPProtocolEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IPProtocolEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAddressOrigin sets the value of AddressOrigin for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyAddressOrigin(value uint16) (err error) {
	return instance.SetProperty("AddressOrigin", value)
}

// GetAddressOrigin gets the value of AddressOrigin for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyAddressOrigin() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressOrigin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAddressType sets the value of AddressType for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyAddressType(value uint16) (err error) {
	return instance.SetProperty("AddressType", value)
}

// GetAddressType gets the value of AddressType for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyAddressType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4Address sets the value of IPv4Address for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyIPv4Address(value string) (err error) {
	return instance.SetProperty("IPv4Address", value)
}

// GetIPv4Address gets the value of IPv4Address for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Address sets the value of IPv6Address for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyIPv6Address(value string) (err error) {
	return instance.SetProperty("IPv6Address", value)
}

// GetIPv6Address gets the value of IPv6Address for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPVersionSupport sets the value of IPVersionSupport for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyIPVersionSupport(value uint16) (err error) {
	return instance.SetProperty("IPVersionSupport", value)
}

// GetIPVersionSupport gets the value of IPVersionSupport for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyIPVersionSupport() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPVersionSupport")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrefixLength sets the value of PrefixLength for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertyPrefixLength(value uint8) (err error) {
	return instance.SetProperty("PrefixLength", value)
}

// GetPrefixLength gets the value of PrefixLength for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertyPrefixLength() (value uint8, err error) {
	retValue, err := instance.GetProperty("PrefixLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *CIM_IPProtocolEndpoint) SetPropertySubnetMask(value string) (err error) {
	return instance.SetProperty("SubnetMask", value)
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *CIM_IPProtocolEndpoint) GetPropertySubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetMask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
