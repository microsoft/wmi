// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerNetworkAdapter struct
type MSFT_ServerNetworkAdapter struct {
	*cim.WmiInstance

	//
	Addresses []string

	//
	ConnectionStatus uint16

	//
	DHCPEnabled bool

	//
	Name string
}

func NewMSFT_ServerNetworkAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerNetworkAdapter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerNetworkAdapter{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerNetworkAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerNetworkAdapter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerNetworkAdapter{
		WmiInstance: tmp,
	}
	return
}

// SetAddresses sets the value of Addresses for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyAddresses(value []string) (err error) {
	return instance.SetProperty("Addresses", (value))
}

// GetAddresses gets the value of Addresses for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("Addresses")
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

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyConnectionStatus(value uint16) (err error) {
	return instance.SetProperty("ConnectionStatus", (value))
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyConnectionStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConnectionStatus")
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

// SetDHCPEnabled sets the value of DHCPEnabled for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", (value))
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyDHCPEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DHCPEnabled")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
