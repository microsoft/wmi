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

// SoapExtensionTypeElement struct
type SoapExtensionTypeElement struct {
	*CollectionElement

	//
	Group int32

	//
	Priority int32

	//
	Type string
}

func NewSoapExtensionTypeElementEx1(instance *cim.WmiInstance) (newInstance *SoapExtensionTypeElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionTypeElement{
		CollectionElement: tmp,
	}
	return
}

func NewSoapExtensionTypeElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SoapExtensionTypeElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionTypeElement{
		CollectionElement: tmp,
	}
	return
}

// SetGroup sets the value of Group for the instance
func (instance *SoapExtensionTypeElement) SetPropertyGroup(value int32) (err error) {
	return instance.SetProperty("Group", (value))
}

// GetGroup gets the value of Group for the instance
func (instance *SoapExtensionTypeElement) GetPropertyGroup() (value int32, err error) {
	retValue, err := instance.GetProperty("Group")
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

// SetPriority sets the value of Priority for the instance
func (instance *SoapExtensionTypeElement) SetPropertyPriority(value int32) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *SoapExtensionTypeElement) GetPropertyPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("Priority")
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

// SetType sets the value of Type for the instance
func (instance *SoapExtensionTypeElement) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *SoapExtensionTypeElement) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
