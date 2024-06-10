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

// WebDavLockStoresSettings struct
type WebDavLockStoresSettings struct {
	*EmbeddedObject

	//
	LockStores []NameImageElement
}

func NewWebDavLockStoresSettingsEx1(instance *cim.WmiInstance) (newInstance *WebDavLockStoresSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavLockStoresSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewWebDavLockStoresSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavLockStoresSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavLockStoresSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetLockStores sets the value of LockStores for the instance
func (instance *WebDavLockStoresSettings) SetPropertyLockStores(value []NameImageElement) (err error) {
	return instance.SetProperty("LockStores", (value))
}

// GetLockStores gets the value of LockStores for the instance
func (instance *WebDavLockStoresSettings) GetPropertyLockStores() (value []NameImageElement, err error) {
	retValue, err := instance.GetProperty("LockStores")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NameImageElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NameImageElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NameImageElement(valuetmp))
	}

	return
}
