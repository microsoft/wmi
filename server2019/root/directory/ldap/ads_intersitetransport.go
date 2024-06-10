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

// ads_intersitetransport struct
type ads_intersitetransport struct {
	*ds_top

	//
	DS_options int32

	//
	DS_replInterval int32

	//
	DS_transportAddressAttribute string

	//
	DS_transportDLLName string
}

func Newads_intersitetransportEx1(instance *cim.WmiInstance) (newInstance *ads_intersitetransport, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_intersitetransport{
		ds_top: tmp,
	}
	return
}

func Newads_intersitetransportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_intersitetransport, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_intersitetransport{
		ds_top: tmp,
	}
	return
}

// SetDS_options sets the value of DS_options for the instance
func (instance *ads_intersitetransport) SetPropertyDS_options(value int32) (err error) {
	return instance.SetProperty("DS_options", (value))
}

// GetDS_options gets the value of DS_options for the instance
func (instance *ads_intersitetransport) GetPropertyDS_options() (value int32, err error) {
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
func (instance *ads_intersitetransport) SetPropertyDS_replInterval(value int32) (err error) {
	return instance.SetProperty("DS_replInterval", (value))
}

// GetDS_replInterval gets the value of DS_replInterval for the instance
func (instance *ads_intersitetransport) GetPropertyDS_replInterval() (value int32, err error) {
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

// SetDS_transportAddressAttribute sets the value of DS_transportAddressAttribute for the instance
func (instance *ads_intersitetransport) SetPropertyDS_transportAddressAttribute(value string) (err error) {
	return instance.SetProperty("DS_transportAddressAttribute", (value))
}

// GetDS_transportAddressAttribute gets the value of DS_transportAddressAttribute for the instance
func (instance *ads_intersitetransport) GetPropertyDS_transportAddressAttribute() (value string, err error) {
	retValue, err := instance.GetProperty("DS_transportAddressAttribute")
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

// SetDS_transportDLLName sets the value of DS_transportDLLName for the instance
func (instance *ads_intersitetransport) SetPropertyDS_transportDLLName(value string) (err error) {
	return instance.SetProperty("DS_transportDLLName", (value))
}

// GetDS_transportDLLName gets the value of DS_transportDLLName for the instance
func (instance *ads_intersitetransport) GetPropertyDS_transportDLLName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_transportDLLName")
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
