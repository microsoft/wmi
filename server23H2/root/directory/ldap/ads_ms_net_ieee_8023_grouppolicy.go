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

// ads_ms_net_ieee_8023_grouppolicy struct
type ads_ms_net_ieee_8023_grouppolicy struct {
	*ds_top

	//
	DS_ms_net_ieee_8023_GP_PolicyData string

	//
	DS_ms_net_ieee_8023_GP_PolicyGUID string

	//
	DS_ms_net_ieee_8023_GP_PolicyReserved Uint8Array
}

func Newads_ms_net_ieee_8023_grouppolicyEx1(instance *cim.WmiInstance) (newInstance *ads_ms_net_ieee_8023_grouppolicy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ms_net_ieee_8023_grouppolicy{
		ds_top: tmp,
	}
	return
}

func Newads_ms_net_ieee_8023_grouppolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ms_net_ieee_8023_grouppolicy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ms_net_ieee_8023_grouppolicy{
		ds_top: tmp,
	}
	return
}

// SetDS_ms_net_ieee_8023_GP_PolicyData sets the value of DS_ms_net_ieee_8023_GP_PolicyData for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) SetPropertyDS_ms_net_ieee_8023_GP_PolicyData(value string) (err error) {
	return instance.SetProperty("DS_ms_net_ieee_8023_GP_PolicyData", (value))
}

// GetDS_ms_net_ieee_8023_GP_PolicyData gets the value of DS_ms_net_ieee_8023_GP_PolicyData for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) GetPropertyDS_ms_net_ieee_8023_GP_PolicyData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ms_net_ieee_8023_GP_PolicyData")
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

// SetDS_ms_net_ieee_8023_GP_PolicyGUID sets the value of DS_ms_net_ieee_8023_GP_PolicyGUID for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) SetPropertyDS_ms_net_ieee_8023_GP_PolicyGUID(value string) (err error) {
	return instance.SetProperty("DS_ms_net_ieee_8023_GP_PolicyGUID", (value))
}

// GetDS_ms_net_ieee_8023_GP_PolicyGUID gets the value of DS_ms_net_ieee_8023_GP_PolicyGUID for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) GetPropertyDS_ms_net_ieee_8023_GP_PolicyGUID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ms_net_ieee_8023_GP_PolicyGUID")
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

// SetDS_ms_net_ieee_8023_GP_PolicyReserved sets the value of DS_ms_net_ieee_8023_GP_PolicyReserved for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) SetPropertyDS_ms_net_ieee_8023_GP_PolicyReserved(value Uint8Array) (err error) {
	return instance.SetProperty("DS_ms_net_ieee_8023_GP_PolicyReserved", (value))
}

// GetDS_ms_net_ieee_8023_GP_PolicyReserved gets the value of DS_ms_net_ieee_8023_GP_PolicyReserved for the instance
func (instance *ads_ms_net_ieee_8023_grouppolicy) GetPropertyDS_ms_net_ieee_8023_GP_PolicyReserved() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_ms_net_ieee_8023_GP_PolicyReserved")
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
