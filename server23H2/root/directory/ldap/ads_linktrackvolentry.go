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

// ads_linktrackvolentry struct
type ads_linktrackvolentry struct {
	*ds_leaf

	//
	DS_currMachineId Uint8Array

	//
	DS_linkTrackSecret Uint8Array

	//
	DS_objectCount int32

	//
	DS_seqNotification int32

	//
	DS_timeRefresh int64

	//
	DS_timeVolChange int64

	//
	DS_volTableGUID Uint8Array

	//
	DS_volTableIdxGUID Uint8Array
}

func Newads_linktrackvolentryEx1(instance *cim.WmiInstance) (newInstance *ads_linktrackvolentry, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_linktrackvolentry{
		ds_leaf: tmp,
	}
	return
}

func Newads_linktrackvolentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_linktrackvolentry, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_linktrackvolentry{
		ds_leaf: tmp,
	}
	return
}

// SetDS_currMachineId sets the value of DS_currMachineId for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_currMachineId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_currMachineId", (value))
}

// GetDS_currMachineId gets the value of DS_currMachineId for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_currMachineId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_currMachineId")
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

// SetDS_linkTrackSecret sets the value of DS_linkTrackSecret for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_linkTrackSecret(value Uint8Array) (err error) {
	return instance.SetProperty("DS_linkTrackSecret", (value))
}

// GetDS_linkTrackSecret gets the value of DS_linkTrackSecret for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_linkTrackSecret() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_linkTrackSecret")
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

// SetDS_objectCount sets the value of DS_objectCount for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_objectCount(value int32) (err error) {
	return instance.SetProperty("DS_objectCount", (value))
}

// GetDS_objectCount gets the value of DS_objectCount for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_objectCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_objectCount")
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

// SetDS_seqNotification sets the value of DS_seqNotification for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_seqNotification(value int32) (err error) {
	return instance.SetProperty("DS_seqNotification", (value))
}

// GetDS_seqNotification gets the value of DS_seqNotification for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_seqNotification() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_seqNotification")
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

// SetDS_timeRefresh sets the value of DS_timeRefresh for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_timeRefresh(value int64) (err error) {
	return instance.SetProperty("DS_timeRefresh", (value))
}

// GetDS_timeRefresh gets the value of DS_timeRefresh for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_timeRefresh() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_timeRefresh")
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

// SetDS_timeVolChange sets the value of DS_timeVolChange for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_timeVolChange(value int64) (err error) {
	return instance.SetProperty("DS_timeVolChange", (value))
}

// GetDS_timeVolChange gets the value of DS_timeVolChange for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_timeVolChange() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_timeVolChange")
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

// SetDS_volTableGUID sets the value of DS_volTableGUID for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_volTableGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_volTableGUID", (value))
}

// GetDS_volTableGUID gets the value of DS_volTableGUID for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_volTableGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_volTableGUID")
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

// SetDS_volTableIdxGUID sets the value of DS_volTableIdxGUID for the instance
func (instance *ads_linktrackvolentry) SetPropertyDS_volTableIdxGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_volTableIdxGUID", (value))
}

// GetDS_volTableIdxGUID gets the value of DS_volTableIdxGUID for the instance
func (instance *ads_linktrackvolentry) GetPropertyDS_volTableIdxGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_volTableIdxGUID")
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
