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

// HttpHandlerAction struct
type HttpHandlerAction struct {
	*CollectionElement

	//
	Path string

	//
	Type string

	//
	Validate bool

	//
	Verb string
}

func NewHttpHandlerActionEx1(instance *cim.WmiInstance) (newInstance *HttpHandlerAction, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpHandlerAction{
		CollectionElement: tmp,
	}
	return
}

func NewHttpHandlerActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpHandlerAction, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpHandlerAction{
		CollectionElement: tmp,
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *HttpHandlerAction) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *HttpHandlerAction) GetPropertyPath() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *HttpHandlerAction) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *HttpHandlerAction) GetPropertyType() (value string, err error) {
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

// SetValidate sets the value of Validate for the instance
func (instance *HttpHandlerAction) SetPropertyValidate(value bool) (err error) {
	return instance.SetProperty("Validate", (value))
}

// GetValidate gets the value of Validate for the instance
func (instance *HttpHandlerAction) GetPropertyValidate() (value bool, err error) {
	retValue, err := instance.GetProperty("Validate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetVerb sets the value of Verb for the instance
func (instance *HttpHandlerAction) SetPropertyVerb(value string) (err error) {
	return instance.SetProperty("Verb", (value))
}

// GetVerb gets the value of Verb for the instance
func (instance *HttpHandlerAction) GetPropertyVerb() (value string, err error) {
	retValue, err := instance.GetProperty("Verb")
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
