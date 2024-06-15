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

// ds_shadowaccount struct
type ds_shadowaccount struct {
	*ds_top

	//
	DS_shadowExpire int32

	//
	DS_shadowFlag int32

	//
	DS_shadowInactive int32

	//
	DS_shadowLastChange int32

	//
	DS_shadowMax int32

	//
	DS_shadowMin int32

	//
	DS_shadowWarning int32

	//
	DS_uid []string

	//
	DS_userPassword []Uint8Array
}

func Newds_shadowaccountEx1(instance *cim.WmiInstance) (newInstance *ds_shadowaccount, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_shadowaccount{
		ds_top: tmp,
	}
	return
}

func Newds_shadowaccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_shadowaccount, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_shadowaccount{
		ds_top: tmp,
	}
	return
}

// SetDS_shadowExpire sets the value of DS_shadowExpire for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowExpire(value int32) (err error) {
	return instance.SetProperty("DS_shadowExpire", (value))
}

// GetDS_shadowExpire gets the value of DS_shadowExpire for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowExpire() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowExpire")
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

// SetDS_shadowFlag sets the value of DS_shadowFlag for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowFlag(value int32) (err error) {
	return instance.SetProperty("DS_shadowFlag", (value))
}

// GetDS_shadowFlag gets the value of DS_shadowFlag for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowFlag() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowFlag")
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

// SetDS_shadowInactive sets the value of DS_shadowInactive for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowInactive(value int32) (err error) {
	return instance.SetProperty("DS_shadowInactive", (value))
}

// GetDS_shadowInactive gets the value of DS_shadowInactive for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowInactive() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowInactive")
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

// SetDS_shadowLastChange sets the value of DS_shadowLastChange for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowLastChange(value int32) (err error) {
	return instance.SetProperty("DS_shadowLastChange", (value))
}

// GetDS_shadowLastChange gets the value of DS_shadowLastChange for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowLastChange() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowLastChange")
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

// SetDS_shadowMax sets the value of DS_shadowMax for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowMax(value int32) (err error) {
	return instance.SetProperty("DS_shadowMax", (value))
}

// GetDS_shadowMax gets the value of DS_shadowMax for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowMax() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowMax")
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

// SetDS_shadowMin sets the value of DS_shadowMin for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowMin(value int32) (err error) {
	return instance.SetProperty("DS_shadowMin", (value))
}

// GetDS_shadowMin gets the value of DS_shadowMin for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowMin")
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

// SetDS_shadowWarning sets the value of DS_shadowWarning for the instance
func (instance *ds_shadowaccount) SetPropertyDS_shadowWarning(value int32) (err error) {
	return instance.SetProperty("DS_shadowWarning", (value))
}

// GetDS_shadowWarning gets the value of DS_shadowWarning for the instance
func (instance *ds_shadowaccount) GetPropertyDS_shadowWarning() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowWarning")
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ds_shadowaccount) SetPropertyDS_uid(value []string) (err error) {
	return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ds_shadowaccount) GetPropertyDS_uid() (value []string, err error) {
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ds_shadowaccount) SetPropertyDS_userPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ds_shadowaccount) GetPropertyDS_userPassword() (value []Uint8Array, err error) {
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
