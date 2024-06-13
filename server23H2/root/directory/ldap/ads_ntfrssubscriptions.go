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

// ads_ntfrssubscriptions struct
type ads_ntfrssubscriptions struct {
	*ds_top

	//
	DS_fRSExtensions Uint8Array

	//
	DS_fRSVersion string

	//
	DS_fRSWorkingPath string
}

func Newads_ntfrssubscriptionsEx1(instance *cim.WmiInstance) (newInstance *ads_ntfrssubscriptions, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrssubscriptions{
		ds_top: tmp,
	}
	return
}

func Newads_ntfrssubscriptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntfrssubscriptions, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrssubscriptions{
		ds_top: tmp,
	}
	return
}

// SetDS_fRSExtensions sets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssubscriptions) SetPropertyDS_fRSExtensions(value Uint8Array) (err error) {
	return instance.SetProperty("DS_fRSExtensions", (value))
}

// GetDS_fRSExtensions gets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssubscriptions) GetPropertyDS_fRSExtensions() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_fRSExtensions")
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

// SetDS_fRSVersion sets the value of DS_fRSVersion for the instance
func (instance *ads_ntfrssubscriptions) SetPropertyDS_fRSVersion(value string) (err error) {
	return instance.SetProperty("DS_fRSVersion", (value))
}

// GetDS_fRSVersion gets the value of DS_fRSVersion for the instance
func (instance *ads_ntfrssubscriptions) GetPropertyDS_fRSVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSVersion")
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

// SetDS_fRSWorkingPath sets the value of DS_fRSWorkingPath for the instance
func (instance *ads_ntfrssubscriptions) SetPropertyDS_fRSWorkingPath(value string) (err error) {
	return instance.SetProperty("DS_fRSWorkingPath", (value))
}

// GetDS_fRSWorkingPath gets the value of DS_fRSWorkingPath for the instance
func (instance *ads_ntfrssubscriptions) GetPropertyDS_fRSWorkingPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSWorkingPath")
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
