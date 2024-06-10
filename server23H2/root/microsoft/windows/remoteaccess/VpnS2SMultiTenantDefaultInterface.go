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

// VpnS2SMultiTenantDefaultInterface struct
type VpnS2SMultiTenantDefaultInterface struct {
	*VpnS2SDefaultInterface

	//
	AutoConnectEnabled bool

	//
	RadiusAttributeClass string

	//
	RoutingDomain string
}

func NewVpnS2SMultiTenantDefaultInterfaceEx1(instance *cim.WmiInstance) (newInstance *VpnS2SMultiTenantDefaultInterface, err error) {
	tmp, err := NewVpnS2SDefaultInterfaceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantDefaultInterface{
		VpnS2SDefaultInterface: tmp,
	}
	return
}

func NewVpnS2SMultiTenantDefaultInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnS2SMultiTenantDefaultInterface, err error) {
	tmp, err := NewVpnS2SDefaultInterfaceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantDefaultInterface{
		VpnS2SDefaultInterface: tmp,
	}
	return
}

// SetAutoConnectEnabled sets the value of AutoConnectEnabled for the instance
func (instance *VpnS2SMultiTenantDefaultInterface) SetPropertyAutoConnectEnabled(value bool) (err error) {
	return instance.SetProperty("AutoConnectEnabled", (value))
}

// GetAutoConnectEnabled gets the value of AutoConnectEnabled for the instance
func (instance *VpnS2SMultiTenantDefaultInterface) GetPropertyAutoConnectEnabled() (value bool, err error) {
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
func (instance *VpnS2SMultiTenantDefaultInterface) SetPropertyRadiusAttributeClass(value string) (err error) {
	return instance.SetProperty("RadiusAttributeClass", (value))
}

// GetRadiusAttributeClass gets the value of RadiusAttributeClass for the instance
func (instance *VpnS2SMultiTenantDefaultInterface) GetPropertyRadiusAttributeClass() (value string, err error) {
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
func (instance *VpnS2SMultiTenantDefaultInterface) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *VpnS2SMultiTenantDefaultInterface) GetPropertyRoutingDomain() (value string, err error) {
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
