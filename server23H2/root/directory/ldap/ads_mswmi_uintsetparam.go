// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_mswmi_uintsetparam struct
type ads_mswmi_uintsetparam struct {
	*ads_mswmi_rangeparam

	//
	DS_msWMI_IntDefault int32

	//
	DS_msWMI_IntValidValues []int32
}

func Newads_mswmi_uintsetparamEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_uintsetparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_uintsetparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

func Newads_mswmi_uintsetparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_uintsetparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_uintsetparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

// SetDS_msWMI_IntDefault sets the value of DS_msWMI_IntDefault for the instance
func (instance *ads_mswmi_uintsetparam) SetPropertyDS_msWMI_IntDefault(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_IntDefault", (value))
}

// GetDS_msWMI_IntDefault gets the value of DS_msWMI_IntDefault for the instance
func (instance *ads_mswmi_uintsetparam) GetPropertyDS_msWMI_IntDefault() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_IntDefault")
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

// SetDS_msWMI_IntValidValues sets the value of DS_msWMI_IntValidValues for the instance
func (instance *ads_mswmi_uintsetparam) SetPropertyDS_msWMI_IntValidValues(value []int32) (err error) {
	return instance.SetProperty("DS_msWMI_IntValidValues", (value))
}

// GetDS_msWMI_IntValidValues gets the value of DS_msWMI_IntValidValues for the instance
func (instance *ads_mswmi_uintsetparam) GetPropertyDS_msWMI_IntValidValues() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_IntValidValues")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}
