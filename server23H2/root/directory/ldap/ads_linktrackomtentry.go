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

// ads_linktrackomtentry struct
type ads_linktrackomtentry struct {
	*ds_leaf

	//
	DS_birthLocation Uint8Array

	//
	DS_currentLocation Uint8Array

	//
	DS_oMTGuid Uint8Array

	//
	DS_oMTIndxGuid Uint8Array

	//
	DS_timeRefresh int64
}

func Newads_linktrackomtentryEx1(instance *cim.WmiInstance) (newInstance *ads_linktrackomtentry, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_linktrackomtentry{
		ds_leaf: tmp,
	}
	return
}

func Newads_linktrackomtentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_linktrackomtentry, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_linktrackomtentry{
		ds_leaf: tmp,
	}
	return
}

// SetDS_birthLocation sets the value of DS_birthLocation for the instance
func (instance *ads_linktrackomtentry) SetPropertyDS_birthLocation(value Uint8Array) (err error) {
	return instance.SetProperty("DS_birthLocation", (value))
}

// GetDS_birthLocation gets the value of DS_birthLocation for the instance
func (instance *ads_linktrackomtentry) GetPropertyDS_birthLocation() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_birthLocation")
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

// SetDS_currentLocation sets the value of DS_currentLocation for the instance
func (instance *ads_linktrackomtentry) SetPropertyDS_currentLocation(value Uint8Array) (err error) {
	return instance.SetProperty("DS_currentLocation", (value))
}

// GetDS_currentLocation gets the value of DS_currentLocation for the instance
func (instance *ads_linktrackomtentry) GetPropertyDS_currentLocation() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_currentLocation")
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

// SetDS_oMTGuid sets the value of DS_oMTGuid for the instance
func (instance *ads_linktrackomtentry) SetPropertyDS_oMTGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_oMTGuid", (value))
}

// GetDS_oMTGuid gets the value of DS_oMTGuid for the instance
func (instance *ads_linktrackomtentry) GetPropertyDS_oMTGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_oMTGuid")
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

// SetDS_oMTIndxGuid sets the value of DS_oMTIndxGuid for the instance
func (instance *ads_linktrackomtentry) SetPropertyDS_oMTIndxGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_oMTIndxGuid", (value))
}

// GetDS_oMTIndxGuid gets the value of DS_oMTIndxGuid for the instance
func (instance *ads_linktrackomtentry) GetPropertyDS_oMTIndxGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_oMTIndxGuid")
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

// SetDS_timeRefresh sets the value of DS_timeRefresh for the instance
func (instance *ads_linktrackomtentry) SetPropertyDS_timeRefresh(value int64) (err error) {
	return instance.SetProperty("DS_timeRefresh", (value))
}

// GetDS_timeRefresh gets the value of DS_timeRefresh for the instance
func (instance *ads_linktrackomtentry) GetPropertyDS_timeRefresh() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_timeRefresh")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}
