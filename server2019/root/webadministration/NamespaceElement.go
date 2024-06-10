// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NamespaceElement struct
type NamespaceElement struct {
	*CollectionElement

	//
	Namespace string
}

func NewNamespaceElementEx1(instance *cim.WmiInstance) (newInstance *NamespaceElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NamespaceElement{
		CollectionElement: tmp,
	}
	return
}

func NewNamespaceElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NamespaceElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NamespaceElement{
		CollectionElement: tmp,
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *NamespaceElement) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *NamespaceElement) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
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
