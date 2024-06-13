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

// ads_volume struct
type ads_volume struct {
	*ds_connectionpoint

	//
	DS_contentIndexingAllowed bool

	//
	DS_lastContentIndexed int64

	//
	DS_uNCName string
}

func Newads_volumeEx1(instance *cim.WmiInstance) (newInstance *ads_volume, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_volume{
		ds_connectionpoint: tmp,
	}
	return
}

func Newads_volumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_volume, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_volume{
		ds_connectionpoint: tmp,
	}
	return
}

// SetDS_contentIndexingAllowed sets the value of DS_contentIndexingAllowed for the instance
func (instance *ads_volume) SetPropertyDS_contentIndexingAllowed(value bool) (err error) {
	return instance.SetProperty("DS_contentIndexingAllowed", (value))
}

// GetDS_contentIndexingAllowed gets the value of DS_contentIndexingAllowed for the instance
func (instance *ads_volume) GetPropertyDS_contentIndexingAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_contentIndexingAllowed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDS_lastContentIndexed sets the value of DS_lastContentIndexed for the instance
func (instance *ads_volume) SetPropertyDS_lastContentIndexed(value int64) (err error) {
	return instance.SetProperty("DS_lastContentIndexed", (value))
}

// GetDS_lastContentIndexed gets the value of DS_lastContentIndexed for the instance
func (instance *ads_volume) GetPropertyDS_lastContentIndexed() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastContentIndexed")
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

// SetDS_uNCName sets the value of DS_uNCName for the instance
func (instance *ads_volume) SetPropertyDS_uNCName(value string) (err error) {
	return instance.SetProperty("DS_uNCName", (value))
}

// GetDS_uNCName gets the value of DS_uNCName for the instance
func (instance *ads_volume) GetPropertyDS_uNCName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_uNCName")
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
