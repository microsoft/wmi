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

// TagPrefixElement struct
type TagPrefixElement struct {
	*CollectionElement

	//
	Assembly string

	//
	Namespace string

	//
	Src string

	//
	TagName string

	//
	TagPrefix string
}

func NewTagPrefixElementEx1(instance *cim.WmiInstance) (newInstance *TagPrefixElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TagPrefixElement{
		CollectionElement: tmp,
	}
	return
}

func NewTagPrefixElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TagPrefixElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TagPrefixElement{
		CollectionElement: tmp,
	}
	return
}

// SetAssembly sets the value of Assembly for the instance
func (instance *TagPrefixElement) SetPropertyAssembly(value string) (err error) {
	return instance.SetProperty("Assembly", (value))
}

// GetAssembly gets the value of Assembly for the instance
func (instance *TagPrefixElement) GetPropertyAssembly() (value string, err error) {
	retValue, err := instance.GetProperty("Assembly")
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

// SetNamespace sets the value of Namespace for the instance
func (instance *TagPrefixElement) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *TagPrefixElement) GetPropertyNamespace() (value string, err error) {
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

// SetSrc sets the value of Src for the instance
func (instance *TagPrefixElement) SetPropertySrc(value string) (err error) {
	return instance.SetProperty("Src", (value))
}

// GetSrc gets the value of Src for the instance
func (instance *TagPrefixElement) GetPropertySrc() (value string, err error) {
	retValue, err := instance.GetProperty("Src")
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

// SetTagName sets the value of TagName for the instance
func (instance *TagPrefixElement) SetPropertyTagName(value string) (err error) {
	return instance.SetProperty("TagName", (value))
}

// GetTagName gets the value of TagName for the instance
func (instance *TagPrefixElement) GetPropertyTagName() (value string, err error) {
	retValue, err := instance.GetProperty("TagName")
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

// SetTagPrefix sets the value of TagPrefix for the instance
func (instance *TagPrefixElement) SetPropertyTagPrefix(value string) (err error) {
	return instance.SetProperty("TagPrefix", (value))
}

// GetTagPrefix gets the value of TagPrefix for the instance
func (instance *TagPrefixElement) GetPropertyTagPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("TagPrefix")
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
