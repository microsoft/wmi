// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root
//////////////////////////////////////////////
package root

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __ListOfEventActiveNamespaces struct
type __ListOfEventActiveNamespaces struct {
	*__SystemClass

	//
	Namespaces []string
}

func New__ListOfEventActiveNamespacesEx1(instance *cim.WmiInstance) (newInstance *__ListOfEventActiveNamespaces, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ListOfEventActiveNamespaces{
		__SystemClass: tmp,
	}
	return
}

func New__ListOfEventActiveNamespacesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ListOfEventActiveNamespaces, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ListOfEventActiveNamespaces{
		__SystemClass: tmp,
	}
	return
}

// SetNamespaces sets the value of Namespaces for the instance
func (instance *__ListOfEventActiveNamespaces) SetPropertyNamespaces(value []string) (err error) {
	return instance.SetProperty("Namespaces", (value))
}

// GetNamespaces gets the value of Namespaces for the instance
func (instance *__ListOfEventActiveNamespaces) GetPropertyNamespaces() (value []string, err error) {
	retValue, err := instance.GetProperty("Namespaces")
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
