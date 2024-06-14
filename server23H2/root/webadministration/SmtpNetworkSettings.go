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

// SmtpNetworkSettings struct
type SmtpNetworkSettings struct {
	*EmbeddedObject

	//
	DefaultCredentials bool

	//
	Host string

	//
	Password string

	//
	Port int32

	//
	UserName string
}

func NewSmtpNetworkSettingsEx1(instance *cim.WmiInstance) (newInstance *SmtpNetworkSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SmtpNetworkSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSmtpNetworkSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SmtpNetworkSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SmtpNetworkSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDefaultCredentials sets the value of DefaultCredentials for the instance
func (instance *SmtpNetworkSettings) SetPropertyDefaultCredentials(value bool) (err error) {
	return instance.SetProperty("DefaultCredentials", (value))
}

// GetDefaultCredentials gets the value of DefaultCredentials for the instance
func (instance *SmtpNetworkSettings) GetPropertyDefaultCredentials() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultCredentials")
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

// SetHost sets the value of Host for the instance
func (instance *SmtpNetworkSettings) SetPropertyHost(value string) (err error) {
	return instance.SetProperty("Host", (value))
}

// GetHost gets the value of Host for the instance
func (instance *SmtpNetworkSettings) GetPropertyHost() (value string, err error) {
	retValue, err := instance.GetProperty("Host")
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

// SetPassword sets the value of Password for the instance
func (instance *SmtpNetworkSettings) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *SmtpNetworkSettings) GetPropertyPassword() (value string, err error) {
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

// SetPort sets the value of Port for the instance
func (instance *SmtpNetworkSettings) SetPropertyPort(value int32) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *SmtpNetworkSettings) GetPropertyPort() (value int32, err error) {
	retValue, err := instance.GetProperty("Port")
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

// SetUserName sets the value of UserName for the instance
func (instance *SmtpNetworkSettings) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *SmtpNetworkSettings) GetPropertyUserName() (value string, err error) {
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
