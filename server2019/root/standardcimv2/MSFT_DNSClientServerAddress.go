// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DNSClientServerAddress struct
type MSFT_DNSClientServerAddress struct {
	*CIM_RemoteServiceAccessPoint

	// 747
	AddressFamily uint16

	// 656
	InterfaceAlias string

	// 655
	InterfaceIndex uint32

	// 746
	ServerAddresses []string
}

func NewMSFT_DNSClientServerAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_DNSClientServerAddress, err error) {
	tmp, err := NewCIM_RemoteServiceAccessPointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientServerAddress{
		CIM_RemoteServiceAccessPoint: tmp,
	}
	return
}

func NewMSFT_DNSClientServerAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DNSClientServerAddress, err error) {
	tmp, err := NewCIM_RemoteServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientServerAddress{
		CIM_RemoteServiceAccessPoint: tmp,
	}
	return
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", (value))
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetServerAddresses sets the value of ServerAddresses for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyServerAddresses(value []string) (err error) {
	return instance.SetProperty("ServerAddresses", (value))
}

// GetServerAddresses gets the value of ServerAddresses for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyServerAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("ServerAddresses")
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
