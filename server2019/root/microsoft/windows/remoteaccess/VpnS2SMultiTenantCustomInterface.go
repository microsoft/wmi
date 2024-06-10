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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnS2SMultiTenantCustomInterface struct
type VpnS2SMultiTenantCustomInterface struct {
	*VpnS2SCustomInterface

	//
	AutoConnectEnabled bool

	//
	RadiusAttributeClass string

	//
	RoutingDomain string
}

func NewVpnS2SMultiTenantCustomInterfaceEx1(instance *cim.WmiInstance) (newInstance *VpnS2SMultiTenantCustomInterface, err error) {
	tmp, err := NewVpnS2SCustomInterfaceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantCustomInterface{
		VpnS2SCustomInterface: tmp,
	}
	return
}

func NewVpnS2SMultiTenantCustomInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnS2SMultiTenantCustomInterface, err error) {
	tmp, err := NewVpnS2SCustomInterfaceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantCustomInterface{
		VpnS2SCustomInterface: tmp,
	}
	return
}

// SetAutoConnectEnabled sets the value of AutoConnectEnabled for the instance
func (instance *VpnS2SMultiTenantCustomInterface) SetPropertyAutoConnectEnabled(value bool) (err error) {
	return instance.SetProperty("AutoConnectEnabled", (value))
}

// GetAutoConnectEnabled gets the value of AutoConnectEnabled for the instance
func (instance *VpnS2SMultiTenantCustomInterface) GetPropertyAutoConnectEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoConnectEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRadiusAttributeClass sets the value of RadiusAttributeClass for the instance
func (instance *VpnS2SMultiTenantCustomInterface) SetPropertyRadiusAttributeClass(value string) (err error) {
	return instance.SetProperty("RadiusAttributeClass", (value))
}

// GetRadiusAttributeClass gets the value of RadiusAttributeClass for the instance
func (instance *VpnS2SMultiTenantCustomInterface) GetPropertyRadiusAttributeClass() (value string, err error) {
	retValue, err := instance.GetProperty("RadiusAttributeClass")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *VpnS2SMultiTenantCustomInterface) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *VpnS2SMultiTenantCustomInterface) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
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
