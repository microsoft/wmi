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

// WebDavMimeTypeElement struct
type WebDavMimeTypeElement struct {
	*CollectionElement

	//
	Access int32

	//
	Path string

	//
	Roles string

	//
	Users string
}

func NewWebDavMimeTypeElementEx1(instance *cim.WmiInstance) (newInstance *WebDavMimeTypeElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavMimeTypeElement{
		CollectionElement: tmp,
	}
	return
}

func NewWebDavMimeTypeElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavMimeTypeElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavMimeTypeElement{
		CollectionElement: tmp,
	}
	return
}

// SetAccess sets the value of Access for the instance
func (instance *WebDavMimeTypeElement) SetPropertyAccess(value int32) (err error) {
	return instance.SetProperty("Access", (value))
}

// GetAccess gets the value of Access for the instance
func (instance *WebDavMimeTypeElement) GetPropertyAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("Access")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *WebDavMimeTypeElement) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *WebDavMimeTypeElement) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetRoles sets the value of Roles for the instance
func (instance *WebDavMimeTypeElement) SetPropertyRoles(value string) (err error) {
	return instance.SetProperty("Roles", (value))
}

// GetRoles gets the value of Roles for the instance
func (instance *WebDavMimeTypeElement) GetPropertyRoles() (value string, err error) {
	retValue, err := instance.GetProperty("Roles")
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

// SetUsers sets the value of Users for the instance
func (instance *WebDavMimeTypeElement) SetPropertyUsers(value string) (err error) {
	return instance.SetProperty("Users", (value))
}

// GetUsers gets the value of Users for the instance
func (instance *WebDavMimeTypeElement) GetPropertyUsers() (value string, err error) {
	retValue, err := instance.GetProperty("Users")
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
