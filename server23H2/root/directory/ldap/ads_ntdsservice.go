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

// ads_ntdsservice struct
type ads_ntdsservice struct {
	*ds_top

	//
	DS_dSHeuristics string

	//
	DS_garbageCollPeriod int32

	//
	DS_msDS_DeletedObjectLifetime int32

	//
	DS_msDS_Other_Settings []string

	//
	DS_replTopologyStayOfExecution int32

	//
	DS_sPNMappings []string

	//
	DS_tombstoneLifetime int32
}

func Newads_ntdsserviceEx1(instance *cim.WmiInstance) (newInstance *ads_ntdsservice, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsservice{
		ds_top: tmp,
	}
	return
}

func Newads_ntdsserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntdsservice, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsservice{
		ds_top: tmp,
	}
	return
}

// SetDS_dSHeuristics sets the value of DS_dSHeuristics for the instance
func (instance *ads_ntdsservice) SetPropertyDS_dSHeuristics(value string) (err error) {
	return instance.SetProperty("DS_dSHeuristics", (value))
}

// GetDS_dSHeuristics gets the value of DS_dSHeuristics for the instance
func (instance *ads_ntdsservice) GetPropertyDS_dSHeuristics() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dSHeuristics")
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

// SetDS_garbageCollPeriod sets the value of DS_garbageCollPeriod for the instance
func (instance *ads_ntdsservice) SetPropertyDS_garbageCollPeriod(value int32) (err error) {
	return instance.SetProperty("DS_garbageCollPeriod", (value))
}

// GetDS_garbageCollPeriod gets the value of DS_garbageCollPeriod for the instance
func (instance *ads_ntdsservice) GetPropertyDS_garbageCollPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_garbageCollPeriod")
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

// SetDS_msDS_DeletedObjectLifetime sets the value of DS_msDS_DeletedObjectLifetime for the instance
func (instance *ads_ntdsservice) SetPropertyDS_msDS_DeletedObjectLifetime(value int32) (err error) {
	return instance.SetProperty("DS_msDS_DeletedObjectLifetime", (value))
}

// GetDS_msDS_DeletedObjectLifetime gets the value of DS_msDS_DeletedObjectLifetime for the instance
func (instance *ads_ntdsservice) GetPropertyDS_msDS_DeletedObjectLifetime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DeletedObjectLifetime")
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

// SetDS_msDS_Other_Settings sets the value of DS_msDS_Other_Settings for the instance
func (instance *ads_ntdsservice) SetPropertyDS_msDS_Other_Settings(value []string) (err error) {
	return instance.SetProperty("DS_msDS_Other_Settings", (value))
}

// GetDS_msDS_Other_Settings gets the value of DS_msDS_Other_Settings for the instance
func (instance *ads_ntdsservice) GetPropertyDS_msDS_Other_Settings() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Other_Settings")
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

// SetDS_replTopologyStayOfExecution sets the value of DS_replTopologyStayOfExecution for the instance
func (instance *ads_ntdsservice) SetPropertyDS_replTopologyStayOfExecution(value int32) (err error) {
	return instance.SetProperty("DS_replTopologyStayOfExecution", (value))
}

// GetDS_replTopologyStayOfExecution gets the value of DS_replTopologyStayOfExecution for the instance
func (instance *ads_ntdsservice) GetPropertyDS_replTopologyStayOfExecution() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_replTopologyStayOfExecution")
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

// SetDS_sPNMappings sets the value of DS_sPNMappings for the instance
func (instance *ads_ntdsservice) SetPropertyDS_sPNMappings(value []string) (err error) {
	return instance.SetProperty("DS_sPNMappings", (value))
}

// GetDS_sPNMappings gets the value of DS_sPNMappings for the instance
func (instance *ads_ntdsservice) GetPropertyDS_sPNMappings() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_sPNMappings")
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

// SetDS_tombstoneLifetime sets the value of DS_tombstoneLifetime for the instance
func (instance *ads_ntdsservice) SetPropertyDS_tombstoneLifetime(value int32) (err error) {
	return instance.SetProperty("DS_tombstoneLifetime", (value))
}

// GetDS_tombstoneLifetime gets the value of DS_tombstoneLifetime for the instance
func (instance *ads_ntdsservice) GetPropertyDS_tombstoneLifetime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_tombstoneLifetime")
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
