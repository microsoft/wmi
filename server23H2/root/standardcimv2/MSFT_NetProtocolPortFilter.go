// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetProtocolPortFilter struct
type MSFT_NetProtocolPortFilter struct {
	*CIM_FilterEntryBase

	//
	DynamicTransport uint32

	//
	IcmpType []string

	//
	LocalPort []string

	//
	Protocol string

	//
	RemotePort []string
}

func NewMSFT_NetProtocolPortFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetProtocolPortFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetProtocolPortFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetProtocolPortFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetProtocolPortFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetProtocolPortFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetDynamicTransport sets the value of DynamicTransport for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyDynamicTransport(value uint32) (err error) {
	return instance.SetProperty("DynamicTransport", (value))
}

// GetDynamicTransport gets the value of DynamicTransport for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyDynamicTransport() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicTransport")
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

// SetIcmpType sets the value of IcmpType for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyIcmpType(value []string) (err error) {
	return instance.SetProperty("IcmpType", (value))
}

// GetIcmpType gets the value of IcmpType for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyIcmpType() (value []string, err error) {
	retValue, err := instance.GetProperty("IcmpType")
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

// SetLocalPort sets the value of LocalPort for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyLocalPort(value []string) (err error) {
	return instance.SetProperty("LocalPort", (value))
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyLocalPort() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalPort")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemotePort sets the value of RemotePort for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyRemotePort(value []string) (err error) {
	return instance.SetProperty("RemotePort", (value))
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyRemotePort() (value []string, err error) {
	retValue, err := instance.GetProperty("RemotePort")
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
