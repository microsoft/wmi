// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7CF930C4_E2B6_4AA7_B00E_814954412812
//////////////////////////////////////////////
package ns7cf930c4_e2b6_4aa7_b00e_814954412812

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __ProviderRegistration struct
type __ProviderRegistration struct {
	*__SystemClass

	//
	provider __Provider
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
func (instance *__ProviderRegistration) SetPropertyprovider(value __Provider) (err error) {
	return instance.SetProperty("provider", (value))
}

// Getprovider gets the value of provider for the instance
func (instance *__ProviderRegistration) GetPropertyprovider() (value __Provider, err error) {
	retValue, err := instance.GetProperty("provider")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(__Provider)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " __Provider is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = __Provider(valuetmp)

	return
}
