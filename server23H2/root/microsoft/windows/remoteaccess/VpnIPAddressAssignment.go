// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnIPAddressAssignment struct
type VpnIPAddressAssignment struct {
	*cim.WmiInstance

	//
	IPAddressRange []VpnIPAddressRange

	//
	IPAssignmentMethod string

	//
	IPv6Prefix string
}

func NewVpnIPAddressAssignmentEx1(instance *cim.WmiInstance) (newInstance *VpnIPAddressAssignment, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnIPAddressAssignment{
		WmiInstance: tmp,
	}
	return
}

func NewVpnIPAddressAssignmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnIPAddressAssignment, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnIPAddressAssignment{
		WmiInstance: tmp,
	}
	return
}

// SetIPAddressRange sets the value of IPAddressRange for the instance
func (instance *VpnIPAddressAssignment) SetPropertyIPAddressRange(value []VpnIPAddressRange) (err error) {
	return instance.SetProperty("IPAddressRange", (value))
}

// GetIPAddressRange gets the value of IPAddressRange for the instance
func (instance *VpnIPAddressAssignment) GetPropertyIPAddressRange() (value []VpnIPAddressRange, err error) {
	retValue, err := instance.GetProperty("IPAddressRange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VpnIPAddressRange)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VpnIPAddressRange is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VpnIPAddressRange(valuetmp))
	}

	return
}

// SetIPAssignmentMethod sets the value of IPAssignmentMethod for the instance
func (instance *VpnIPAddressAssignment) SetPropertyIPAssignmentMethod(value string) (err error) {
	return instance.SetProperty("IPAssignmentMethod", (value))
}

// GetIPAssignmentMethod gets the value of IPAssignmentMethod for the instance
func (instance *VpnIPAddressAssignment) GetPropertyIPAssignmentMethod() (value string, err error) {
	retValue, err := instance.GetProperty("IPAssignmentMethod")
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

// SetIPv6Prefix sets the value of IPv6Prefix for the instance
func (instance *VpnIPAddressAssignment) SetPropertyIPv6Prefix(value string) (err error) {
	return instance.SetProperty("IPv6Prefix", (value))
}

// GetIPv6Prefix gets the value of IPv6Prefix for the instance
func (instance *VpnIPAddressAssignment) GetPropertyIPv6Prefix() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Prefix")
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
