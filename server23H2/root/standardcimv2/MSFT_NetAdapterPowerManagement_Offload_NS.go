// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterPowerManagement_Offload_NS struct
type MSFT_NetAdapterPowerManagement_Offload_NS struct {
	*MSFT_NetAdapterPowerManagement_Offload

	//
	MacAddress string

	//
	RemoteIPv6Address string

	//
	SolicitedNodeIPv6Address string

	//
	TargetIPv6Addresses []string
}

func NewMSFT_NetAdapterPowerManagement_Offload_NSEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_Offload_NS, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_NS{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_Offload_NSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_Offload_NS, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_NS{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", (value))
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
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

// SetRemoteIPv6Address sets the value of RemoteIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyRemoteIPv6Address(value string) (err error) {
	return instance.SetProperty("RemoteIPv6Address", (value))
}

// GetRemoteIPv6Address gets the value of RemoteIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyRemoteIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteIPv6Address")
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

// SetSolicitedNodeIPv6Address sets the value of SolicitedNodeIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertySolicitedNodeIPv6Address(value string) (err error) {
	return instance.SetProperty("SolicitedNodeIPv6Address", (value))
}

// GetSolicitedNodeIPv6Address gets the value of SolicitedNodeIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertySolicitedNodeIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("SolicitedNodeIPv6Address")
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

// SetTargetIPv6Addresses sets the value of TargetIPv6Addresses for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyTargetIPv6Addresses(value []string) (err error) {
	return instance.SetProperty("TargetIPv6Addresses", (value))
}

// GetTargetIPv6Addresses gets the value of TargetIPv6Addresses for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyTargetIPv6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("TargetIPv6Addresses")
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
