// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnConnectionTriggerTrustedNetwork struct
type VpnConnectionTriggerTrustedNetwork struct {
	*cim.WmiInstance

	//
	ConnectionName string

	//
	DnsSuffix []string
}

func NewVpnConnectionTriggerTrustedNetworkEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionTriggerTrustedNetwork, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTriggerTrustedNetwork{
		WmiInstance: tmp,
	}
	return
}

func NewVpnConnectionTriggerTrustedNetworkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnectionTriggerTrustedNetwork, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTriggerTrustedNetwork{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", (value))
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) GetPropertyConnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionName")
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

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) SetPropertyDnsSuffix(value []string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) GetPropertyDnsSuffix() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
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
