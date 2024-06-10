// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// VpnConnectionTrigger struct
type VpnConnectionTrigger struct {
	*cim.WmiInstance

	//
	ApplicationID []string

	//
	ConnectionName string

	//
	dnsConfig []VpnConnectionTriggerDnsConfiguration

	//
	DnsSuffixSearchList []string

	//
	TrustedNetwork []string
}

func NewVpnConnectionTriggerEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionTrigger, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTrigger{
		WmiInstance: tmp,
	}
	return
}

func NewVpnConnectionTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnectionTrigger, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnectionTrigger{
		WmiInstance: tmp,
	}
	return
}

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *VpnConnectionTrigger) SetPropertyApplicationID(value []string) (err error) {
	return instance.SetProperty("ApplicationID", (value))
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *VpnConnectionTrigger) GetPropertyApplicationID() (value []string, err error) {
	retValue, err := instance.GetProperty("ApplicationID")
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

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTrigger) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", (value))
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTrigger) GetPropertyConnectionName() (value string, err error) {
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

// SetdnsConfig sets the value of dnsConfig for the instance
func (instance *VpnConnectionTrigger) SetPropertydnsConfig(value []VpnConnectionTriggerDnsConfiguration) (err error) {
	return instance.SetProperty("dnsConfig", (value))
}

// GetdnsConfig gets the value of dnsConfig for the instance
func (instance *VpnConnectionTrigger) GetPropertydnsConfig() (value []VpnConnectionTriggerDnsConfiguration, err error) {
	retValue, err := instance.GetProperty("dnsConfig")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VpnConnectionTriggerDnsConfiguration)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VpnConnectionTriggerDnsConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VpnConnectionTriggerDnsConfiguration(valuetmp))
	}

	return
}

// SetDnsSuffixSearchList sets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTrigger) SetPropertyDnsSuffixSearchList(value []string) (err error) {
	return instance.SetProperty("DnsSuffixSearchList", (value))
}

// GetDnsSuffixSearchList gets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTrigger) GetPropertyDnsSuffixSearchList() (value []string, err error) {
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

// SetTrustedNetwork sets the value of TrustedNetwork for the instance
func (instance *VpnConnectionTrigger) SetPropertyTrustedNetwork(value []string) (err error) {
	return instance.SetProperty("TrustedNetwork", (value))
}

// GetTrustedNetwork gets the value of TrustedNetwork for the instance
func (instance *VpnConnectionTrigger) GetPropertyTrustedNetwork() (value []string, err error) {
	retValue, err := instance.GetProperty("TrustedNetwork")
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
