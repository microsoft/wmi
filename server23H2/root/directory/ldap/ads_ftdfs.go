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

// ads_ftdfs struct
type ads_ftdfs struct {
	*ds_top

	//
	DS_keywords []string

	//
	DS_managedBy string

	//
	DS_pKT Uint8Array

	//
	DS_pKTGuid Uint8Array

	//
	DS_remoteServerName []string

	//
	DS_uNCName string
}

func Newads_ftdfsEx1(instance *cim.WmiInstance) (newInstance *ads_ftdfs, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ftdfs{
		ds_top: tmp,
	}
	return
}

func Newads_ftdfsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ftdfs, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ftdfs{
		ds_top: tmp,
	}
	return
}

// SetDS_keywords sets the value of DS_keywords for the instance
func (instance *ads_ftdfs) SetPropertyDS_keywords(value []string) (err error) {
	return instance.SetProperty("DS_keywords", (value))
}

// GetDS_keywords gets the value of DS_keywords for the instance
func (instance *ads_ftdfs) GetPropertyDS_keywords() (value []string, err error) {
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
func (instance *ads_ftdfs) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_ftdfs) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_pKT sets the value of DS_pKT for the instance
func (instance *ads_ftdfs) SetPropertyDS_pKT(value Uint8Array) (err error) {
	return instance.SetProperty("DS_pKT", (value))
}

// GetDS_pKT gets the value of DS_pKT for the instance
func (instance *ads_ftdfs) GetPropertyDS_pKT() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_pKT")
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

// SetDS_pKTGuid sets the value of DS_pKTGuid for the instance
func (instance *ads_ftdfs) SetPropertyDS_pKTGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_pKTGuid", (value))
}

// GetDS_pKTGuid gets the value of DS_pKTGuid for the instance
func (instance *ads_ftdfs) GetPropertyDS_pKTGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_pKTGuid")
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

// SetDS_remoteServerName sets the value of DS_remoteServerName for the instance
func (instance *ads_ftdfs) SetPropertyDS_remoteServerName(value []string) (err error) {
	return instance.SetProperty("DS_remoteServerName", (value))
}

// GetDS_remoteServerName gets the value of DS_remoteServerName for the instance
func (instance *ads_ftdfs) GetPropertyDS_remoteServerName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_remoteServerName")
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

// SetDS_uNCName sets the value of DS_uNCName for the instance
func (instance *ads_ftdfs) SetPropertyDS_uNCName(value string) (err error) {
	return instance.SetProperty("DS_uNCName", (value))
}

// GetDS_uNCName gets the value of DS_uNCName for the instance
func (instance *ads_ftdfs) GetPropertyDS_uNCName() (value string, err error) {
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
