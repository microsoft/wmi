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

// ads_categoryregistration struct
type ads_categoryregistration struct {
	*ds_leaf

	//
	DS_categoryId Uint8Array

	//
	DS_localeID []int32

	//
	DS_localizedDescription []string

	//
	DS_managedBy string
}

func Newads_categoryregistrationEx1(instance *cim.WmiInstance) (newInstance *ads_categoryregistration, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_categoryregistration{
		ds_leaf: tmp,
	}
	return
}

func Newads_categoryregistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_categoryregistration, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_categoryregistration{
		ds_leaf: tmp,
	}
	return
}

// SetDS_categoryId sets the value of DS_categoryId for the instance
func (instance *ads_categoryregistration) SetPropertyDS_categoryId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_categoryId", (value))
}

// GetDS_categoryId gets the value of DS_categoryId for the instance
func (instance *ads_categoryregistration) GetPropertyDS_categoryId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_categoryId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_localeID sets the value of DS_localeID for the instance
func (instance *ads_categoryregistration) SetPropertyDS_localeID(value []int32) (err error) {
	return instance.SetProperty("DS_localeID", (value))
}

// GetDS_localeID gets the value of DS_localeID for the instance
func (instance *ads_categoryregistration) GetPropertyDS_localeID() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_localeID")
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

// SetDS_localizedDescription sets the value of DS_localizedDescription for the instance
func (instance *ads_categoryregistration) SetPropertyDS_localizedDescription(value []string) (err error) {
	return instance.SetProperty("DS_localizedDescription", (value))
}

// GetDS_localizedDescription gets the value of DS_localizedDescription for the instance
func (instance *ads_categoryregistration) GetPropertyDS_localizedDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_localizedDescription")
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
func (instance *ads_categoryregistration) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_categoryregistration) GetPropertyDS_managedBy() (value string, err error) {
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
