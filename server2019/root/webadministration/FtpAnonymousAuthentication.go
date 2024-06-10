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

// FtpAnonymousAuthentication struct
type FtpAnonymousAuthentication struct {
	*EmbeddedObject

	//
	DefaultLogonDomain string

	//
	Enabled bool

	//
	LogonMethod int32

	//
	Password string

	//
	UserName string
}

func NewFtpAnonymousAuthenticationEx1(instance *cim.WmiInstance) (newInstance *FtpAnonymousAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpAnonymousAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpAnonymousAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpAnonymousAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpAnonymousAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

// SetDefaultLogonDomain sets the value of DefaultLogonDomain for the instance
func (instance *FtpAnonymousAuthentication) SetPropertyDefaultLogonDomain(value string) (err error) {
	return instance.SetProperty("DefaultLogonDomain", (value))
}

// GetDefaultLogonDomain gets the value of DefaultLogonDomain for the instance
func (instance *FtpAnonymousAuthentication) GetPropertyDefaultLogonDomain() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultLogonDomain")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *FtpAnonymousAuthentication) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *FtpAnonymousAuthentication) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetLogonMethod sets the value of LogonMethod for the instance
func (instance *FtpAnonymousAuthentication) SetPropertyLogonMethod(value int32) (err error) {
	return instance.SetProperty("LogonMethod", (value))
}

// GetLogonMethod gets the value of LogonMethod for the instance
func (instance *FtpAnonymousAuthentication) GetPropertyLogonMethod() (value int32, err error) {
	retValue, err := instance.GetProperty("LogonMethod")
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

// SetPassword sets the value of Password for the instance
func (instance *FtpAnonymousAuthentication) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *FtpAnonymousAuthentication) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetUserName sets the value of UserName for the instance
func (instance *FtpAnonymousAuthentication) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *FtpAnonymousAuthentication) GetPropertyUserName() (value string, err error) {
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
