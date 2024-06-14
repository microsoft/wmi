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

// ads_country struct
type ads_country struct {
	*ds_top

	//
	DS_c string

	//
	DS_co string

	//
	DS_searchGuide []Uint8Array
}

func Newads_countryEx1(instance *cim.WmiInstance) (newInstance *ads_country, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_country{
		ds_top: tmp,
	}
	return
}

func Newads_countryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_country, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_country{
		ds_top: tmp,
	}
	return
}

// SetDS_c sets the value of DS_c for the instance
func (instance *ads_country) SetPropertyDS_c(value string) (err error) {
	return instance.SetProperty("DS_c", (value))
}

// GetDS_c gets the value of DS_c for the instance
func (instance *ads_country) GetPropertyDS_c() (value string, err error) {
	retValue, err := instance.GetProperty("DS_c")
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

// SetDS_co sets the value of DS_co for the instance
func (instance *ads_country) SetPropertyDS_co(value string) (err error) {
	return instance.SetProperty("DS_co", (value))
}

// GetDS_co gets the value of DS_co for the instance
func (instance *ads_country) GetPropertyDS_co() (value string, err error) {
	retValue, err := instance.GetProperty("DS_co")
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

// SetDS_searchGuide sets the value of DS_searchGuide for the instance
func (instance *ads_country) SetPropertyDS_searchGuide(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_searchGuide", (value))
}

// GetDS_searchGuide gets the value of DS_searchGuide for the instance
func (instance *ads_country) GetPropertyDS_searchGuide() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_searchGuide")
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
