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

// TagPrefixInfo struct
type TagPrefixInfo struct {
	*EmbeddedObject

	//
	Controls []TagPrefixElement
}

func NewTagPrefixInfoEx1(instance *cim.WmiInstance) (newInstance *TagPrefixInfo, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TagPrefixInfo{
		EmbeddedObject: tmp,
	}
	return
}

func NewTagPrefixInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TagPrefixInfo, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TagPrefixInfo{
		EmbeddedObject: tmp,
	}
	return
}

// SetControls sets the value of Controls for the instance
func (instance *TagPrefixInfo) SetPropertyControls(value []TagPrefixElement) (err error) {
	return instance.SetProperty("Controls", (value))
}

// GetControls gets the value of Controls for the instance
func (instance *TagPrefixInfo) GetPropertyControls() (value []TagPrefixElement, err error) {
	retValue, err := instance.GetProperty("Controls")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TagPrefixElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TagPrefixElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TagPrefixElement(valuetmp))
	}

	return
}
