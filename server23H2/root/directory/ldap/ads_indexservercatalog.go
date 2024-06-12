// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_indexservercatalog struct
type ads_indexservercatalog struct {
	*ds_connectionpoint

	//
	DS_creator string

	//
	DS_friendlyNames []string

	//
	DS_indexedScopes []string

	//
	DS_queryPoint string

	//
	DS_uNCName string
}

func Newads_indexservercatalogEx1(instance *cim.WmiInstance) (newInstance *ads_indexservercatalog, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_indexservercatalog{
		ds_connectionpoint: tmp,
	}
	return
}

func Newads_indexservercatalogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_indexservercatalog, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_indexservercatalog{
		ds_connectionpoint: tmp,
	}
	return
}

// SetDS_creator sets the value of DS_creator for the instance
func (instance *ads_indexservercatalog) SetPropertyDS_creator(value string) (err error) {
	return instance.SetProperty("DS_creator", (value))
}

// GetDS_creator gets the value of DS_creator for the instance
func (instance *ads_indexservercatalog) GetPropertyDS_creator() (value string, err error) {
	retValue, err := instance.GetProperty("DS_creator")
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

// SetDS_friendlyNames sets the value of DS_friendlyNames for the instance
func (instance *ads_indexservercatalog) SetPropertyDS_friendlyNames(value []string) (err error) {
	return instance.SetProperty("DS_friendlyNames", (value))
}

// GetDS_friendlyNames gets the value of DS_friendlyNames for the instance
func (instance *ads_indexservercatalog) GetPropertyDS_friendlyNames() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_friendlyNames")
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

// SetDS_indexedScopes sets the value of DS_indexedScopes for the instance
func (instance *ads_indexservercatalog) SetPropertyDS_indexedScopes(value []string) (err error) {
	return instance.SetProperty("DS_indexedScopes", (value))
}

// GetDS_indexedScopes gets the value of DS_indexedScopes for the instance
func (instance *ads_indexservercatalog) GetPropertyDS_indexedScopes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_indexedScopes")
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

// SetDS_queryPoint sets the value of DS_queryPoint for the instance
func (instance *ads_indexservercatalog) SetPropertyDS_queryPoint(value string) (err error) {
	return instance.SetProperty("DS_queryPoint", (value))
}

// GetDS_queryPoint gets the value of DS_queryPoint for the instance
func (instance *ads_indexservercatalog) GetPropertyDS_queryPoint() (value string, err error) {
	retValue, err := instance.GetProperty("DS_queryPoint")
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

// SetDS_uNCName sets the value of DS_uNCName for the instance
func (instance *ads_indexservercatalog) SetPropertyDS_uNCName(value string) (err error) {
	return instance.SetProperty("DS_uNCName", (value))
}

// GetDS_uNCName gets the value of DS_uNCName for the instance
func (instance *ads_indexservercatalog) GetPropertyDS_uNCName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_uNCName")
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
