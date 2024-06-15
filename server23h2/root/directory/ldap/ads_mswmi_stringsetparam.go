// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_mswmi_stringsetparam struct
type ads_mswmi_stringsetparam struct {
	*ads_mswmi_rangeparam

	//
	DS_msWMI_StringDefault string

	//
	DS_msWMI_StringValidValues []string
}

func Newads_mswmi_stringsetparamEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_stringsetparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_stringsetparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

func Newads_mswmi_stringsetparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_stringsetparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_stringsetparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

// SetDS_msWMI_StringDefault sets the value of DS_msWMI_StringDefault for the instance
func (instance *ads_mswmi_stringsetparam) SetPropertyDS_msWMI_StringDefault(value string) (err error) {
	return instance.SetProperty("DS_msWMI_StringDefault", (value))
}

// GetDS_msWMI_StringDefault gets the value of DS_msWMI_StringDefault for the instance
func (instance *ads_mswmi_stringsetparam) GetPropertyDS_msWMI_StringDefault() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_StringDefault")
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

// SetDS_msWMI_StringValidValues sets the value of DS_msWMI_StringValidValues for the instance
func (instance *ads_mswmi_stringsetparam) SetPropertyDS_msWMI_StringValidValues(value []string) (err error) {
	return instance.SetProperty("DS_msWMI_StringValidValues", (value))
}

// GetDS_msWMI_StringValidValues gets the value of DS_msWMI_StringValidValues for the instance
func (instance *ads_mswmi_stringsetparam) GetPropertyDS_msWMI_StringValidValues() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_StringValidValues")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
