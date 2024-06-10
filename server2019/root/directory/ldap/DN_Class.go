// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DN_Class struct
type DN_Class struct {
	*cim.WmiInstance

	//
	DN string
}

func NewDN_ClassEx1(instance *cim.WmiInstance) (newInstance *DN_Class, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DN_Class{
		WmiInstance: tmp,
	}
	return
}

func NewDN_ClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DN_Class, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DN_Class{
		WmiInstance: tmp,
	}
	return
}

// SetDN sets the value of DN for the instance
func (instance *DN_Class) SetPropertyDN(value string) (err error) {
	return instance.SetProperty("DN", (value))
}

// GetDN gets the value of DN for the instance
func (instance *DN_Class) GetPropertyDN() (value string, err error) {
	retValue, err := instance.GetProperty("DN")
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
