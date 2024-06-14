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

// GlobalModuleElement struct
type GlobalModuleElement struct {
	*CollectionElement

	//
	Image string

	//
	Name string

	//
	PreCondition string
}

func NewGlobalModuleElementEx1(instance *cim.WmiInstance) (newInstance *GlobalModuleElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &GlobalModuleElement{
		CollectionElement: tmp,
	}
	return
}

func NewGlobalModuleElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GlobalModuleElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GlobalModuleElement{
		CollectionElement: tmp,
	}
	return
}

// SetImage sets the value of Image for the instance
func (instance *GlobalModuleElement) SetPropertyImage(value string) (err error) {
	return instance.SetProperty("Image", (value))
}

// GetImage gets the value of Image for the instance
func (instance *GlobalModuleElement) GetPropertyImage() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *GlobalModuleElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *GlobalModuleElement) GetPropertyName() (value string, err error) {
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

// SetPreCondition sets the value of PreCondition for the instance
func (instance *GlobalModuleElement) SetPropertyPreCondition(value string) (err error) {
	return instance.SetProperty("PreCondition", (value))
}

// GetPreCondition gets the value of PreCondition for the instance
func (instance *GlobalModuleElement) GetPropertyPreCondition() (value string, err error) {
	retValue, err := instance.GetProperty("PreCondition")
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
