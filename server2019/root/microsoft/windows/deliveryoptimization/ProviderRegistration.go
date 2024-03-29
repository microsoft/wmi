// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("provider", value)
}

// Getprovider gets the value of provider for the instance
func (instance *__ProviderRegistration) GetPropertyprovider() (value __Provider, err error) {
	retValue, err := instance.GetProperty("provider")
	if err != nil {
		return
	}
	value, ok := retValue.(__Provider)
	if !ok {
		// TODO: Set an error
	}
	return
}
