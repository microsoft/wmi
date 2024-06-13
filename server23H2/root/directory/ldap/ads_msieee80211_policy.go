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

// ads_msieee80211_policy struct
type ads_msieee80211_policy struct {
	*ds_top

	//
	DS_msieee80211_Data Uint8Array

	//
	DS_msieee80211_DataType int32

	//
	DS_msieee80211_ID string
}

func Newads_msieee80211_policyEx1(instance *cim.WmiInstance) (newInstance *ads_msieee80211_policy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msieee80211_policy{
		ds_top: tmp,
	}
	return
}

func Newads_msieee80211_policyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msieee80211_policy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msieee80211_policy{
		ds_top: tmp,
	}
	return
}

// SetDS_msieee80211_Data sets the value of DS_msieee80211_Data for the instance
func (instance *ads_msieee80211_policy) SetPropertyDS_msieee80211_Data(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msieee80211_Data", (value))
}

// GetDS_msieee80211_Data gets the value of DS_msieee80211_Data for the instance
func (instance *ads_msieee80211_policy) GetPropertyDS_msieee80211_Data() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msieee80211_Data")
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

// SetDS_msieee80211_DataType sets the value of DS_msieee80211_DataType for the instance
func (instance *ads_msieee80211_policy) SetPropertyDS_msieee80211_DataType(value int32) (err error) {
	return instance.SetProperty("DS_msieee80211_DataType", (value))
}

// GetDS_msieee80211_DataType gets the value of DS_msieee80211_DataType for the instance
func (instance *ads_msieee80211_policy) GetPropertyDS_msieee80211_DataType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msieee80211_DataType")
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

// SetDS_msieee80211_ID sets the value of DS_msieee80211_ID for the instance
func (instance *ads_msieee80211_policy) SetPropertyDS_msieee80211_ID(value string) (err error) {
	return instance.SetProperty("DS_msieee80211_ID", (value))
}

// GetDS_msieee80211_ID gets the value of DS_msieee80211_ID for the instance
func (instance *ads_msieee80211_policy) GetPropertyDS_msieee80211_ID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msieee80211_ID")
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
