// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_GuestNetworkAdapterConfiguration struct
type Msvm_GuestNetworkAdapterConfiguration struct {
	*cim.WmiInstance

	//
	DefaultGateways []string

	//
	DHCPEnabled bool

	//
	DNSServers []string

	//
	InstanceID string

	//
	IPAddresses []string

	//
	IPAddressOrigins []GuestNetworkAdapterConfiguration_IPAddressOrigins

	//
	ProtocolIFType GuestNetworkAdapterConfiguration_ProtocolIFType

	//
	Subnets []string
}

func NewMsvm_GuestNetworkAdapterConfigurationEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestNetworkAdapterConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestNetworkAdapterConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_GuestNetworkAdapterConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestNetworkAdapterConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestNetworkAdapterConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetDefaultGateways sets the value of DefaultGateways for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyDefaultGateways(value []string) (err error) {
	return instance.SetProperty("DefaultGateways", (value))
}

// GetDefaultGateways gets the value of DefaultGateways for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDefaultGateways() (value []string, err error) {
	retValue, err := instance.GetProperty("DefaultGateways")
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

// SetDHCPEnabled sets the value of DHCPEnabled for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", (value))
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDHCPEnabled() (value bool, err error) {
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

// SetDNSServers sets the value of DNSServers for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyDNSServers(value []string) (err error) {
	return instance.SetProperty("DNSServers", (value))
}

// GetDNSServers gets the value of DNSServers for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDNSServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DNSServers")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", (value))
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddresses")
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

// SetIPAddressOrigins sets the value of IPAddressOrigins for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyIPAddressOrigins(value []GuestNetworkAdapterConfiguration_IPAddressOrigins) (err error) {
	return instance.SetProperty("IPAddressOrigins", (value))
}

// GetIPAddressOrigins gets the value of IPAddressOrigins for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyIPAddressOrigins() (value []GuestNetworkAdapterConfiguration_IPAddressOrigins, err error) {
	retValue, err := instance.GetProperty("IPAddressOrigins")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, GuestNetworkAdapterConfiguration_IPAddressOrigins(valuetmp))
	}

	return
}

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyProtocolIFType(value GuestNetworkAdapterConfiguration_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", (value))
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyProtocolIFType() (value GuestNetworkAdapterConfiguration_ProtocolIFType, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = GuestNetworkAdapterConfiguration_ProtocolIFType(valuetmp)

	return
}

// SetSubnets sets the value of Subnets for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertySubnets(value []string) (err error) {
	return instance.SetProperty("Subnets", (value))
}

// GetSubnets gets the value of Subnets for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertySubnets() (value []string, err error) {
	retValue, err := instance.GetProperty("Subnets")
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
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetRelatedSyntheticEthernetPortSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SyntheticEthernetPortSettingData")
}
