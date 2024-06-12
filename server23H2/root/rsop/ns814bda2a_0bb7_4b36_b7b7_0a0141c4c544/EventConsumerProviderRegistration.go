// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS814BDA2A_0BB7_4B36_B7B7_0A0141C4C544
//////////////////////////////////////////////
package ns814bda2a_0bb7_4b36_b7b7_0a0141c4c544

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
