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

// ads_subnet struct
type ads_subnet struct {
	*ds_top

	//
	DS_location string

	//
	DS_physicalLocationObject string

	//
	DS_siteObject string
}

func Newads_subnetEx1(instance *cim.WmiInstance) (newInstance *ads_subnet, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_subnet{
		ds_top: tmp,
	}
	return
}

func Newads_subnetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_subnet, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_subnet{
		ds_top: tmp,
	}
	return
}

// SetDS_location sets the value of DS_location for the instance
func (instance *ads_subnet) SetPropertyDS_location(value string) (err error) {
	return instance.SetProperty("DS_location", (value))
}

// GetDS_location gets the value of DS_location for the instance
func (instance *ads_subnet) GetPropertyDS_location() (value string, err error) {
	retValue, err := instance.GetProperty("DS_location")
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

// SetDS_physicalLocationObject sets the value of DS_physicalLocationObject for the instance
func (instance *ads_subnet) SetPropertyDS_physicalLocationObject(value string) (err error) {
	return instance.SetProperty("DS_physicalLocationObject", (value))
}

// GetDS_physicalLocationObject gets the value of DS_physicalLocationObject for the instance
func (instance *ads_subnet) GetPropertyDS_physicalLocationObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_physicalLocationObject")
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

// SetDS_siteObject sets the value of DS_siteObject for the instance
func (instance *ads_subnet) SetPropertyDS_siteObject(value string) (err error) {
	return instance.SetProperty("DS_siteObject", (value))
}

// GetDS_siteObject gets the value of DS_siteObject for the instance
func (instance *ads_subnet) GetPropertyDS_siteObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_siteObject")
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
