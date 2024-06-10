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

// ads_msdfsr_localsettings struct
type ads_msdfsr_localsettings struct {
	*ds_top

	//
	DS_msDFSR_CommonStagingPath string

	//
	DS_msDFSR_CommonStagingSizeInMb int64

	//
	DS_msDFSR_Extension Uint8Array

	//
	DS_msDFSR_Flags int32

	//
	DS_msDFSR_Options int32

	//
	DS_msDFSR_Options2 int32

	//
	DS_msDFSR_StagingCleanupTriggerInPercent int32

	//
	DS_msDFSR_Version string
}

func Newads_msdfsr_localsettingsEx1(instance *cim.WmiInstance) (newInstance *ads_msdfsr_localsettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_localsettings{
		ds_top: tmp,
	}
	return
}

func Newads_msdfsr_localsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfsr_localsettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_localsettings{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFSR_CommonStagingPath sets the value of DS_msDFSR_CommonStagingPath for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_CommonStagingPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_CommonStagingPath", (value))
}

// GetDS_msDFSR_CommonStagingPath gets the value of DS_msDFSR_CommonStagingPath for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_CommonStagingPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_CommonStagingPath")
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

// SetDS_msDFSR_CommonStagingSizeInMb sets the value of DS_msDFSR_CommonStagingSizeInMb for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_CommonStagingSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_CommonStagingSizeInMb", (value))
}

// GetDS_msDFSR_CommonStagingSizeInMb gets the value of DS_msDFSR_CommonStagingSizeInMb for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_CommonStagingSizeInMb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_CommonStagingSizeInMb")
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

// SetDS_msDFSR_Extension sets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_Extension(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Extension", (value))
}

// GetDS_msDFSR_Extension gets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_Extension() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Extension")
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

// SetDS_msDFSR_Flags sets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_Flags(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Flags", (value))
}

// GetDS_msDFSR_Flags gets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_Flags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Flags")
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

// SetDS_msDFSR_Options sets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_Options(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options", (value))
}

// GetDS_msDFSR_Options gets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_Options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Options")
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

// SetDS_msDFSR_Options2 sets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_Options2(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options2", (value))
}

// GetDS_msDFSR_Options2 gets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_Options2() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Options2")
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

// SetDS_msDFSR_StagingCleanupTriggerInPercent sets the value of DS_msDFSR_StagingCleanupTriggerInPercent for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_StagingCleanupTriggerInPercent(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingCleanupTriggerInPercent", (value))
}

// GetDS_msDFSR_StagingCleanupTriggerInPercent gets the value of DS_msDFSR_StagingCleanupTriggerInPercent for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_StagingCleanupTriggerInPercent() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_StagingCleanupTriggerInPercent")
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

// SetDS_msDFSR_Version sets the value of DS_msDFSR_Version for the instance
func (instance *ads_msdfsr_localsettings) SetPropertyDS_msDFSR_Version(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_Version", (value))
}

// GetDS_msDFSR_Version gets the value of DS_msDFSR_Version for the instance
func (instance *ads_msdfsr_localsettings) GetPropertyDS_msDFSR_Version() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Version")
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
