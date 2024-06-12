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

// ads_mswmi_realrangeparam struct
type ads_mswmi_realrangeparam struct {
	*ads_mswmi_rangeparam

	//
	DS_msWMI_Int8Default int64

	//
	DS_msWMI_Int8Max int64

	//
	DS_msWMI_Int8Min int64
}

func Newads_mswmi_realrangeparamEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_realrangeparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_realrangeparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

func Newads_mswmi_realrangeparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_realrangeparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_realrangeparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

// SetDS_msWMI_Int8Default sets the value of DS_msWMI_Int8Default for the instance
func (instance *ads_mswmi_realrangeparam) SetPropertyDS_msWMI_Int8Default(value int64) (err error) {
	return instance.SetProperty("DS_msWMI_Int8Default", (value))
}

// GetDS_msWMI_Int8Default gets the value of DS_msWMI_Int8Default for the instance
func (instance *ads_mswmi_realrangeparam) GetPropertyDS_msWMI_Int8Default() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Int8Default")
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

// SetDS_msWMI_Int8Max sets the value of DS_msWMI_Int8Max for the instance
func (instance *ads_mswmi_realrangeparam) SetPropertyDS_msWMI_Int8Max(value int64) (err error) {
	return instance.SetProperty("DS_msWMI_Int8Max", (value))
}

// GetDS_msWMI_Int8Max gets the value of DS_msWMI_Int8Max for the instance
func (instance *ads_mswmi_realrangeparam) GetPropertyDS_msWMI_Int8Max() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Int8Max")
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

// SetDS_msWMI_Int8Min sets the value of DS_msWMI_Int8Min for the instance
func (instance *ads_mswmi_realrangeparam) SetPropertyDS_msWMI_Int8Min(value int64) (err error) {
	return instance.SetProperty("DS_msWMI_Int8Min", (value))
}

// GetDS_msWMI_Int8Min gets the value of DS_msWMI_Int8Min for the instance
func (instance *ads_mswmi_realrangeparam) GetPropertyDS_msWMI_Int8Min() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Int8Min")
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
