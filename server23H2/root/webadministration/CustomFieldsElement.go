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

// CustomFieldsElement struct
type CustomFieldsElement struct {
	*CollectionElement

	//
	LogFieldName string

	//
	SourceName string

	//
	SourceType int32
}

func NewCustomFieldsElementEx1(instance *cim.WmiInstance) (newInstance *CustomFieldsElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CustomFieldsElement{
		CollectionElement: tmp,
	}
	return
}

func NewCustomFieldsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CustomFieldsElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CustomFieldsElement{
		CollectionElement: tmp,
	}
	return
}

// SetLogFieldName sets the value of LogFieldName for the instance
func (instance *CustomFieldsElement) SetPropertyLogFieldName(value string) (err error) {
	return instance.SetProperty("LogFieldName", (value))
}

// GetLogFieldName gets the value of LogFieldName for the instance
func (instance *CustomFieldsElement) GetPropertyLogFieldName() (value string, err error) {
	retValue, err := instance.GetProperty("LogFieldName")
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

// SetSourceName sets the value of SourceName for the instance
func (instance *CustomFieldsElement) SetPropertySourceName(value string) (err error) {
	return instance.SetProperty("SourceName", (value))
}

// GetSourceName gets the value of SourceName for the instance
func (instance *CustomFieldsElement) GetPropertySourceName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceName")
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

// SetSourceType sets the value of SourceType for the instance
func (instance *CustomFieldsElement) SetPropertySourceType(value int32) (err error) {
	return instance.SetProperty("SourceType", (value))
}

// GetSourceType gets the value of SourceType for the instance
func (instance *CustomFieldsElement) GetPropertySourceType() (value int32, err error) {
	retValue, err := instance.GetProperty("SourceType")
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
