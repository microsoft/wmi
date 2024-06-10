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

// WebDavLocksSettings struct
type WebDavLocksSettings struct {
	*EmbeddedObject

	//
	Enabled bool

	//
	LockStore string

	//
	RequireLockForWriting bool
}

func NewWebDavLocksSettingsEx1(instance *cim.WmiInstance) (newInstance *WebDavLocksSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavLocksSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewWebDavLocksSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavLocksSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavLocksSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *WebDavLocksSettings) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *WebDavLocksSettings) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetLockStore sets the value of LockStore for the instance
func (instance *WebDavLocksSettings) SetPropertyLockStore(value string) (err error) {
	return instance.SetProperty("LockStore", (value))
}

// GetLockStore gets the value of LockStore for the instance
func (instance *WebDavLocksSettings) GetPropertyLockStore() (value string, err error) {
	retValue, err := instance.GetProperty("LockStore")
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

// SetRequireLockForWriting sets the value of RequireLockForWriting for the instance
func (instance *WebDavLocksSettings) SetPropertyRequireLockForWriting(value bool) (err error) {
	return instance.SetProperty("RequireLockForWriting", (value))
}

// GetRequireLockForWriting gets the value of RequireLockForWriting for the instance
func (instance *WebDavLocksSettings) GetPropertyRequireLockForWriting() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireLockForWriting")
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
