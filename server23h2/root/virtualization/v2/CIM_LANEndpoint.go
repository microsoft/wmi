// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LANEndpoint struct
type CIM_LANEndpoint struct {
	*CIM_ProtocolEndpoint

	// Other unicast addresses that may be used to communicate with the LANEndpoint.
	AliasAddresses []string

	// Multicast addresses to which the LANEndpoint listens.
	GroupAddresses []string

	// A label or identifier for the LAN Segment to which the Endpoint is connected. If the Endpoint is not currently active/connected or this information is not known, then LANID is NULL.
	LANID string

	// An indication of the kind of technology used on the LAN. This property is deprecated in lieu of ProtocolType, which is an enumeration inherited from ProtocolEndpoint and which includes the Values specified here.
	LANType LANEndpoint_LANType

	// The principal unicast address used in communication with the LANEndpoint. The MAC address is formatted as twelve hexadecimal digits (e.g., "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order according to RFC 2469.
	MACAddress string

	// The largest information field that may be sent or received by the LANEndpoint.
	MaxDataSize uint32

	// A free-form string that describes the type of technology used on the LAN when the value of the LANType property is equal to 1 (i.e., "Other"). This property is deprecated since its purpose overlaps with OtherTypeDescription, which which is inherited from ProtocolEndpoint.
	OtherLANType string
}

func NewCIM_LANEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_LANEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_LANEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_LANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LANEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LANEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetAliasAddresses sets the value of AliasAddresses for the instance
func (instance *CIM_LANEndpoint) SetPropertyAliasAddresses(value []string) (err error) {
	return instance.SetProperty("AliasAddresses", (value))
}

// GetAliasAddresses gets the value of AliasAddresses for the instance
func (instance *CIM_LANEndpoint) GetPropertyAliasAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("AliasAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetGroupAddresses sets the value of GroupAddresses for the instance
func (instance *CIM_LANEndpoint) SetPropertyGroupAddresses(value []string) (err error) {
	return instance.SetProperty("GroupAddresses", (value))
}

// GetGroupAddresses gets the value of GroupAddresses for the instance
func (instance *CIM_LANEndpoint) GetPropertyGroupAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("GroupAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetLANID sets the value of LANID for the instance
func (instance *CIM_LANEndpoint) SetPropertyLANID(value string) (err error) {
	return instance.SetProperty("LANID", (value))
}

// GetLANID gets the value of LANID for the instance
func (instance *CIM_LANEndpoint) GetPropertyLANID() (value string, err error) {
	retValue, err := instance.GetProperty("LANID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLANType sets the value of LANType for the instance
func (instance *CIM_LANEndpoint) SetPropertyLANType(value LANEndpoint_LANType) (err error) {
	return instance.SetProperty("LANType", (value))
}

// GetLANType gets the value of LANType for the instance
func (instance *CIM_LANEndpoint) GetPropertyLANType() (value LANEndpoint_LANType, err error) {
	retValue, err := instance.GetProperty("LANType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = LANEndpoint_LANType(valuetmp)

	return
}

// SetMACAddress sets the value of MACAddress for the instance
func (instance *CIM_LANEndpoint) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", (value))
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *CIM_LANEndpoint) GetPropertyMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MACAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMaxDataSize sets the value of MaxDataSize for the instance
func (instance *CIM_LANEndpoint) SetPropertyMaxDataSize(value uint32) (err error) {
	return instance.SetProperty("MaxDataSize", (value))
}

// GetMaxDataSize gets the value of MaxDataSize for the instance
func (instance *CIM_LANEndpoint) GetPropertyMaxDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDataSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOtherLANType sets the value of OtherLANType for the instance
func (instance *CIM_LANEndpoint) SetPropertyOtherLANType(value string) (err error) {
	return instance.SetProperty("OtherLANType", (value))
}

// GetOtherLANType gets the value of OtherLANType for the instance
func (instance *CIM_LANEndpoint) GetPropertyOtherLANType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLANType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
