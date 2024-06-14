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

// ads_typelibrary struct
type ads_typelibrary struct {
	*ds_top

	//
	DS_cOMClassID []string

	//
	DS_cOMInterfaceID []string

	//
	DS_cOMUniqueLIBID string
}

func Newads_typelibraryEx1(instance *cim.WmiInstance) (newInstance *ads_typelibrary, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_typelibrary{
		ds_top: tmp,
	}
	return
}

func Newads_typelibraryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_typelibrary, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_typelibrary{
		ds_top: tmp,
	}
	return
}

// SetDS_cOMClassID sets the value of DS_cOMClassID for the instance
func (instance *ads_typelibrary) SetPropertyDS_cOMClassID(value []string) (err error) {
	return instance.SetProperty("DS_cOMClassID", (value))
}

// GetDS_cOMClassID gets the value of DS_cOMClassID for the instance
func (instance *ads_typelibrary) GetPropertyDS_cOMClassID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMClassID")
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

// SetDS_cOMInterfaceID sets the value of DS_cOMInterfaceID for the instance
func (instance *ads_typelibrary) SetPropertyDS_cOMInterfaceID(value []string) (err error) {
	return instance.SetProperty("DS_cOMInterfaceID", (value))
}

// GetDS_cOMInterfaceID gets the value of DS_cOMInterfaceID for the instance
func (instance *ads_typelibrary) GetPropertyDS_cOMInterfaceID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMInterfaceID")
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

// SetDS_cOMUniqueLIBID sets the value of DS_cOMUniqueLIBID for the instance
func (instance *ads_typelibrary) SetPropertyDS_cOMUniqueLIBID(value string) (err error) {
	return instance.SetProperty("DS_cOMUniqueLIBID", (value))
}

// GetDS_cOMUniqueLIBID gets the value of DS_cOMUniqueLIBID for the instance
func (instance *ads_typelibrary) GetPropertyDS_cOMUniqueLIBID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cOMUniqueLIBID")
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
