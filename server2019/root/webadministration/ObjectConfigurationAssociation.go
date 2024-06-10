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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ObjectConfigurationAssociation struct
type ObjectConfigurationAssociation struct {
	*cim.WmiInstance

	//
	ConfigurationSection interface{}

	//
	ConfiguredObject interface{}
}

func NewObjectConfigurationAssociationEx1(instance *cim.WmiInstance) (newInstance *ObjectConfigurationAssociation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ObjectConfigurationAssociation{
		WmiInstance: tmp,
	}
	return
}

func NewObjectConfigurationAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObjectConfigurationAssociation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObjectConfigurationAssociation{
		WmiInstance: tmp,
	}
	return
}

// SetConfigurationSection sets the value of ConfigurationSection for the instance
func (instance *ObjectConfigurationAssociation) SetPropertyConfigurationSection(value interface{}) (err error) {
	return instance.SetProperty("ConfigurationSection", (value))
}

// GetConfigurationSection gets the value of ConfigurationSection for the instance
func (instance *ObjectConfigurationAssociation) GetPropertyConfigurationSection() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ConfigurationSection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetConfiguredObject sets the value of ConfiguredObject for the instance
func (instance *ObjectConfigurationAssociation) SetPropertyConfiguredObject(value interface{}) (err error) {
	return instance.SetProperty("ConfiguredObject", (value))
}

// GetConfiguredObject gets the value of ConfiguredObject for the instance
func (instance *ObjectConfigurationAssociation) GetPropertyConfiguredObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ConfiguredObject")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
