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

// DSClass_To_DNInstance struct
type DSClass_To_DNInstance struct {
	*cim.WmiInstance

	//
	DSClass string

	//
	RootDNForSearchAndQuery DN_Class
}

func NewDSClass_To_DNInstanceEx1(instance *cim.WmiInstance) (newInstance *DSClass_To_DNInstance, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DSClass_To_DNInstance{
		WmiInstance: tmp,
	}
	return
}

func NewDSClass_To_DNInstanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DSClass_To_DNInstance, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DSClass_To_DNInstance{
		WmiInstance: tmp,
	}
	return
}

// SetDSClass sets the value of DSClass for the instance
func (instance *DSClass_To_DNInstance) SetPropertyDSClass(value string) (err error) {
	return instance.SetProperty("DSClass", (value))
}

// GetDSClass gets the value of DSClass for the instance
func (instance *DSClass_To_DNInstance) GetPropertyDSClass() (value string, err error) {
	retValue, err := instance.GetProperty("DSClass")
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

// SetRootDNForSearchAndQuery sets the value of RootDNForSearchAndQuery for the instance
func (instance *DSClass_To_DNInstance) SetPropertyRootDNForSearchAndQuery(value DN_Class) (err error) {
	return instance.SetProperty("RootDNForSearchAndQuery", (value))
}

// GetRootDNForSearchAndQuery gets the value of RootDNForSearchAndQuery for the instance
func (instance *DSClass_To_DNInstance) GetPropertyRootDNForSearchAndQuery() (value DN_Class, err error) {
	retValue, err := instance.GetProperty("RootDNForSearchAndQuery")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DN_Class)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DN_Class is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DN_Class(valuetmp)

	return
}
