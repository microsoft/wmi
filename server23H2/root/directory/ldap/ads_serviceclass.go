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

// ads_serviceclass struct
type ads_serviceclass struct {
	*ds_leaf

	//
	DS_serviceClassID Uint8Array

	//
	DS_serviceClassInfo []Uint8Array
}

func Newads_serviceclassEx1(instance *cim.WmiInstance) (newInstance *ads_serviceclass, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_serviceclass{
		ds_leaf: tmp,
	}
	return
}

func Newads_serviceclassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_serviceclass, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_serviceclass{
		ds_leaf: tmp,
	}
	return
}

// SetDS_serviceClassID sets the value of DS_serviceClassID for the instance
func (instance *ads_serviceclass) SetPropertyDS_serviceClassID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_serviceClassID", (value))
}

// GetDS_serviceClassID gets the value of DS_serviceClassID for the instance
func (instance *ads_serviceclass) GetPropertyDS_serviceClassID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_serviceClassID")
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

// SetDS_serviceClassInfo sets the value of DS_serviceClassInfo for the instance
func (instance *ads_serviceclass) SetPropertyDS_serviceClassInfo(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_serviceClassInfo", (value))
}

// GetDS_serviceClassInfo gets the value of DS_serviceClassInfo for the instance
func (instance *ads_serviceclass) GetPropertyDS_serviceClassInfo() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_serviceClassInfo")
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
