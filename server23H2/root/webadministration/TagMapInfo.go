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

// TagMapInfo struct
type TagMapInfo struct {
	*EmbeddedObject

	//
	TagMapping []TagMapElement
}

func NewTagMapInfoEx1(instance *cim.WmiInstance) (newInstance *TagMapInfo, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TagMapInfo{
		EmbeddedObject: tmp,
	}
	return
}

func NewTagMapInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TagMapInfo, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TagMapInfo{
		EmbeddedObject: tmp,
	}
	return
}

// SetTagMapping sets the value of TagMapping for the instance
func (instance *TagMapInfo) SetPropertyTagMapping(value []TagMapElement) (err error) {
	return instance.SetProperty("TagMapping", (value))
}

// GetTagMapping gets the value of TagMapping for the instance
func (instance *TagMapInfo) GetPropertyTagMapping() (value []TagMapElement, err error) {
	retValue, err := instance.GetProperty("TagMapping")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TagMapElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TagMapElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TagMapElement(valuetmp))
	}

	return
}
