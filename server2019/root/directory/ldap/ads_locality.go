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

// ads_locality struct
type ads_locality struct {
	*ds_top

	//
	DS_l string

	//
	DS_searchGuide []Uint8Array

	//
	DS_seeAlso []string

	//
	DS_st string

	//
	DS_street string
}

func Newads_localityEx1(instance *cim.WmiInstance) (newInstance *ads_locality, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_locality{
		ds_top: tmp,
	}
	return
}

func Newads_localityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_locality, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_locality{
		ds_top: tmp,
	}
	return
}

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_locality) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_locality) GetPropertyDS_l() (value string, err error) {
	retValue, err := instance.GetProperty("DS_l")
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
func (instance *ads_locality) SetPropertyDS_searchGuide(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_searchGuide", (value))
}

// GetDS_searchGuide gets the value of DS_searchGuide for the instance
func (instance *ads_locality) GetPropertyDS_searchGuide() (value []Uint8Array, err error) {
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_locality) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_locality) GetPropertyDS_seeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_seeAlso")
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

// SetDS_st sets the value of DS_st for the instance
func (instance *ads_locality) SetPropertyDS_st(value string) (err error) {
	return instance.SetProperty("DS_st", (value))
}

// GetDS_st gets the value of DS_st for the instance
func (instance *ads_locality) GetPropertyDS_st() (value string, err error) {
	retValue, err := instance.GetProperty("DS_st")
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

// SetDS_street sets the value of DS_street for the instance
func (instance *ads_locality) SetPropertyDS_street(value string) (err error) {
	return instance.SetProperty("DS_street", (value))
}

// GetDS_street gets the value of DS_street for the instance
func (instance *ads_locality) GetPropertyDS_street() (value string, err error) {
	retValue, err := instance.GetProperty("DS_street")
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
