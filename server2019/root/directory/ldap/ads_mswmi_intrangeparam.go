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

// ads_mswmi_intrangeparam struct
type ads_mswmi_intrangeparam struct {
	*ads_mswmi_rangeparam

	//
	DS_msWMI_IntDefault int32

	//
	DS_msWMI_IntMax int32

	//
	DS_msWMI_IntMin int32
}

func Newads_mswmi_intrangeparamEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_intrangeparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_intrangeparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

func Newads_mswmi_intrangeparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_intrangeparam, err error) {
	tmp, err := Newads_mswmi_rangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_intrangeparam{
		ads_mswmi_rangeparam: tmp,
	}
	return
}

// SetDS_msWMI_IntDefault sets the value of DS_msWMI_IntDefault for the instance
func (instance *ads_mswmi_intrangeparam) SetPropertyDS_msWMI_IntDefault(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_IntDefault", (value))
}

// GetDS_msWMI_IntDefault gets the value of DS_msWMI_IntDefault for the instance
func (instance *ads_mswmi_intrangeparam) GetPropertyDS_msWMI_IntDefault() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_IntDefault")
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

// SetDS_msWMI_IntMax sets the value of DS_msWMI_IntMax for the instance
func (instance *ads_mswmi_intrangeparam) SetPropertyDS_msWMI_IntMax(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_IntMax", (value))
}

// GetDS_msWMI_IntMax gets the value of DS_msWMI_IntMax for the instance
func (instance *ads_mswmi_intrangeparam) GetPropertyDS_msWMI_IntMax() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_IntMax")
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

// SetDS_msWMI_IntMin sets the value of DS_msWMI_IntMin for the instance
func (instance *ads_mswmi_intrangeparam) SetPropertyDS_msWMI_IntMin(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_IntMin", (value))
}

// GetDS_msWMI_IntMin gets the value of DS_msWMI_IntMin for the instance
func (instance *ads_mswmi_intrangeparam) GetPropertyDS_msWMI_IntMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_IntMin")
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
