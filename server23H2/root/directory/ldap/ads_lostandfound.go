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

// ads_lostandfound struct
type ads_lostandfound struct {
	*ds_top

	//
	DS_moveTreeState []Uint8Array
}

func Newads_lostandfoundEx1(instance *cim.WmiInstance) (newInstance *ads_lostandfound, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_lostandfound{
		ds_top: tmp,
	}
	return
}

func Newads_lostandfoundEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_lostandfound, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_lostandfound{
		ds_top: tmp,
	}
	return
}

// SetDS_moveTreeState sets the value of DS_moveTreeState for the instance
func (instance *ads_lostandfound) SetPropertyDS_moveTreeState(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_moveTreeState", (value))
}

// GetDS_moveTreeState gets the value of DS_moveTreeState for the instance
func (instance *ads_lostandfound) GetPropertyDS_moveTreeState() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_moveTreeState")
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
