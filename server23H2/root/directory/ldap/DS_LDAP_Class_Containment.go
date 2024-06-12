// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DS_LDAP_Class_Containment struct
type DS_LDAP_Class_Containment struct {
	*cim.WmiInstance

	//
	ChildClass interface{}

	//
	ParentClass interface{}
}

func NewDS_LDAP_Class_ContainmentEx1(instance *cim.WmiInstance) (newInstance *DS_LDAP_Class_Containment, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Class_Containment{
		WmiInstance: tmp,
	}
	return
}

func NewDS_LDAP_Class_ContainmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DS_LDAP_Class_Containment, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Class_Containment{
		WmiInstance: tmp,
	}
	return
}

// SetChildClass sets the value of ChildClass for the instance
func (instance *DS_LDAP_Class_Containment) SetPropertyChildClass(value interface{}) (err error) {
	return instance.SetProperty("ChildClass", (value))
}

// GetChildClass gets the value of ChildClass for the instance
func (instance *DS_LDAP_Class_Containment) GetPropertyChildClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ChildClass")
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

// SetParentClass sets the value of ParentClass for the instance
func (instance *DS_LDAP_Class_Containment) SetPropertyParentClass(value interface{}) (err error) {
	return instance.SetProperty("ParentClass", (value))
}

// GetParentClass gets the value of ParentClass for the instance
func (instance *DS_LDAP_Class_Containment) GetPropertyParentClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ParentClass")
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
