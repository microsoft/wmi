// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DnsClientPolicyConfiguration struct
type DnsClientPolicyConfiguration struct {
	*cim.WmiInstance

	//
	DirectAccessDnsServers []string

	//
	DirectAccessEnabled bool

	//
	DirectAccessIPsecCARestriction string

	//
	DirectAccessProxyName string

	//
	DirectAccessProxyType string

	//
	DirectAccessQueryIPsecEncryption string

	//
	DirectAccessQueryIPsecRequired bool

	//
	DnsSecIPsecCARestriction string

	//
	DnsSecQueryIPsecEncryption string

	//
	DnsSecQueryIPsecRequired bool

	//
	DnsSecValidationRequired bool

	//
	NameEncoding string

	//
	NameServers []string

	//
	Namespace string

	//
	QueryPolicy string

	//
	SecureNameQueryFallback string
}

func NewDnsClientPolicyConfigurationEx1(instance *cim.WmiInstance) (newInstance *DnsClientPolicyConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DnsClientPolicyConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewDnsClientPolicyConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DnsClientPolicyConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DnsClientPolicyConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetDirectAccessDnsServers sets the value of DirectAccessDnsServers for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessDnsServers(value []string) (err error) {
	return instance.SetProperty("DirectAccessDnsServers", (value))
}

// GetDirectAccessDnsServers gets the value of DirectAccessDnsServers for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessDnsServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DirectAccessDnsServers")
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

// SetDirectAccessEnabled sets the value of DirectAccessEnabled for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessEnabled(value bool) (err error) {
	return instance.SetProperty("DirectAccessEnabled", (value))
}

// GetDirectAccessEnabled gets the value of DirectAccessEnabled for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DirectAccessEnabled")
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

// SetDirectAccessIPsecCARestriction sets the value of DirectAccessIPsecCARestriction for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessIPsecCARestriction(value string) (err error) {
	return instance.SetProperty("DirectAccessIPsecCARestriction", (value))
}

// GetDirectAccessIPsecCARestriction gets the value of DirectAccessIPsecCARestriction for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessIPsecCARestriction() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessIPsecCARestriction")
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

// SetDirectAccessProxyName sets the value of DirectAccessProxyName for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessProxyName(value string) (err error) {
	return instance.SetProperty("DirectAccessProxyName", (value))
}

// GetDirectAccessProxyName gets the value of DirectAccessProxyName for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessProxyName() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessProxyName")
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

// SetDirectAccessProxyType sets the value of DirectAccessProxyType for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessProxyType(value string) (err error) {
	return instance.SetProperty("DirectAccessProxyType", (value))
}

// GetDirectAccessProxyType gets the value of DirectAccessProxyType for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessProxyType() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessProxyType")
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

// SetDirectAccessQueryIPsecEncryption sets the value of DirectAccessQueryIPsecEncryption for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessQueryIPsecEncryption(value string) (err error) {
	return instance.SetProperty("DirectAccessQueryIPsecEncryption", (value))
}

// GetDirectAccessQueryIPsecEncryption gets the value of DirectAccessQueryIPsecEncryption for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessQueryIPsecEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessQueryIPsecEncryption")
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

// SetDirectAccessQueryIPsecRequired sets the value of DirectAccessQueryIPsecRequired for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDirectAccessQueryIPsecRequired(value bool) (err error) {
	return instance.SetProperty("DirectAccessQueryIPsecRequired", (value))
}

// GetDirectAccessQueryIPsecRequired gets the value of DirectAccessQueryIPsecRequired for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDirectAccessQueryIPsecRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DirectAccessQueryIPsecRequired")
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

// SetDnsSecIPsecCARestriction sets the value of DnsSecIPsecCARestriction for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDnsSecIPsecCARestriction(value string) (err error) {
	return instance.SetProperty("DnsSecIPsecCARestriction", (value))
}

// GetDnsSecIPsecCARestriction gets the value of DnsSecIPsecCARestriction for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDnsSecIPsecCARestriction() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSecIPsecCARestriction")
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

// SetDnsSecQueryIPsecEncryption sets the value of DnsSecQueryIPsecEncryption for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDnsSecQueryIPsecEncryption(value string) (err error) {
	return instance.SetProperty("DnsSecQueryIPsecEncryption", (value))
}

// GetDnsSecQueryIPsecEncryption gets the value of DnsSecQueryIPsecEncryption for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDnsSecQueryIPsecEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSecQueryIPsecEncryption")
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

// SetDnsSecQueryIPsecRequired sets the value of DnsSecQueryIPsecRequired for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDnsSecQueryIPsecRequired(value bool) (err error) {
	return instance.SetProperty("DnsSecQueryIPsecRequired", (value))
}

// GetDnsSecQueryIPsecRequired gets the value of DnsSecQueryIPsecRequired for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDnsSecQueryIPsecRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsSecQueryIPsecRequired")
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

// SetDnsSecValidationRequired sets the value of DnsSecValidationRequired for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyDnsSecValidationRequired(value bool) (err error) {
	return instance.SetProperty("DnsSecValidationRequired", (value))
}

// GetDnsSecValidationRequired gets the value of DnsSecValidationRequired for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyDnsSecValidationRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsSecValidationRequired")
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

// SetNameEncoding sets the value of NameEncoding for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyNameEncoding(value string) (err error) {
	return instance.SetProperty("NameEncoding", (value))
}

// GetNameEncoding gets the value of NameEncoding for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyNameEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("NameEncoding")
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

// SetNameServers sets the value of NameServers for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyNameServers(value []string) (err error) {
	return instance.SetProperty("NameServers", (value))
}

// GetNameServers gets the value of NameServers for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyNameServers() (value []string, err error) {
	retValue, err := instance.GetProperty("NameServers")
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

// SetNamespace sets the value of Namespace for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
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

// SetQueryPolicy sets the value of QueryPolicy for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertyQueryPolicy(value string) (err error) {
	return instance.SetProperty("QueryPolicy", (value))
}

// GetQueryPolicy gets the value of QueryPolicy for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertyQueryPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("QueryPolicy")
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

// SetSecureNameQueryFallback sets the value of SecureNameQueryFallback for the instance
func (instance *DnsClientPolicyConfiguration) SetPropertySecureNameQueryFallback(value string) (err error) {
	return instance.SetProperty("SecureNameQueryFallback", (value))
}

// GetSecureNameQueryFallback gets the value of SecureNameQueryFallback for the instance
func (instance *DnsClientPolicyConfiguration) GetPropertySecureNameQueryFallback() (value string, err error) {
	retValue, err := instance.GetProperty("SecureNameQueryFallback")
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
