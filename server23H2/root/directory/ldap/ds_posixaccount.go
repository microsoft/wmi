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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ds_posixaccount struct
type ds_posixaccount struct {
	*ds_top

	//
	DS_gecos string

	//
	DS_gidNumber int32

	//
	DS_homeDirectory string

	//
	DS_loginShell string

	//
	DS_uid []string

	//
	DS_uidNumber int32

	//
	DS_unixHomeDirectory string

	//
	DS_unixUserPassword []Uint8Array

	//
	DS_userPassword []Uint8Array
}

func Newds_posixaccountEx1(instance *cim.WmiInstance) (newInstance *ds_posixaccount, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_posixaccount{
		ds_top: tmp,
	}
	return
}

func Newds_posixaccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_posixaccount, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_posixaccount{
		ds_top: tmp,
	}
	return
}

// SetDS_gecos sets the value of DS_gecos for the instance
func (instance *ds_posixaccount) SetPropertyDS_gecos(value string) (err error) {
	return instance.SetProperty("DS_gecos", (value))
}

// GetDS_gecos gets the value of DS_gecos for the instance
func (instance *ds_posixaccount) GetPropertyDS_gecos() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gecos")
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

// SetDS_gidNumber sets the value of DS_gidNumber for the instance
func (instance *ds_posixaccount) SetPropertyDS_gidNumber(value int32) (err error) {
	return instance.SetProperty("DS_gidNumber", (value))
}

// GetDS_gidNumber gets the value of DS_gidNumber for the instance
func (instance *ds_posixaccount) GetPropertyDS_gidNumber() (value int32, err error) {
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

// SetDS_homeDirectory sets the value of DS_homeDirectory for the instance
func (instance *ds_posixaccount) SetPropertyDS_homeDirectory(value string) (err error) {
	return instance.SetProperty("DS_homeDirectory", (value))
}

// GetDS_homeDirectory gets the value of DS_homeDirectory for the instance
func (instance *ds_posixaccount) GetPropertyDS_homeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_homeDirectory")
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

// SetDS_loginShell sets the value of DS_loginShell for the instance
func (instance *ds_posixaccount) SetPropertyDS_loginShell(value string) (err error) {
	return instance.SetProperty("DS_loginShell", (value))
}

// GetDS_loginShell gets the value of DS_loginShell for the instance
func (instance *ds_posixaccount) GetPropertyDS_loginShell() (value string, err error) {
	retValue, err := instance.GetProperty("DS_loginShell")
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ds_posixaccount) SetPropertyDS_uid(value []string) (err error) {
	return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ds_posixaccount) GetPropertyDS_uid() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_uid")
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

// SetDS_uidNumber sets the value of DS_uidNumber for the instance
func (instance *ds_posixaccount) SetPropertyDS_uidNumber(value int32) (err error) {
	return instance.SetProperty("DS_uidNumber", (value))
}

// GetDS_uidNumber gets the value of DS_uidNumber for the instance
func (instance *ds_posixaccount) GetPropertyDS_uidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_uidNumber")
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

// SetDS_unixHomeDirectory sets the value of DS_unixHomeDirectory for the instance
func (instance *ds_posixaccount) SetPropertyDS_unixHomeDirectory(value string) (err error) {
	return instance.SetProperty("DS_unixHomeDirectory", (value))
}

// GetDS_unixHomeDirectory gets the value of DS_unixHomeDirectory for the instance
func (instance *ds_posixaccount) GetPropertyDS_unixHomeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_unixHomeDirectory")
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

// SetDS_unixUserPassword sets the value of DS_unixUserPassword for the instance
func (instance *ds_posixaccount) SetPropertyDS_unixUserPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_unixUserPassword", (value))
}

// GetDS_unixUserPassword gets the value of DS_unixUserPassword for the instance
func (instance *ds_posixaccount) GetPropertyDS_unixUserPassword() (value []Uint8Array, err error) {
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
func (instance *ds_posixaccount) SetPropertyDS_userPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ds_posixaccount) GetPropertyDS_userPassword() (value []Uint8Array, err error) {
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
