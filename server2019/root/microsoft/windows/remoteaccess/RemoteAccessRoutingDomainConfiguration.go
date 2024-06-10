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

// RemoteAccessRoutingDomainConfiguration struct
type RemoteAccessRoutingDomainConfiguration struct {
	*cim.WmiInstance

	//
	Ipv4TransportInfo []uint8

	//
	Ipv6TransportInfo []uint8

	//
	RoutingDomainConfig interface{}
}

func NewRemoteAccessRoutingDomainConfigurationEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessRoutingDomainConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessRoutingDomainConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessRoutingDomainConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetIpv4TransportInfo sets the value of Ipv4TransportInfo for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) SetPropertyIpv4TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("Ipv4TransportInfo", (value))
}

// GetIpv4TransportInfo gets the value of Ipv4TransportInfo for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) GetPropertyIpv4TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ipv4TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetIpv6TransportInfo sets the value of Ipv6TransportInfo for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) SetPropertyIpv6TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("Ipv6TransportInfo", (value))
}

// GetIpv6TransportInfo gets the value of Ipv6TransportInfo for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) GetPropertyIpv6TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ipv6TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetRoutingDomainConfig sets the value of RoutingDomainConfig for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) SetPropertyRoutingDomainConfig(value interface{}) (err error) {
	return instance.SetProperty("RoutingDomainConfig", (value))
}

// GetRoutingDomainConfig gets the value of RoutingDomainConfig for the instance
func (instance *RemoteAccessRoutingDomainConfiguration) GetPropertyRoutingDomainConfig() (value interface{}, err error) {
	retValue, err := instance.GetProperty("RoutingDomainConfig")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
