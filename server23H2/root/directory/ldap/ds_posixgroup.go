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

// ds_posixgroup struct
type ds_posixgroup struct {
	*ds_top

	//
	DS_gidNumber int32

	//
	DS_memberUid []string

	//
	DS_unixUserPassword []Uint8Array

	//
	DS_userPassword []Uint8Array
}

func Newds_posixgroupEx1(instance *cim.WmiInstance) (newInstance *ds_posixgroup, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_posixgroup{
		ds_top: tmp,
	}
	return
}

func Newds_posixgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_posixgroup, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_posixgroup{
		ds_top: tmp,
	}
	return
}

// SetDS_gidNumber sets the value of DS_gidNumber for the instance
func (instance *ds_posixgroup) SetPropertyDS_gidNumber(value int32) (err error) {
	return instance.SetProperty("DS_gidNumber", (value))
}

// GetDS_gidNumber gets the value of DS_gidNumber for the instance
func (instance *ds_posixgroup) GetPropertyDS_gidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_gidNumber")
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

// SetDS_memberUid sets the value of DS_memberUid for the instance
func (instance *ds_posixgroup) SetPropertyDS_memberUid(value []string) (err error) {
	return instance.SetProperty("DS_memberUid", (value))
}

// GetDS_memberUid gets the value of DS_memberUid for the instance
func (instance *ds_posixgroup) GetPropertyDS_memberUid() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_memberUid")
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

// SetDS_unixUserPassword sets the value of DS_unixUserPassword for the instance
func (instance *ds_posixgroup) SetPropertyDS_unixUserPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_unixUserPassword", (value))
}

// GetDS_unixUserPassword gets the value of DS_unixUserPassword for the instance
func (instance *ds_posixgroup) GetPropertyDS_unixUserPassword() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_unixUserPassword")
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ds_posixgroup) SetPropertyDS_userPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ds_posixgroup) GetPropertyDS_userPassword() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userPassword")
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
