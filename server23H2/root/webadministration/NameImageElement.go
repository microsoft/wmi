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

// NameImageElement struct
type NameImageElement struct {
	*CollectionElement

	//
	Image string

	//
	Image32 string

	//
	Name string
}

func NewNameImageElementEx1(instance *cim.WmiInstance) (newInstance *NameImageElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NameImageElement{
		CollectionElement: tmp,
	}
	return
}

func NewNameImageElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NameImageElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NameImageElement{
		CollectionElement: tmp,
	}
	return
}

// SetImage sets the value of Image for the instance
func (instance *NameImageElement) SetPropertyImage(value string) (err error) {
	return instance.SetProperty("Image", (value))
}

// GetImage gets the value of Image for the instance
func (instance *NameImageElement) GetPropertyImage() (value string, err error) {
	retValue, err := instance.GetProperty("Image")
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

// SetImage32 sets the value of Image32 for the instance
func (instance *NameImageElement) SetPropertyImage32(value string) (err error) {
	return instance.SetProperty("Image32", (value))
}

// GetImage32 gets the value of Image32 for the instance
func (instance *NameImageElement) GetPropertyImage32() (value string, err error) {
	retValue, err := instance.GetProperty("Image32")
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

// SetName sets the value of Name for the instance
func (instance *NameImageElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *NameImageElement) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
