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

// ads_msds_azoperation struct
type ads_msds_azoperation struct {
	*ds_top

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzObjectGuid Uint8Array

	//
	DS_msDS_AzOperationID int32
}

func Newads_msds_azoperationEx1(instance *cim.WmiInstance) (newInstance *ads_msds_azoperation, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azoperation{
		ds_top: tmp,
	}
	return
}

func Newads_msds_azoperationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_azoperation, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azoperation{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azoperation) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azoperation) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
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
func (instance *ads_msds_azoperation) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azoperation) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
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
func (instance *ads_msds_azoperation) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azoperation) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
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

// SetDS_msDS_AzOperationID sets the value of DS_msDS_AzOperationID for the instance
func (instance *ads_msds_azoperation) SetPropertyDS_msDS_AzOperationID(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzOperationID", (value))
}

// GetDS_msDS_AzOperationID gets the value of DS_msDS_AzOperationID for the instance
func (instance *ads_msds_azoperation) GetPropertyDS_msDS_AzOperationID() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzOperationID")
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
