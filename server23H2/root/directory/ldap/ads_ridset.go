// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ridset struct
type ads_ridset struct {
	*ds_top

	//
	DS_rIDAllocationPool int64

	//
	DS_rIDNextRID int32

	//
	DS_rIDPreviousAllocationPool int64

	//
	DS_rIDUsedPool int64
}

func Newads_ridsetEx1(instance *cim.WmiInstance) (newInstance *ads_ridset, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ridset{
		ds_top: tmp,
	}
	return
}

func Newads_ridsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ridset, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ridset{
		ds_top: tmp,
	}
	return
}

// SetDS_rIDAllocationPool sets the value of DS_rIDAllocationPool for the instance
func (instance *ads_ridset) SetPropertyDS_rIDAllocationPool(value int64) (err error) {
	return instance.SetProperty("DS_rIDAllocationPool", (value))
}

// GetDS_rIDAllocationPool gets the value of DS_rIDAllocationPool for the instance
func (instance *ads_ridset) GetPropertyDS_rIDAllocationPool() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_rIDAllocationPool")
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

// SetDS_rIDNextRID sets the value of DS_rIDNextRID for the instance
func (instance *ads_ridset) SetPropertyDS_rIDNextRID(value int32) (err error) {
	return instance.SetProperty("DS_rIDNextRID", (value))
}

// GetDS_rIDNextRID gets the value of DS_rIDNextRID for the instance
func (instance *ads_ridset) GetPropertyDS_rIDNextRID() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_rIDNextRID")
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

// SetDS_rIDPreviousAllocationPool sets the value of DS_rIDPreviousAllocationPool for the instance
func (instance *ads_ridset) SetPropertyDS_rIDPreviousAllocationPool(value int64) (err error) {
	return instance.SetProperty("DS_rIDPreviousAllocationPool", (value))
}

// GetDS_rIDPreviousAllocationPool gets the value of DS_rIDPreviousAllocationPool for the instance
func (instance *ads_ridset) GetPropertyDS_rIDPreviousAllocationPool() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_rIDPreviousAllocationPool")
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

// SetDS_rIDUsedPool sets the value of DS_rIDUsedPool for the instance
func (instance *ads_ridset) SetPropertyDS_rIDUsedPool(value int64) (err error) {
	return instance.SetProperty("DS_rIDUsedPool", (value))
}

// GetDS_rIDUsedPool gets the value of DS_rIDUsedPool for the instance
func (instance *ads_ridset) GetPropertyDS_rIDUsedPool() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_rIDUsedPool")
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
