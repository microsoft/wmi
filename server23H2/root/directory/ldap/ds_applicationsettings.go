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

// ds_applicationsettings struct
type ds_applicationsettings struct {
	*ds_top

	//
	DS_applicationName string

	//
	DS_msDS_Settings []string

	//
	DS_notificationList string
}

func Newds_applicationsettingsEx1(instance *cim.WmiInstance) (newInstance *ds_applicationsettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_applicationsettings{
		ds_top: tmp,
	}
	return
}

func Newds_applicationsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_applicationsettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_applicationsettings{
		ds_top: tmp,
	}
	return
}

// SetDS_applicationName sets the value of DS_applicationName for the instance
func (instance *ds_applicationsettings) SetPropertyDS_applicationName(value string) (err error) {
	return instance.SetProperty("DS_applicationName", (value))
}

// GetDS_applicationName gets the value of DS_applicationName for the instance
func (instance *ds_applicationsettings) GetPropertyDS_applicationName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_applicationName")
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

// SetDS_msDS_Settings sets the value of DS_msDS_Settings for the instance
func (instance *ds_applicationsettings) SetPropertyDS_msDS_Settings(value []string) (err error) {
	return instance.SetProperty("DS_msDS_Settings", (value))
}

// GetDS_msDS_Settings gets the value of DS_msDS_Settings for the instance
func (instance *ds_applicationsettings) GetPropertyDS_msDS_Settings() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Settings")
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

// SetDS_notificationList sets the value of DS_notificationList for the instance
func (instance *ds_applicationsettings) SetPropertyDS_notificationList(value string) (err error) {
	return instance.SetProperty("DS_notificationList", (value))
}

// GetDS_notificationList gets the value of DS_notificationList for the instance
func (instance *ds_applicationsettings) GetPropertyDS_notificationList() (value string, err error) {
	retValue, err := instance.GetProperty("DS_notificationList")
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
