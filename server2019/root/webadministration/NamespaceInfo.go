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

// NamespaceInfo struct
type NamespaceInfo struct {
	*EmbeddedObject

	//
	AutoImportVBNamespace bool

	//
	Namespaces []NamespaceElement
}

func NewNamespaceInfoEx1(instance *cim.WmiInstance) (newInstance *NamespaceInfo, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NamespaceInfo{
		EmbeddedObject: tmp,
	}
	return
}

func NewNamespaceInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NamespaceInfo, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NamespaceInfo{
		EmbeddedObject: tmp,
	}
	return
}

// SetAutoImportVBNamespace sets the value of AutoImportVBNamespace for the instance
func (instance *NamespaceInfo) SetPropertyAutoImportVBNamespace(value bool) (err error) {
	return instance.SetProperty("AutoImportVBNamespace", (value))
}

// GetAutoImportVBNamespace gets the value of AutoImportVBNamespace for the instance
func (instance *NamespaceInfo) GetPropertyAutoImportVBNamespace() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoImportVBNamespace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetNamespaces sets the value of Namespaces for the instance
func (instance *NamespaceInfo) SetPropertyNamespaces(value []NamespaceElement) (err error) {
	return instance.SetProperty("Namespaces", (value))
}

// GetNamespaces gets the value of Namespaces for the instance
func (instance *NamespaceInfo) GetPropertyNamespaces() (value []NamespaceElement, err error) {
	retValue, err := instance.GetProperty("Namespaces")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NamespaceElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NamespaceElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NamespaceElement(valuetmp))
	}

	return
}
