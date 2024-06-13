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

// ads_container struct
type ads_container struct {
	*ds_top

	//
	DS_defaultClassStore []string

	//
	DS_msDS_ObjectReference []string

	//
	DS_schemaVersion []int32
}

func Newads_containerEx1(instance *cim.WmiInstance) (newInstance *ads_container, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_container{
		ds_top: tmp,
	}
	return
}

func Newads_containerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_container, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_container{
		ds_top: tmp,
	}
	return
}

// SetDS_defaultClassStore sets the value of DS_defaultClassStore for the instance
func (instance *ads_container) SetPropertyDS_defaultClassStore(value []string) (err error) {
	return instance.SetProperty("DS_defaultClassStore", (value))
}

// GetDS_defaultClassStore gets the value of DS_defaultClassStore for the instance
func (instance *ads_container) GetPropertyDS_defaultClassStore() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_defaultClassStore")
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

// SetDS_msDS_ObjectReference sets the value of DS_msDS_ObjectReference for the instance
func (instance *ads_container) SetPropertyDS_msDS_ObjectReference(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ObjectReference", (value))
}

// GetDS_msDS_ObjectReference gets the value of DS_msDS_ObjectReference for the instance
func (instance *ads_container) GetPropertyDS_msDS_ObjectReference() (value []string, err error) {
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

// SetDS_schemaVersion sets the value of DS_schemaVersion for the instance
func (instance *ads_container) SetPropertyDS_schemaVersion(value []int32) (err error) {
	return instance.SetProperty("DS_schemaVersion", (value))
}

// GetDS_schemaVersion gets the value of DS_schemaVersion for the instance
func (instance *ads_container) GetPropertyDS_schemaVersion() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_schemaVersion")
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
