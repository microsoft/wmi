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

// ads_msds_groupmanagedserviceaccount struct
type ads_msds_groupmanagedserviceaccount struct {
	*ads_computer

	//
	DS_msDS_GroupMSAMembership Uint8Array

	//
	DS_msDS_ManagedPassword Uint8Array

	//
	DS_msDS_ManagedPasswordId Uint8Array

	//
	DS_msDS_ManagedPasswordInterval int32

	//
	DS_msDS_ManagedPasswordPreviousId Uint8Array
}

func Newads_msds_groupmanagedserviceaccountEx1(instance *cim.WmiInstance) (newInstance *ads_msds_groupmanagedserviceaccount, err error) {
	tmp, err := Newads_computerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_groupmanagedserviceaccount{
		ads_computer: tmp,
	}
	return
}

func Newads_msds_groupmanagedserviceaccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_groupmanagedserviceaccount, err error) {
	tmp, err := Newads_computerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_groupmanagedserviceaccount{
		ads_computer: tmp,
	}
	return
}

// SetDS_msDS_GroupMSAMembership sets the value of DS_msDS_GroupMSAMembership for the instance
func (instance *ads_msds_groupmanagedserviceaccount) SetPropertyDS_msDS_GroupMSAMembership(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_GroupMSAMembership", (value))
}

// GetDS_msDS_GroupMSAMembership gets the value of DS_msDS_GroupMSAMembership for the instance
func (instance *ads_msds_groupmanagedserviceaccount) GetPropertyDS_msDS_GroupMSAMembership() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GroupMSAMembership")
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

// SetDS_msDS_ManagedPassword sets the value of DS_msDS_ManagedPassword for the instance
func (instance *ads_msds_groupmanagedserviceaccount) SetPropertyDS_msDS_ManagedPassword(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ManagedPassword", (value))
}

// GetDS_msDS_ManagedPassword gets the value of DS_msDS_ManagedPassword for the instance
func (instance *ads_msds_groupmanagedserviceaccount) GetPropertyDS_msDS_ManagedPassword() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ManagedPassword")
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

// SetDS_msDS_ManagedPasswordId sets the value of DS_msDS_ManagedPasswordId for the instance
func (instance *ads_msds_groupmanagedserviceaccount) SetPropertyDS_msDS_ManagedPasswordId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ManagedPasswordId", (value))
}

// GetDS_msDS_ManagedPasswordId gets the value of DS_msDS_ManagedPasswordId for the instance
func (instance *ads_msds_groupmanagedserviceaccount) GetPropertyDS_msDS_ManagedPasswordId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ManagedPasswordId")
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

// SetDS_msDS_ManagedPasswordInterval sets the value of DS_msDS_ManagedPasswordInterval for the instance
func (instance *ads_msds_groupmanagedserviceaccount) SetPropertyDS_msDS_ManagedPasswordInterval(value int32) (err error) {
	return instance.SetProperty("DS_msDS_ManagedPasswordInterval", (value))
}

// GetDS_msDS_ManagedPasswordInterval gets the value of DS_msDS_ManagedPasswordInterval for the instance
func (instance *ads_msds_groupmanagedserviceaccount) GetPropertyDS_msDS_ManagedPasswordInterval() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ManagedPasswordInterval")
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

// SetDS_msDS_ManagedPasswordPreviousId sets the value of DS_msDS_ManagedPasswordPreviousId for the instance
func (instance *ads_msds_groupmanagedserviceaccount) SetPropertyDS_msDS_ManagedPasswordPreviousId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ManagedPasswordPreviousId", (value))
}

// GetDS_msDS_ManagedPasswordPreviousId gets the value of DS_msDS_ManagedPasswordPreviousId for the instance
func (instance *ads_msds_groupmanagedserviceaccount) GetPropertyDS_msDS_ManagedPasswordPreviousId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ManagedPasswordPreviousId")
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
