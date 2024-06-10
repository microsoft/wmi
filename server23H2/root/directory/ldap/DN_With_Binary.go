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

// DN_With_Binary struct
type DN_With_Binary struct {
	*cim.WmiInstance

	//
	dnString string

	//
	value []uint8
}

func NewDN_With_BinaryEx1(instance *cim.WmiInstance) (newInstance *DN_With_Binary, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DN_With_Binary{
		WmiInstance: tmp,
	}
	return
}

func NewDN_With_BinaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DN_With_Binary, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DN_With_Binary{
		WmiInstance: tmp,
	}
	return
}

// SetdnString sets the value of dnString for the instance
func (instance *DN_With_Binary) SetPropertydnString(value string) (err error) {
	return instance.SetProperty("dnString", (value))
}

// GetdnString gets the value of dnString for the instance
func (instance *DN_With_Binary) GetPropertydnString() (value string, err error) {
	retValue, err := instance.GetProperty("dnString")
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

// Setvalue sets the value of value for the instance
func (instance *DN_With_Binary) SetPropertyvalue(value []uint8) (err error) {
	return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *DN_With_Binary) GetPropertyvalue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
