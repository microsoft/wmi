// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("DefaultGateways", value)
}

// GetDefaultGateways gets the value of DefaultGateways for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDefaultGateways() (value []string, err error) {
	retValue, err := instance.GetProperty("DefaultGateways")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDHCPEnabled sets the value of DHCPEnabled for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", value)
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDHCPEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DHCPEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDNSServers sets the value of DNSServers for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyDNSServers(value []string) (err error) {
	return instance.SetProperty("DNSServers", value)
}

// GetDNSServers gets the value of DNSServers for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyDNSServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DNSServers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", value)
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddressOrigins sets the value of IPAddressOrigins for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyIPAddressOrigins(value []GuestNetworkAdapterConfiguration_IPAddressOrigins) (err error) {
	return instance.SetProperty("IPAddressOrigins", value)
}

// GetIPAddressOrigins gets the value of IPAddressOrigins for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyIPAddressOrigins() (value []GuestNetworkAdapterConfiguration_IPAddressOrigins, err error) {
	retValue, err := instance.GetProperty("IPAddressOrigins")
	if err != nil {
		return
	}
	value, ok := retValue.([]GuestNetworkAdapterConfiguration_IPAddressOrigins)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertyProtocolIFType(value GuestNetworkAdapterConfiguration_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", value)
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertyProtocolIFType() (value GuestNetworkAdapterConfiguration_ProtocolIFType, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
	if err != nil {
		return
	}
	value, ok := retValue.(GuestNetworkAdapterConfiguration_ProtocolIFType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubnets sets the value of Subnets for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) SetPropertySubnets(value []string) (err error) {
	return instance.SetProperty("Subnets", value)
}

// GetSubnets gets the value of Subnets for the instance
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetPropertySubnets() (value []string, err error) {
	retValue, err := instance.GetProperty("Subnets")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_GuestNetworkAdapterConfiguration) GetRelatedSyntheticEthernetPortSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SyntheticEthernetPortSettingData")
}
