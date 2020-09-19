// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LANEndpoint struct
type CIM_LANEndpoint struct {
	*CIM_ProtocolEndpoint

	//
	AliasAddresses []string

	//
	GroupAddresses []string

	//
	LANID string

	//
	LANType uint16

	//
	MACAddress string

	//
	MaxDataSize uint32

	//
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
func (instance *CIM_LANEndpoint) SetPropertyLANType(value uint16) (err error) {
	return instance.SetProperty("LANType", (value))
}

// GetLANType gets the value of LANType for the instance
func (instance *CIM_LANEndpoint) GetPropertyLANType() (value uint16, err error) {
	retValue, err := instance.GetProperty("LANType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

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
