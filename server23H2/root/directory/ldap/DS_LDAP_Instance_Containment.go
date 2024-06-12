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

// DS_LDAP_Instance_Containment struct
type DS_LDAP_Instance_Containment struct {
	*cim.WmiInstance

	//
	ChildInstance DS_LDAP_Root_Class

	//
	ParentInstance DS_LDAP_Root_Class
}

func NewDS_LDAP_Instance_ContainmentEx1(instance *cim.WmiInstance) (newInstance *DS_LDAP_Instance_Containment, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Instance_Containment{
		WmiInstance: tmp,
	}
	return
}

func NewDS_LDAP_Instance_ContainmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DS_LDAP_Instance_Containment, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Instance_Containment{
		WmiInstance: tmp,
	}
	return
}

// SetChildInstance sets the value of ChildInstance for the instance
func (instance *DS_LDAP_Instance_Containment) SetPropertyChildInstance(value DS_LDAP_Root_Class) (err error) {
	return instance.SetProperty("ChildInstance", (value))
}

// GetChildInstance gets the value of ChildInstance for the instance
func (instance *DS_LDAP_Instance_Containment) GetPropertyChildInstance() (value DS_LDAP_Root_Class, err error) {
	retValue, err := instance.GetProperty("ChildInstance")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DS_LDAP_Root_Class)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DS_LDAP_Root_Class is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DS_LDAP_Root_Class(valuetmp)

	return
}

// SetParentInstance sets the value of ParentInstance for the instance
func (instance *DS_LDAP_Instance_Containment) SetPropertyParentInstance(value DS_LDAP_Root_Class) (err error) {
	return instance.SetProperty("ParentInstance", (value))
}

// GetParentInstance gets the value of ParentInstance for the instance
func (instance *DS_LDAP_Instance_Containment) GetPropertyParentInstance() (value DS_LDAP_Root_Class, err error) {
	retValue, err := instance.GetProperty("ParentInstance")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DS_LDAP_Root_Class)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DS_LDAP_Root_Class is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DS_LDAP_Root_Class(valuetmp)

	return
}
