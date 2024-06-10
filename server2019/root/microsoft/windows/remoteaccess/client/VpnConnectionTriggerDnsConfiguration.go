// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Client
//
// ////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnConnectionTriggerDnsConfiguration struct
type VpnConnectionTriggerDnsConfiguration struct {
	*cim.WmiInstance

	//
	ConnectionName string

	//
	DnsIPAddress []string

	//
	DnsSuffix string

	//
	DnsSuffixSearchList []string
}

func NewVpnConnectionTriggerDnsConfigurationEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionTriggerDnsConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTriggerDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewVpnConnectionTriggerDnsConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnectionTriggerDnsConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTriggerDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", (value))
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyConnectionName() (value string, err error) {
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

// SetDnsIPAddress sets the value of DnsIPAddress for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsIPAddress(value []string) (err error) {
	return instance.SetProperty("DnsIPAddress", (value))
}

// GetDnsIPAddress gets the value of DnsIPAddress for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsIPAddress")
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

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
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

// SetDnsSuffixSearchList sets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffixSearchList(value []string) (err error) {
	return instance.SetProperty("DnsSuffixSearchList", (value))
}

// GetDnsSuffixSearchList gets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffixSearchList")
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
