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

// AuthorizationRule struct
type AuthorizationRule struct {
	*CollectionElement

	//
	AccessType int32

	//
	Roles string

	//
	Users string

	//
	Verbs string
}

func NewAuthorizationRuleEx1(instance *cim.WmiInstance) (newInstance *AuthorizationRule, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AuthorizationRule{
		CollectionElement: tmp,
	}
	return
}

func NewAuthorizationRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AuthorizationRule, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AuthorizationRule{
		CollectionElement: tmp,
	}
	return
}

// SetAccessType sets the value of AccessType for the instance
func (instance *AuthorizationRule) SetPropertyAccessType(value int32) (err error) {
	return instance.SetProperty("AccessType", (value))
}

// GetAccessType gets the value of AccessType for the instance
func (instance *AuthorizationRule) GetPropertyAccessType() (value int32, err error) {
	retValue, err := instance.GetProperty("AccessType")
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

// SetRoles sets the value of Roles for the instance
func (instance *AuthorizationRule) SetPropertyRoles(value string) (err error) {
	return instance.SetProperty("Roles", (value))
}

// GetRoles gets the value of Roles for the instance
func (instance *AuthorizationRule) GetPropertyRoles() (value string, err error) {
	retValue, err := instance.GetProperty("Roles")
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

// SetUsers sets the value of Users for the instance
func (instance *AuthorizationRule) SetPropertyUsers(value string) (err error) {
	return instance.SetProperty("Users", (value))
}

// GetUsers gets the value of Users for the instance
func (instance *AuthorizationRule) GetPropertyUsers() (value string, err error) {
	retValue, err := instance.GetProperty("Users")
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

// SetVerbs sets the value of Verbs for the instance
func (instance *AuthorizationRule) SetPropertyVerbs(value string) (err error) {
	return instance.SetProperty("Verbs", (value))
}

// GetVerbs gets the value of Verbs for the instance
func (instance *AuthorizationRule) GetPropertyVerbs() (value string, err error) {
	retValue, err := instance.GetProperty("Verbs")
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
