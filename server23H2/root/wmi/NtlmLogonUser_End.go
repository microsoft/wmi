// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NtlmLogonUser_End struct
type NtlmLogonUser_End struct {
	*NtlmLogonUser

	// Domain Name
	DomainName string

	// Logon Type
	LogonType uint32

	// Status
	Status uint32

	// User Name
	UserName string
}

func NewNtlmLogonUser_EndEx1(instance *cim.WmiInstance) (newInstance *NtlmLogonUser_End, err error) {
	tmp, err := NewNtlmLogonUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser_End{
		NtlmLogonUser: tmp,
	}
	return
}

func NewNtlmLogonUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmLogonUser_End, err error) {
	tmp, err := NewNtlmLogonUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser_End{
		NtlmLogonUser: tmp,
	}
	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *NtlmLogonUser_End) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *NtlmLogonUser_End) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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

// SetLogonType sets the value of LogonType for the instance
func (instance *NtlmLogonUser_End) SetPropertyLogonType(value uint32) (err error) {
	return instance.SetProperty("LogonType", (value))
}

// GetLogonType gets the value of LogonType for the instance
func (instance *NtlmLogonUser_End) GetPropertyLogonType() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogonType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *NtlmLogonUser_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *NtlmLogonUser_End) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUserName sets the value of UserName for the instance
func (instance *NtlmLogonUser_End) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *NtlmLogonUser_End) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
