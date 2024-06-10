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

// VpnSstpProxyRule struct
type VpnSstpProxyRule struct {
	*cim.WmiInstance

	//
	GatewayAddress []string

	//
	TenantID string
}

func NewVpnSstpProxyRuleEx1(instance *cim.WmiInstance) (newInstance *VpnSstpProxyRule, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnSstpProxyRule{
		WmiInstance: tmp,
	}
	return
}

func NewVpnSstpProxyRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnSstpProxyRule, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnSstpProxyRule{
		WmiInstance: tmp,
	}
	return
}

// SetGatewayAddress sets the value of GatewayAddress for the instance
func (instance *VpnSstpProxyRule) SetPropertyGatewayAddress(value []string) (err error) {
	return instance.SetProperty("GatewayAddress", (value))
}

// GetGatewayAddress gets the value of GatewayAddress for the instance
func (instance *VpnSstpProxyRule) GetPropertyGatewayAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("GatewayAddress")
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

// SetTenantID sets the value of TenantID for the instance
func (instance *VpnSstpProxyRule) SetPropertyTenantID(value string) (err error) {
	return instance.SetProperty("TenantID", (value))
}

// GetTenantID gets the value of TenantID for the instance
func (instance *VpnSstpProxyRule) GetPropertyTenantID() (value string, err error) {
	retValue, err := instance.GetProperty("TenantID")
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
