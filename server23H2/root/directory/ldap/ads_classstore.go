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

// ads_classstore struct
type ads_classstore struct {
	*ds_top

	//
	DS_appSchemaVersion int32

	//
	DS_lastUpdateSequence string

	//
	DS_nextLevelStore string

	//
	DS_versionNumber int32
}

func Newads_classstoreEx1(instance *cim.WmiInstance) (newInstance *ads_classstore, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_classstore{
		ds_top: tmp,
	}
	return
}

func Newads_classstoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_classstore, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_classstore{
		ds_top: tmp,
	}
	return
}

// SetDS_appSchemaVersion sets the value of DS_appSchemaVersion for the instance
func (instance *ads_classstore) SetPropertyDS_appSchemaVersion(value int32) (err error) {
	return instance.SetProperty("DS_appSchemaVersion", (value))
}

// GetDS_appSchemaVersion gets the value of DS_appSchemaVersion for the instance
func (instance *ads_classstore) GetPropertyDS_appSchemaVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_appSchemaVersion")
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

// SetDS_lastUpdateSequence sets the value of DS_lastUpdateSequence for the instance
func (instance *ads_classstore) SetPropertyDS_lastUpdateSequence(value string) (err error) {
	return instance.SetProperty("DS_lastUpdateSequence", (value))
}

// GetDS_lastUpdateSequence gets the value of DS_lastUpdateSequence for the instance
func (instance *ads_classstore) GetPropertyDS_lastUpdateSequence() (value string, err error) {
	retValue, err := instance.GetProperty("DS_lastUpdateSequence")
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

// SetDS_nextLevelStore sets the value of DS_nextLevelStore for the instance
func (instance *ads_classstore) SetPropertyDS_nextLevelStore(value string) (err error) {
	return instance.SetProperty("DS_nextLevelStore", (value))
}

// GetDS_nextLevelStore gets the value of DS_nextLevelStore for the instance
func (instance *ads_classstore) GetPropertyDS_nextLevelStore() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nextLevelStore")
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

// SetDS_versionNumber sets the value of DS_versionNumber for the instance
func (instance *ads_classstore) SetPropertyDS_versionNumber(value int32) (err error) {
	return instance.SetProperty("DS_versionNumber", (value))
}

// GetDS_versionNumber gets the value of DS_versionNumber for the instance
func (instance *ads_classstore) GetPropertyDS_versionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumber")
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
