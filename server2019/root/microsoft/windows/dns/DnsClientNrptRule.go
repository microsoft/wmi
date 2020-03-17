// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DnsClientNrptRule struct
type DnsClientNrptRule struct {
	cim.WmiInstance

	//
	Comment string

	//
	DirectAccessDnsServers []string

	//
	DirectAccessEnabled bool

	//
	DirectAccessProxyName string

	//
	DirectAccessProxyType string

	//
	DirectAccessQueryIPsecEncryption string

	//
	DirectAccessQueryIPsecRequired bool

	//
	DisplayName string

	//
	DnsSecEnabled bool

	//
	DnsSecQueryIPsecEncryption string

	//
	DnsSecQueryIPsecRequired bool

	//
	DnsSecValidationRequired bool

	//
	IPsecCARestriction string

	//
	Name string

	//
	NameEncoding string

	//
	NameServers []string

	//
	Namespace []string

	//
	Version uint32
}

// SetComment sets the value of Comment for the instance
func (instance *DnsClientNrptRule) SetPropertyComment(value string) (err error) {
	return instance.SetProperty("Comment", value)
}

// GetComment gets the value of Comment for the instance
func (instance *DnsClientNrptRule) GetPropertyComment() (value string, err error) {
	retValue, err := instance.GetProperty("Comment")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessDnsServers sets the value of DirectAccessDnsServers for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessDnsServers(value []string) (err error) {
	return instance.SetProperty("DirectAccessDnsServers", value)
}

// GetDirectAccessDnsServers gets the value of DirectAccessDnsServers for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessDnsServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DirectAccessDnsServers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessEnabled sets the value of DirectAccessEnabled for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessEnabled(value bool) (err error) {
	return instance.SetProperty("DirectAccessEnabled", value)
}

// GetDirectAccessEnabled gets the value of DirectAccessEnabled for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DirectAccessEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessProxyName sets the value of DirectAccessProxyName for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessProxyName(value string) (err error) {
	return instance.SetProperty("DirectAccessProxyName", value)
}

// GetDirectAccessProxyName gets the value of DirectAccessProxyName for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessProxyName() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessProxyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessProxyType sets the value of DirectAccessProxyType for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessProxyType(value string) (err error) {
	return instance.SetProperty("DirectAccessProxyType", value)
}

// GetDirectAccessProxyType gets the value of DirectAccessProxyType for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessProxyType() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessProxyType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessQueryIPsecEncryption sets the value of DirectAccessQueryIPsecEncryption for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessQueryIPsecEncryption(value string) (err error) {
	return instance.SetProperty("DirectAccessQueryIPsecEncryption", value)
}

// GetDirectAccessQueryIPsecEncryption gets the value of DirectAccessQueryIPsecEncryption for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessQueryIPsecEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("DirectAccessQueryIPsecEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectAccessQueryIPsecRequired sets the value of DirectAccessQueryIPsecRequired for the instance
func (instance *DnsClientNrptRule) SetPropertyDirectAccessQueryIPsecRequired(value bool) (err error) {
	return instance.SetProperty("DirectAccessQueryIPsecRequired", value)
}

// GetDirectAccessQueryIPsecRequired gets the value of DirectAccessQueryIPsecRequired for the instance
func (instance *DnsClientNrptRule) GetPropertyDirectAccessQueryIPsecRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DirectAccessQueryIPsecRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *DnsClientNrptRule) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *DnsClientNrptRule) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSecEnabled sets the value of DnsSecEnabled for the instance
func (instance *DnsClientNrptRule) SetPropertyDnsSecEnabled(value bool) (err error) {
	return instance.SetProperty("DnsSecEnabled", value)
}

// GetDnsSecEnabled gets the value of DnsSecEnabled for the instance
func (instance *DnsClientNrptRule) GetPropertyDnsSecEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsSecEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSecQueryIPsecEncryption sets the value of DnsSecQueryIPsecEncryption for the instance
func (instance *DnsClientNrptRule) SetPropertyDnsSecQueryIPsecEncryption(value string) (err error) {
	return instance.SetProperty("DnsSecQueryIPsecEncryption", value)
}

// GetDnsSecQueryIPsecEncryption gets the value of DnsSecQueryIPsecEncryption for the instance
func (instance *DnsClientNrptRule) GetPropertyDnsSecQueryIPsecEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSecQueryIPsecEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSecQueryIPsecRequired sets the value of DnsSecQueryIPsecRequired for the instance
func (instance *DnsClientNrptRule) SetPropertyDnsSecQueryIPsecRequired(value bool) (err error) {
	return instance.SetProperty("DnsSecQueryIPsecRequired", value)
}

// GetDnsSecQueryIPsecRequired gets the value of DnsSecQueryIPsecRequired for the instance
func (instance *DnsClientNrptRule) GetPropertyDnsSecQueryIPsecRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsSecQueryIPsecRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSecValidationRequired sets the value of DnsSecValidationRequired for the instance
func (instance *DnsClientNrptRule) SetPropertyDnsSecValidationRequired(value bool) (err error) {
	return instance.SetProperty("DnsSecValidationRequired", value)
}

// GetDnsSecValidationRequired gets the value of DnsSecValidationRequired for the instance
func (instance *DnsClientNrptRule) GetPropertyDnsSecValidationRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsSecValidationRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPsecCARestriction sets the value of IPsecCARestriction for the instance
func (instance *DnsClientNrptRule) SetPropertyIPsecCARestriction(value string) (err error) {
	return instance.SetProperty("IPsecCARestriction", value)
}

// GetIPsecCARestriction gets the value of IPsecCARestriction for the instance
func (instance *DnsClientNrptRule) GetPropertyIPsecCARestriction() (value string, err error) {
	retValue, err := instance.GetProperty("IPsecCARestriction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *DnsClientNrptRule) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *DnsClientNrptRule) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameEncoding sets the value of NameEncoding for the instance
func (instance *DnsClientNrptRule) SetPropertyNameEncoding(value string) (err error) {
	return instance.SetProperty("NameEncoding", value)
}

// GetNameEncoding gets the value of NameEncoding for the instance
func (instance *DnsClientNrptRule) GetPropertyNameEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("NameEncoding")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameServers sets the value of NameServers for the instance
func (instance *DnsClientNrptRule) SetPropertyNameServers(value []string) (err error) {
	return instance.SetProperty("NameServers", value)
}

// GetNameServers gets the value of NameServers for the instance
func (instance *DnsClientNrptRule) GetPropertyNameServers() (value []string, err error) {
	retValue, err := instance.GetProperty("NameServers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *DnsClientNrptRule) SetPropertyNamespace(value []string) (err error) {
	return instance.SetProperty("Namespace", value)
}

// GetNamespace gets the value of Namespace for the instance
func (instance *DnsClientNrptRule) GetPropertyNamespace() (value []string, err error) {
	retValue, err := instance.GetProperty("Namespace")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *DnsClientNrptRule) SetPropertyVersion(value uint32) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *DnsClientNrptRule) GetPropertyVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
