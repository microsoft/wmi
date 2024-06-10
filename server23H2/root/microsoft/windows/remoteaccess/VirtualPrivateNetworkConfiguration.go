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

// VirtualPrivateNetworkConfiguration struct
type VirtualPrivateNetworkConfiguration struct {
	*cim.WmiInstance

	//
	AuthenticationPolicy VpnAuth

	//
	IPAddressAssignmentPolicy VpnIPAddressAssignment
}

func NewVirtualPrivateNetworkConfigurationEx1(instance *cim.WmiInstance) (newInstance *VirtualPrivateNetworkConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VirtualPrivateNetworkConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewVirtualPrivateNetworkConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VirtualPrivateNetworkConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VirtualPrivateNetworkConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAuthenticationPolicy sets the value of AuthenticationPolicy for the instance
func (instance *VirtualPrivateNetworkConfiguration) SetPropertyAuthenticationPolicy(value VpnAuth) (err error) {
	return instance.SetProperty("AuthenticationPolicy", (value))
}

// GetAuthenticationPolicy gets the value of AuthenticationPolicy for the instance
func (instance *VirtualPrivateNetworkConfiguration) GetPropertyAuthenticationPolicy() (value VpnAuth, err error) {
	retValue, err := instance.GetProperty("AuthenticationPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VpnAuth)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VpnAuth is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VpnAuth(valuetmp)

	return
}

// SetIPAddressAssignmentPolicy sets the value of IPAddressAssignmentPolicy for the instance
func (instance *VirtualPrivateNetworkConfiguration) SetPropertyIPAddressAssignmentPolicy(value VpnIPAddressAssignment) (err error) {
	return instance.SetProperty("IPAddressAssignmentPolicy", (value))
}

// GetIPAddressAssignmentPolicy gets the value of IPAddressAssignmentPolicy for the instance
func (instance *VirtualPrivateNetworkConfiguration) GetPropertyIPAddressAssignmentPolicy() (value VpnIPAddressAssignment, err error) {
	retValue, err := instance.GetProperty("IPAddressAssignmentPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VpnIPAddressAssignment)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VpnIPAddressAssignment is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VpnIPAddressAssignment(valuetmp)

	return
}
