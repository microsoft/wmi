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

// ads_person struct
type ads_person struct {
	*ds_top

	//
	DS_attributeCertificateAttribute []Uint8Array

	//
	DS_seeAlso []string

	//
	DS_serialNumber []string

	//
	DS_sn string

	//
	DS_telephoneNumber string

	//
	DS_userPassword []Uint8Array
}

func Newads_personEx1(instance *cim.WmiInstance) (newInstance *ads_person, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_person{
		ds_top: tmp,
	}
	return
}

func Newads_personEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_person, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_person{
		ds_top: tmp,
	}
	return
}

// SetDS_attributeCertificateAttribute sets the value of DS_attributeCertificateAttribute for the instance
func (instance *ads_person) SetPropertyDS_attributeCertificateAttribute(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_attributeCertificateAttribute", (value))
}

// GetDS_attributeCertificateAttribute gets the value of DS_attributeCertificateAttribute for the instance
func (instance *ads_person) GetPropertyDS_attributeCertificateAttribute() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_attributeCertificateAttribute")
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
func (instance *ads_person) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_person) GetPropertyDS_seeAlso() (value []string, err error) {
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

// SetDS_serialNumber sets the value of DS_serialNumber for the instance
func (instance *ads_person) SetPropertyDS_serialNumber(value []string) (err error) {
	return instance.SetProperty("DS_serialNumber", (value))
}

// GetDS_serialNumber gets the value of DS_serialNumber for the instance
func (instance *ads_person) GetPropertyDS_serialNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_serialNumber")
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

// SetDS_sn sets the value of DS_sn for the instance
func (instance *ads_person) SetPropertyDS_sn(value string) (err error) {
	return instance.SetProperty("DS_sn", (value))
}

// GetDS_sn gets the value of DS_sn for the instance
func (instance *ads_person) GetPropertyDS_sn() (value string, err error) {
	retValue, err := instance.GetProperty("DS_sn")
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

// SetDS_telephoneNumber sets the value of DS_telephoneNumber for the instance
func (instance *ads_person) SetPropertyDS_telephoneNumber(value string) (err error) {
	return instance.SetProperty("DS_telephoneNumber", (value))
}

// GetDS_telephoneNumber gets the value of DS_telephoneNumber for the instance
func (instance *ads_person) GetPropertyDS_telephoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_telephoneNumber")
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ads_person) SetPropertyDS_userPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ads_person) GetPropertyDS_userPassword() (value []Uint8Array, err error) {
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
