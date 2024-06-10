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

// TagMapElement struct
type TagMapElement struct {
	*CollectionElement

	//
	MappedTagType string

	//
	TagType string
}

func NewTagMapElementEx1(instance *cim.WmiInstance) (newInstance *TagMapElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TagMapElement{
		CollectionElement: tmp,
	}
	return
}

func NewTagMapElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TagMapElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TagMapElement{
		CollectionElement: tmp,
	}
	return
}

// SetMappedTagType sets the value of MappedTagType for the instance
func (instance *TagMapElement) SetPropertyMappedTagType(value string) (err error) {
	return instance.SetProperty("MappedTagType", (value))
}

// GetMappedTagType gets the value of MappedTagType for the instance
func (instance *TagMapElement) GetPropertyMappedTagType() (value string, err error) {
	retValue, err := instance.GetProperty("MappedTagType")
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

// SetTagType sets the value of TagType for the instance
func (instance *TagMapElement) SetPropertyTagType(value string) (err error) {
	return instance.SetProperty("TagType", (value))
}

// GetTagType gets the value of TagType for the instance
func (instance *TagMapElement) GetPropertyTagType() (value string, err error) {
	retValue, err := instance.GetProperty("TagType")
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
