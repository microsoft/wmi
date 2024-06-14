// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS5B3DCE03_EB70_4EA2_8019_E0415C3C4C7B
//////////////////////////////////////////////
package ns5b3dce03_eb70_4ea2_8019_e0415c3c4c7b

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __EventConsumerProviderRegistration struct
type __EventConsumerProviderRegistration struct {
	*__ProviderRegistration

	//
	ConsumerClassNames []string
}

func New__EventConsumerProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__EventConsumerProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventConsumerProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__EventConsumerProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventConsumerProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventConsumerProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

// SetConsumerClassNames sets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) SetPropertyConsumerClassNames(value []string) (err error) {
	return instance.SetProperty("ConsumerClassNames", (value))
}

// GetConsumerClassNames gets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) GetPropertyConsumerClassNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ConsumerClassNames")
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