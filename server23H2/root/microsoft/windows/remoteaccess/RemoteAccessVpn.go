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

// RemoteAccessVpn struct
type RemoteAccessVpn struct {
	*RemoteAccessCommon

	//
	VpnConfiguration VirtualPrivateNetworkConfiguration
}

func NewRemoteAccessVpnEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessVpn, err error) {
	tmp, err := NewRemoteAccessCommonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessVpn{
		RemoteAccessCommon: tmp,
	}
	return
}

func NewRemoteAccessVpnEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessVpn, err error) {
	tmp, err := NewRemoteAccessCommonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessVpn{
		RemoteAccessCommon: tmp,
	}
	return
}

// SetVpnConfiguration sets the value of VpnConfiguration for the instance
func (instance *RemoteAccessVpn) SetPropertyVpnConfiguration(value VirtualPrivateNetworkConfiguration) (err error) {
	return instance.SetProperty("VpnConfiguration", (value))
}

// GetVpnConfiguration gets the value of VpnConfiguration for the instance
func (instance *RemoteAccessVpn) GetPropertyVpnConfiguration() (value VirtualPrivateNetworkConfiguration, err error) {
	retValue, err := instance.GetProperty("VpnConfiguration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VirtualPrivateNetworkConfiguration)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VirtualPrivateNetworkConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualPrivateNetworkConfiguration(valuetmp)

	return
}
