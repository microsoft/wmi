// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
	root "server23h2/root"
)

// __ProviderRegistration struct
type __ProviderRegistration struct {
	*__SystemClass

	//
	provider root.__Provider
}

func New__ProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__ProviderRegistration, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ProviderRegistration{
		__SystemClass: tmp,
	}
	return
}

func New__ProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ProviderRegistration, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ProviderRegistration{
		__SystemClass: tmp,
	}
	return
}

// Setprovider sets the value of provider for the instance
func (instance *__ProviderRegistration) SetPropertyprovider(value root.__Provider) (err error) {
	return instance.SetProperty("provider", (value))
}

// Getprovider gets the value of provider for the instance
func (instance *__ProviderRegistration) GetPropertyprovider() (value root.__Provider, err error) {
	retValue, err := instance.GetProperty("provider")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(root.__Provider)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " root.__Provider is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = root.__Provider(valuetmp)

	return
}
