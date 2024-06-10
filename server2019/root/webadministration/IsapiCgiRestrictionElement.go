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

// IsapiCgiRestrictionElement struct
type IsapiCgiRestrictionElement struct {
	*CollectionElement

	//
	Allowed bool

	//
	Description string

	//
	GroupId string

	//
	Path string
}

func NewIsapiCgiRestrictionElementEx1(instance *cim.WmiInstance) (newInstance *IsapiCgiRestrictionElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IsapiCgiRestrictionElement{
		CollectionElement: tmp,
	}
	return
}

func NewIsapiCgiRestrictionElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IsapiCgiRestrictionElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IsapiCgiRestrictionElement{
		CollectionElement: tmp,
	}
	return
}

// SetAllowed sets the value of Allowed for the instance
func (instance *IsapiCgiRestrictionElement) SetPropertyAllowed(value bool) (err error) {
	return instance.SetProperty("Allowed", (value))
}

// GetAllowed gets the value of Allowed for the instance
func (instance *IsapiCgiRestrictionElement) GetPropertyAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("Allowed")
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

// SetDescription sets the value of Description for the instance
func (instance *IsapiCgiRestrictionElement) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *IsapiCgiRestrictionElement) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetGroupId sets the value of GroupId for the instance
func (instance *IsapiCgiRestrictionElement) SetPropertyGroupId(value string) (err error) {
	return instance.SetProperty("GroupId", (value))
}

// GetGroupId gets the value of GroupId for the instance
func (instance *IsapiCgiRestrictionElement) GetPropertyGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("GroupId")
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

// SetPath sets the value of Path for the instance
func (instance *IsapiCgiRestrictionElement) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *IsapiCgiRestrictionElement) GetPropertyPath() (value string, err error) {
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
