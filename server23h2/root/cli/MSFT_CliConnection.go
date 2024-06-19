// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CliConnection struct
type MSFT_CliConnection struct {
	*cim.WmiInstance

	//
	Authority string

	//
	Locale string

	//
	NameSpace string

	//
	Password string

	//
	Server string

	//
	User string
}

func NewMSFT_CliConnectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliConnection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliConnection{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliConnection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliConnection{
		WmiInstance: tmp,
	}
	return
}

// SetAuthority sets the value of Authority for the instance
func (instance *MSFT_CliConnection) SetPropertyAuthority(value string) (err error) {
	return instance.SetProperty("Authority", (value))
}

// GetAuthority gets the value of Authority for the instance
func (instance *MSFT_CliConnection) GetPropertyAuthority() (value string, err error) {
	retValue, err := instance.GetProperty("Authority")
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

// SetLocale sets the value of Locale for the instance
func (instance *MSFT_CliConnection) SetPropertyLocale(value string) (err error) {
	return instance.SetProperty("Locale", (value))
}

// GetLocale gets the value of Locale for the instance
func (instance *MSFT_CliConnection) GetPropertyLocale() (value string, err error) {
	retValue, err := instance.GetProperty("Locale")
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

// SetNameSpace sets the value of NameSpace for the instance
func (instance *MSFT_CliConnection) SetPropertyNameSpace(value string) (err error) {
	return instance.SetProperty("NameSpace", (value))
}

// GetNameSpace gets the value of NameSpace for the instance
func (instance *MSFT_CliConnection) GetPropertyNameSpace() (value string, err error) {
	retValue, err := instance.GetProperty("NameSpace")
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
func (instance *MSFT_CliConnection) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *MSFT_CliConnection) GetPropertyPassword() (value string, err error) {
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

// SetServer sets the value of Server for the instance
func (instance *MSFT_CliConnection) SetPropertyServer(value string) (err error) {
	return instance.SetProperty("Server", (value))
}

// GetServer gets the value of Server for the instance
func (instance *MSFT_CliConnection) GetPropertyServer() (value string, err error) {
	retValue, err := instance.GetProperty("Server")
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

// SetUser sets the value of User for the instance
func (instance *MSFT_CliConnection) SetPropertyUser(value string) (err error) {
	return instance.SetProperty("User", (value))
}

// GetUser gets the value of User for the instance
func (instance *MSFT_CliConnection) GetPropertyUser() (value string, err error) {
	retValue, err := instance.GetProperty("User")
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
