// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpIpSecuritySection struct
type FtpIpSecuritySection struct {
	*ConfigurationSectionWithCollection

	//
	AllowUnlisted bool

	//
	EnableReverseDns bool

	//
	IpSecurity []IPAddressFilterElement
}

func NewFtpIpSecuritySectionEx1(instance *cim.WmiInstance) (newInstance *FtpIpSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpIpSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFtpIpSecuritySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpIpSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpIpSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *FtpIpSecuritySection) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *FtpIpSecuritySection) GetPropertyAllowUnlisted() (value bool, err error) {
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

// SetEnableReverseDns sets the value of EnableReverseDns for the instance
func (instance *FtpIpSecuritySection) SetPropertyEnableReverseDns(value bool) (err error) {
	return instance.SetProperty("EnableReverseDns", (value))
}

// GetEnableReverseDns gets the value of EnableReverseDns for the instance
func (instance *FtpIpSecuritySection) GetPropertyEnableReverseDns() (value bool, err error) {
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
func (instance *FtpIpSecuritySection) SetPropertyIpSecurity(value []IPAddressFilterElement) (err error) {
	return instance.SetProperty("IpSecurity", (value))
}

// GetIpSecurity gets the value of IpSecurity for the instance
func (instance *FtpIpSecuritySection) GetPropertyIpSecurity() (value []IPAddressFilterElement, err error) {
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
