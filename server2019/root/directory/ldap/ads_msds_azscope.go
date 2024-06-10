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

// ads_msds_azscope struct
type ads_msds_azscope struct {
	*ds_top

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzObjectGuid Uint8Array

	//
	DS_msDS_AzScopeName string
}

func Newads_msds_azscopeEx1(instance *cim.WmiInstance) (newInstance *ads_msds_azscope, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azscope{
		ds_top: tmp,
	}
	return
}

func Newads_msds_azscopeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_azscope, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azscope{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azscope) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azscope) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationData")
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

// SetDS_msDS_AzGenericData sets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azscope) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azscope) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzGenericData")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azscope) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azscope) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzObjectGuid")
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

// SetDS_msDS_AzScopeName sets the value of DS_msDS_AzScopeName for the instance
func (instance *ads_msds_azscope) SetPropertyDS_msDS_AzScopeName(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzScopeName", (value))
}

// GetDS_msDS_AzScopeName gets the value of DS_msDS_AzScopeName for the instance
func (instance *ads_msds_azscope) GetPropertyDS_msDS_AzScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzScopeName")
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
