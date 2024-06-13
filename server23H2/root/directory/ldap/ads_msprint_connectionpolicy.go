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

// ads_msprint_connectionpolicy struct
type ads_msprint_connectionpolicy struct {
	*ds_top

	//
	DS_printAttributes int32

	//
	DS_printerName string

	//
	DS_serverName string

	//
	DS_uNCName string
}

func Newads_msprint_connectionpolicyEx1(instance *cim.WmiInstance) (newInstance *ads_msprint_connectionpolicy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msprint_connectionpolicy{
		ds_top: tmp,
	}
	return
}

func Newads_msprint_connectionpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msprint_connectionpolicy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msprint_connectionpolicy{
		ds_top: tmp,
	}
	return
}

// SetDS_printAttributes sets the value of DS_printAttributes for the instance
func (instance *ads_msprint_connectionpolicy) SetPropertyDS_printAttributes(value int32) (err error) {
	return instance.SetProperty("DS_printAttributes", (value))
}

// GetDS_printAttributes gets the value of DS_printAttributes for the instance
func (instance *ads_msprint_connectionpolicy) GetPropertyDS_printAttributes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printAttributes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_printerName sets the value of DS_printerName for the instance
func (instance *ads_msprint_connectionpolicy) SetPropertyDS_printerName(value string) (err error) {
	return instance.SetProperty("DS_printerName", (value))
}

// GetDS_printerName gets the value of DS_printerName for the instance
func (instance *ads_msprint_connectionpolicy) GetPropertyDS_printerName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printerName")
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

// SetDS_serverName sets the value of DS_serverName for the instance
func (instance *ads_msprint_connectionpolicy) SetPropertyDS_serverName(value string) (err error) {
	return instance.SetProperty("DS_serverName", (value))
}

// GetDS_serverName gets the value of DS_serverName for the instance
func (instance *ads_msprint_connectionpolicy) GetPropertyDS_serverName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serverName")
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
func (instance *ads_msprint_connectionpolicy) SetPropertyDS_uNCName(value string) (err error) {
	return instance.SetProperty("DS_uNCName", (value))
}

// GetDS_uNCName gets the value of DS_uNCName for the instance
func (instance *ads_msprint_connectionpolicy) GetPropertyDS_uNCName() (value string, err error) {
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
