// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// VpnServerAddress struct
type VpnServerAddress struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	ServerAddress string
}

func NewVpnServerAddressEx1(instance *cim.WmiInstance) (newInstance *VpnServerAddress, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnServerAddress{
		WmiInstance: tmp,
	}
	return
}

func NewVpnServerAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnServerAddress, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnServerAddress{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *VpnServerAddress) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *VpnServerAddress) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *VpnServerAddress) SetPropertyServerAddress(value string) (err error) {
	return instance.SetProperty("ServerAddress", value)
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *VpnServerAddress) GetPropertyServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
