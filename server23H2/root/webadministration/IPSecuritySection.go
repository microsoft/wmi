// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// IPSecuritySection struct
type IPSecuritySection struct {
	*ConfigurationSectionWithCollection

	//
	AllowUnlisted bool

	//
	DenyAction int32

	//
	EnableProxyMode bool

	//
	EnableReverseDns bool

	//
	IpSecurity []IPAddressFilterElement
}

func NewIPSecuritySectionEx1(instance *cim.WmiInstance) (newInstance *IPSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IPSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewIPSecuritySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IPSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IPSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *IPSecuritySection) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *IPSecuritySection) GetPropertyAllowUnlisted() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnlisted")
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

// SetDenyAction sets the value of DenyAction for the instance
func (instance *IPSecuritySection) SetPropertyDenyAction(value int32) (err error) {
	return instance.SetProperty("DenyAction", (value))
}

// GetDenyAction gets the value of DenyAction for the instance
func (instance *IPSecuritySection) GetPropertyDenyAction() (value int32, err error) {
	retValue, err := instance.GetProperty("DenyAction")
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

	value = int32(valuetmp)

	return
}

// SetEnableProxyMode sets the value of EnableProxyMode for the instance
func (instance *IPSecuritySection) SetPropertyEnableProxyMode(value bool) (err error) {
	return instance.SetProperty("EnableProxyMode", (value))
}

// GetEnableProxyMode gets the value of EnableProxyMode for the instance
func (instance *IPSecuritySection) GetPropertyEnableProxyMode() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableProxyMode")
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

// SetEnableReverseDns sets the value of EnableReverseDns for the instance
func (instance *IPSecuritySection) SetPropertyEnableReverseDns(value bool) (err error) {
	return instance.SetProperty("EnableReverseDns", (value))
}

// GetEnableReverseDns gets the value of EnableReverseDns for the instance
func (instance *IPSecuritySection) GetPropertyEnableReverseDns() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableReverseDns")
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

// SetIpSecurity sets the value of IpSecurity for the instance
func (instance *IPSecuritySection) SetPropertyIpSecurity(value []IPAddressFilterElement) (err error) {
	return instance.SetProperty("IpSecurity", (value))
}

// GetIpSecurity gets the value of IpSecurity for the instance
func (instance *IPSecuritySection) GetPropertyIpSecurity() (value []IPAddressFilterElement, err error) {
	retValue, err := instance.GetProperty("IpSecurity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(IPAddressFilterElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " IPAddressFilterElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, IPAddressFilterElement(valuetmp))
	}

	return
}
