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

// ads_mswmi_rangeparam struct
type ads_mswmi_rangeparam struct {
	*ds_top

	//
	DS_msWMI_PropertyName string

	//
	DS_msWMI_TargetClass string

	//
	DS_msWMI_TargetType string
}

func Newads_mswmi_rangeparamEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_rangeparam, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_rangeparam{
		ds_top: tmp,
	}
	return
}

func Newads_mswmi_rangeparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_rangeparam, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_rangeparam{
		ds_top: tmp,
	}
	return
}

// SetDS_msWMI_PropertyName sets the value of DS_msWMI_PropertyName for the instance
func (instance *ads_mswmi_rangeparam) SetPropertyDS_msWMI_PropertyName(value string) (err error) {
	return instance.SetProperty("DS_msWMI_PropertyName", (value))
}

// GetDS_msWMI_PropertyName gets the value of DS_msWMI_PropertyName for the instance
func (instance *ads_mswmi_rangeparam) GetPropertyDS_msWMI_PropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_PropertyName")
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

// SetDS_msWMI_TargetClass sets the value of DS_msWMI_TargetClass for the instance
func (instance *ads_mswmi_rangeparam) SetPropertyDS_msWMI_TargetClass(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetClass", (value))
}

// GetDS_msWMI_TargetClass gets the value of DS_msWMI_TargetClass for the instance
func (instance *ads_mswmi_rangeparam) GetPropertyDS_msWMI_TargetClass() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetClass")
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

// SetDS_msWMI_TargetType sets the value of DS_msWMI_TargetType for the instance
func (instance *ads_mswmi_rangeparam) SetPropertyDS_msWMI_TargetType(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetType", (value))
}

// GetDS_msWMI_TargetType gets the value of DS_msWMI_TargetType for the instance
func (instance *ads_mswmi_rangeparam) GetPropertyDS_msWMI_TargetType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetType")
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
