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

// ads_sitelink struct
type ads_sitelink struct {
	*ds_top

	//
	DS_cost int32

	//
	DS_options int32

	//
	DS_replInterval int32

	//
	DS_schedule Uint8Array

	//
	DS_siteList []string
}

func Newads_sitelinkEx1(instance *cim.WmiInstance) (newInstance *ads_sitelink, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_sitelink{
		ds_top: tmp,
	}
	return
}

func Newads_sitelinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_sitelink, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_sitelink{
		ds_top: tmp,
	}
	return
}

// SetDS_cost sets the value of DS_cost for the instance
func (instance *ads_sitelink) SetPropertyDS_cost(value int32) (err error) {
	return instance.SetProperty("DS_cost", (value))
}

// GetDS_cost gets the value of DS_cost for the instance
func (instance *ads_sitelink) GetPropertyDS_cost() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_cost")
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

// SetDS_options sets the value of DS_options for the instance
func (instance *ads_sitelink) SetPropertyDS_options(value int32) (err error) {
	return instance.SetProperty("DS_options", (value))
}

// GetDS_options gets the value of DS_options for the instance
func (instance *ads_sitelink) GetPropertyDS_options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_options")
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

// SetDS_replInterval sets the value of DS_replInterval for the instance
func (instance *ads_sitelink) SetPropertyDS_replInterval(value int32) (err error) {
	return instance.SetProperty("DS_replInterval", (value))
}

// GetDS_replInterval gets the value of DS_replInterval for the instance
func (instance *ads_sitelink) GetPropertyDS_replInterval() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_replInterval")
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

// SetDS_schedule sets the value of DS_schedule for the instance
func (instance *ads_sitelink) SetPropertyDS_schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schedule", (value))
}

// GetDS_schedule gets the value of DS_schedule for the instance
func (instance *ads_sitelink) GetPropertyDS_schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schedule")
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

// SetDS_siteList sets the value of DS_siteList for the instance
func (instance *ads_sitelink) SetPropertyDS_siteList(value []string) (err error) {
	return instance.SetProperty("DS_siteList", (value))
}

// GetDS_siteList gets the value of DS_siteList for the instance
func (instance *ads_sitelink) GetPropertyDS_siteList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_siteList")
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
