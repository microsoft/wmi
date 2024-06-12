// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7344EDEA_49A9_449C_A6C2_37ECE9C9272F
//////////////////////////////////////////////
package ns7344edea_49a9_449c_a6c2_37ece9c9272f

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __EventProviderRegistration struct
type __EventProviderRegistration struct {
	*__ProviderRegistration

	//
	EventQueryList []string
}

func New__EventProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__EventProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__EventProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

// SetEventQueryList sets the value of EventQueryList for the instance
func (instance *__EventProviderRegistration) SetPropertyEventQueryList(value []string) (err error) {
	return instance.SetProperty("EventQueryList", (value))
}

// GetEventQueryList gets the value of EventQueryList for the instance
func (instance *__EventProviderRegistration) GetPropertyEventQueryList() (value []string, err error) {
	retValue, err := instance.GetProperty("EventQueryList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
