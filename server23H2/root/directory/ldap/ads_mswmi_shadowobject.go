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

// ads_mswmi_shadowobject struct
type ads_mswmi_shadowobject struct {
	*ds_top

	//
	DS_msWMI_TargetObject []Uint8Array
}

func Newads_mswmi_shadowobjectEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_shadowobject, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_shadowobject{
		ds_top: tmp,
	}
	return
}

func Newads_mswmi_shadowobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_shadowobject, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_shadowobject{
		ds_top: tmp,
	}
	return
}

// SetDS_msWMI_TargetObject sets the value of DS_msWMI_TargetObject for the instance
func (instance *ads_mswmi_shadowobject) SetPropertyDS_msWMI_TargetObject(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msWMI_TargetObject", (value))
}

// GetDS_msWMI_TargetObject gets the value of DS_msWMI_TargetObject for the instance
func (instance *ads_mswmi_shadowobject) GetPropertyDS_msWMI_TargetObject() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetObject")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}
