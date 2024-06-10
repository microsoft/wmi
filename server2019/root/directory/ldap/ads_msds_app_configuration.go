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

// ads_msds_app_configuration struct
type ads_msds_app_configuration struct {
	*ds_applicationsettings

	//
	DS_keywords []string

	//
	DS_managedBy string

	//
	DS_msDS_ByteArray []Uint8Array

	//
	DS_msDS_DateTime []string

	//
	DS_msDS_Integer []int32

	//
	DS_msDS_ObjectReference []string

	//
	DS_owner string
}

func Newads_msds_app_configurationEx1(instance *cim.WmiInstance) (newInstance *ads_msds_app_configuration, err error) {
	tmp, err := Newds_applicationsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_app_configuration{
		ds_applicationsettings: tmp,
	}
	return
}

func Newads_msds_app_configurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_app_configuration, err error) {
	tmp, err := Newds_applicationsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_app_configuration{
		ds_applicationsettings: tmp,
	}
	return
}

// SetDS_keywords sets the value of DS_keywords for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_keywords(value []string) (err error) {
	return instance.SetProperty("DS_keywords", (value))
}

// GetDS_keywords gets the value of DS_keywords for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_keywords() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_keywords")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msDS_ByteArray sets the value of DS_msDS_ByteArray for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_msDS_ByteArray(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ByteArray", (value))
}

// GetDS_msDS_ByteArray gets the value of DS_msDS_ByteArray for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_msDS_ByteArray() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ByteArray")
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

// SetDS_msDS_DateTime sets the value of DS_msDS_DateTime for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_msDS_DateTime(value []string) (err error) {
	return instance.SetProperty("DS_msDS_DateTime", (value))
}

// GetDS_msDS_DateTime gets the value of DS_msDS_DateTime for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_msDS_DateTime() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DateTime")
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

// SetDS_msDS_Integer sets the value of DS_msDS_Integer for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_msDS_Integer(value []int32) (err error) {
	return instance.SetProperty("DS_msDS_Integer", (value))
}

// GetDS_msDS_Integer gets the value of DS_msDS_Integer for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_msDS_Integer() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Integer")
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

// SetDS_msDS_ObjectReference sets the value of DS_msDS_ObjectReference for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_msDS_ObjectReference(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ObjectReference", (value))
}

// GetDS_msDS_ObjectReference gets the value of DS_msDS_ObjectReference for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_msDS_ObjectReference() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ObjectReference")
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

// SetDS_owner sets the value of DS_owner for the instance
func (instance *ads_msds_app_configuration) SetPropertyDS_owner(value string) (err error) {
	return instance.SetProperty("DS_owner", (value))
}

// GetDS_owner gets the value of DS_owner for the instance
func (instance *ads_msds_app_configuration) GetPropertyDS_owner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_owner")
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
