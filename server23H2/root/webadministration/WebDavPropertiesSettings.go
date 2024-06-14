// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WebDavPropertiesSettings struct
type WebDavPropertiesSettings struct {
	*EmbeddedObject

	//
	AllowAnonymousPropfind bool

	//
	AllowCustomProperties bool

	//
	AllowInfinitePropfindDepth bool

	//
	Properties []WebDavPropertiesElement
}

func NewWebDavPropertiesSettingsEx1(instance *cim.WmiInstance) (newInstance *WebDavPropertiesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavPropertiesSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewWebDavPropertiesSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavPropertiesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavPropertiesSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowAnonymousPropfind sets the value of AllowAnonymousPropfind for the instance
func (instance *WebDavPropertiesSettings) SetPropertyAllowAnonymousPropfind(value bool) (err error) {
	return instance.SetProperty("AllowAnonymousPropfind", (value))
}

// GetAllowAnonymousPropfind gets the value of AllowAnonymousPropfind for the instance
func (instance *WebDavPropertiesSettings) GetPropertyAllowAnonymousPropfind() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowAnonymousPropfind")
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

// SetAllowCustomProperties sets the value of AllowCustomProperties for the instance
func (instance *WebDavPropertiesSettings) SetPropertyAllowCustomProperties(value bool) (err error) {
	return instance.SetProperty("AllowCustomProperties", (value))
}

// GetAllowCustomProperties gets the value of AllowCustomProperties for the instance
func (instance *WebDavPropertiesSettings) GetPropertyAllowCustomProperties() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowCustomProperties")
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

// SetAllowInfinitePropfindDepth sets the value of AllowInfinitePropfindDepth for the instance
func (instance *WebDavPropertiesSettings) SetPropertyAllowInfinitePropfindDepth(value bool) (err error) {
	return instance.SetProperty("AllowInfinitePropfindDepth", (value))
}

// GetAllowInfinitePropfindDepth gets the value of AllowInfinitePropfindDepth for the instance
func (instance *WebDavPropertiesSettings) GetPropertyAllowInfinitePropfindDepth() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInfinitePropfindDepth")
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

// SetProperties sets the value of Properties for the instance
func (instance *WebDavPropertiesSettings) SetPropertyProperties(value []WebDavPropertiesElement) (err error) {
	return instance.SetProperty("Properties", (value))
}

// GetProperties gets the value of Properties for the instance
func (instance *WebDavPropertiesSettings) GetPropertyProperties() (value []WebDavPropertiesElement, err error) {
	retValue, err := instance.GetProperty("Properties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WebDavPropertiesElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WebDavPropertiesElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WebDavPropertiesElement(valuetmp))
	}

	return
}
