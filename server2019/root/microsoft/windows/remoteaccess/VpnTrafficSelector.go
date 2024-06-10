// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnTrafficSelector struct
type VpnTrafficSelector struct {
	*cim.WmiInstance

	//
	IPAddressRange []string

	//
	PortRange []uint32

	//
	ProtocolId uint32

	//
	TsPayloadId uint16

	//
	Type uint32
}

func NewVpnTrafficSelectorEx1(instance *cim.WmiInstance) (newInstance *VpnTrafficSelector, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnTrafficSelector{
		WmiInstance: tmp,
	}
	return
}

func NewVpnTrafficSelectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnTrafficSelector, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnTrafficSelector{
		WmiInstance: tmp,
	}
	return
}

// SetIPAddressRange sets the value of IPAddressRange for the instance
func (instance *VpnTrafficSelector) SetPropertyIPAddressRange(value []string) (err error) {
	return instance.SetProperty("IPAddressRange", (value))
}

// GetIPAddressRange gets the value of IPAddressRange for the instance
func (instance *VpnTrafficSelector) GetPropertyIPAddressRange() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddressRange")
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

// SetPortRange sets the value of PortRange for the instance
func (instance *VpnTrafficSelector) SetPropertyPortRange(value []uint32) (err error) {
	return instance.SetProperty("PortRange", (value))
}

// GetPortRange gets the value of PortRange for the instance
func (instance *VpnTrafficSelector) GetPropertyPortRange() (value []uint32, err error) {
	retValue, err := instance.GetProperty("PortRange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetProtocolId sets the value of ProtocolId for the instance
func (instance *VpnTrafficSelector) SetPropertyProtocolId(value uint32) (err error) {
	return instance.SetProperty("ProtocolId", (value))
}

// GetProtocolId gets the value of ProtocolId for the instance
func (instance *VpnTrafficSelector) GetPropertyProtocolId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProtocolId")
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

// SetTsPayloadId sets the value of TsPayloadId for the instance
func (instance *VpnTrafficSelector) SetPropertyTsPayloadId(value uint16) (err error) {
	return instance.SetProperty("TsPayloadId", (value))
}

// GetTsPayloadId gets the value of TsPayloadId for the instance
func (instance *VpnTrafficSelector) GetPropertyTsPayloadId() (value uint16, err error) {
	retValue, err := instance.GetProperty("TsPayloadId")
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

// SetType sets the value of Type for the instance
func (instance *VpnTrafficSelector) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *VpnTrafficSelector) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
