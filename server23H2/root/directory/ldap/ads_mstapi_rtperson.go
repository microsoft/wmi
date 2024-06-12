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

// ads_mstapi_rtperson struct
type ads_mstapi_rtperson struct {
	*ds_top

	//
	DS_msTAPI_IpAddress []string

	//
	DS_msTAPI_uid string
}

func Newads_mstapi_rtpersonEx1(instance *cim.WmiInstance) (newInstance *ads_mstapi_rtperson, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mstapi_rtperson{
		ds_top: tmp,
	}
	return
}

func Newads_mstapi_rtpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mstapi_rtperson, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mstapi_rtperson{
		ds_top: tmp,
	}
	return
}

// SetDS_msTAPI_IpAddress sets the value of DS_msTAPI_IpAddress for the instance
func (instance *ads_mstapi_rtperson) SetPropertyDS_msTAPI_IpAddress(value []string) (err error) {
	return instance.SetProperty("DS_msTAPI_IpAddress", (value))
}

// GetDS_msTAPI_IpAddress gets the value of DS_msTAPI_IpAddress for the instance
func (instance *ads_mstapi_rtperson) GetPropertyDS_msTAPI_IpAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTAPI_IpAddress")
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

// SetDS_msTAPI_uid sets the value of DS_msTAPI_uid for the instance
func (instance *ads_mstapi_rtperson) SetPropertyDS_msTAPI_uid(value string) (err error) {
	return instance.SetProperty("DS_msTAPI_uid", (value))
}

// GetDS_msTAPI_uid gets the value of DS_msTAPI_uid for the instance
func (instance *ads_mstapi_rtperson) GetPropertyDS_msTAPI_uid() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTAPI_uid")
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
