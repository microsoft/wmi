// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// GreS2SInterface struct
type GreS2SInterface struct {
	*VpnS2SInterface

	//
	GreKey uint32

	//
	IPv4Address string

	//
	IPv6Address string
}

func NewGreS2SInterfaceEx1(instance *cim.WmiInstance) (newInstance *GreS2SInterface, err error) {
	tmp, err := NewVpnS2SInterfaceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &GreS2SInterface{
		VpnS2SInterface: tmp,
	}
	return
}

func NewGreS2SInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GreS2SInterface, err error) {
	tmp, err := NewVpnS2SInterfaceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GreS2SInterface{
		VpnS2SInterface: tmp,
	}
	return
}

// SetGreKey sets the value of GreKey for the instance
func (instance *GreS2SInterface) SetPropertyGreKey(value uint32) (err error) {
	return instance.SetProperty("GreKey", (value))
}

// GetGreKey gets the value of GreKey for the instance
func (instance *GreS2SInterface) GetPropertyGreKey() (value uint32, err error) {
	retValue, err := instance.GetProperty("GreKey")
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

// SetIPv4Address sets the value of IPv4Address for the instance
func (instance *GreS2SInterface) SetPropertyIPv4Address(value string) (err error) {
	return instance.SetProperty("IPv4Address", (value))
}

// GetIPv4Address gets the value of IPv4Address for the instance
func (instance *GreS2SInterface) GetPropertyIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4Address")
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

// SetIPv6Address sets the value of IPv6Address for the instance
func (instance *GreS2SInterface) SetPropertyIPv6Address(value string) (err error) {
	return instance.SetProperty("IPv6Address", (value))
}

// GetIPv6Address gets the value of IPv6Address for the instance
func (instance *GreS2SInterface) GetPropertyIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Address")
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
