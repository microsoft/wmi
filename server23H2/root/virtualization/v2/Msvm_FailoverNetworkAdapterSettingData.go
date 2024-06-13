// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_FailoverNetworkAdapterSettingData struct
type Msvm_FailoverNetworkAdapterSettingData struct {
	*CIM_SettingData

	//
	DefaultGateways []string

	//
	DHCPEnabled bool

	//
	DNSServers []string

	//
	IPAddresses []string

	//
	ProtocolIFType FailoverNetworkAdapterSettingData_ProtocolIFType

	//
	Subnets []string
}

func NewMsvm_FailoverNetworkAdapterSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_FailoverNetworkAdapterSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FailoverNetworkAdapterSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_FailoverNetworkAdapterSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FailoverNetworkAdapterSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FailoverNetworkAdapterSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetDefaultGateways sets the value of DefaultGateways for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertyDefaultGateways(value []string) (err error) {
	return instance.SetProperty("DefaultGateways", (value))
}

// GetDefaultGateways gets the value of DefaultGateways for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertyDefaultGateways() (value []string, err error) {
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
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", (value))
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertyDHCPEnabled() (value bool, err error) {
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
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertyDNSServers(value []string) (err error) {
	return instance.SetProperty("DNSServers", (value))
}

// GetDNSServers gets the value of DNSServers for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertyDNSServers() (value []string, err error) {
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

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", (value))
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertyIPAddresses() (value []string, err error) {
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

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertyProtocolIFType(value FailoverNetworkAdapterSettingData_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", (value))
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertyProtocolIFType() (value FailoverNetworkAdapterSettingData_ProtocolIFType, err error) {
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

	value = FailoverNetworkAdapterSettingData_ProtocolIFType(valuetmp)

	return
}

// SetSubnets sets the value of Subnets for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) SetPropertySubnets(value []string) (err error) {
	return instance.SetProperty("Subnets", (value))
}

// GetSubnets gets the value of Subnets for the instance
func (instance *Msvm_FailoverNetworkAdapterSettingData) GetPropertySubnets() (value []string, err error) {
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
