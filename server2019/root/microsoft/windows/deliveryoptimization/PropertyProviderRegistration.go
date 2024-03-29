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

// __PropertyProviderRegistration struct
type __PropertyProviderRegistration struct {
	*__ProviderRegistration

	//
	SupportsGet bool

	//
	SupportsPut bool
}

func New__PropertyProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__PropertyProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__PropertyProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__PropertyProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__PropertyProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__PropertyProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

// SetSupportsGet sets the value of SupportsGet for the instance
func (instance *__PropertyProviderRegistration) SetPropertySupportsGet(value bool) (err error) {
	return instance.SetProperty("SupportsGet", value)
}

// GetSupportsGet gets the value of SupportsGet for the instance
func (instance *__PropertyProviderRegistration) GetPropertySupportsGet() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsGet")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsPut sets the value of SupportsPut for the instance
func (instance *__PropertyProviderRegistration) SetPropertySupportsPut(value bool) (err error) {
	return instance.SetProperty("SupportsPut", value)
}

// GetSupportsPut gets the value of SupportsPut for the instance
func (instance *__PropertyProviderRegistration) GetPropertySupportsPut() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsPut")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
