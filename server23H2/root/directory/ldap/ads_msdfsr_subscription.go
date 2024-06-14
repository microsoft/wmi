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

// ads_msdfsr_subscription struct
type ads_msdfsr_subscription struct {
	*ds_top

	//
	DS_msDFSR_CachePolicy int32

	//
	DS_msDFSR_ConflictPath string

	//
	DS_msDFSR_ConflictSizeInMb int64

	//
	DS_msDFSR_ContentSetGuid Uint8Array

	//
	DS_msDFSR_DeletedPath string

	//
	DS_msDFSR_DeletedSizeInMb int64

	//
	DS_msDFSR_DfsLinkTarget string

	//
	DS_msDFSR_Enabled bool

	//
	DS_msDFSR_Extension Uint8Array

	//
	DS_msDFSR_Flags int32

	//
	DS_msDFSR_MaxAgeInCacheInMin int32

	//
	DS_msDFSR_MinDurationCacheInMin int32

	//
	DS_msDFSR_OnDemandExclusionDirectoryFilter string

	//
	DS_msDFSR_OnDemandExclusionFileFilter string

	//
	DS_msDFSR_Options int32

	//
	DS_msDFSR_Options2 int32

	//
	DS_msDFSR_ReadOnly bool

	//
	DS_msDFSR_ReplicationGroupGuid Uint8Array

	//
	DS_msDFSR_RootFence int32

	//
	DS_msDFSR_RootPath string

	//
	DS_msDFSR_RootSizeInMb int64

	//
	DS_msDFSR_StagingCleanupTriggerInPercent int32

	//
	DS_msDFSR_StagingPath string

	//
	DS_msDFSR_StagingSizeInMb int64
}

func Newads_msdfsr_subscriptionEx1(instance *cim.WmiInstance) (newInstance *ads_msdfsr_subscription, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_subscription{
		ds_top: tmp,
	}
	return
}

func Newads_msdfsr_subscriptionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfsr_subscription, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_subscription{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFSR_CachePolicy sets the value of DS_msDFSR_CachePolicy for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_CachePolicy(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_CachePolicy", (value))
}

// GetDS_msDFSR_CachePolicy gets the value of DS_msDFSR_CachePolicy for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_CachePolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_CachePolicy")
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

// SetDS_msDFSR_ConflictPath sets the value of DS_msDFSR_ConflictPath for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_ConflictPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_ConflictPath", (value))
}

// GetDS_msDFSR_ConflictPath gets the value of DS_msDFSR_ConflictPath for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_ConflictPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ConflictPath")
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

// SetDS_msDFSR_ConflictSizeInMb sets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_ConflictSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_ConflictSizeInMb", (value))
}

// GetDS_msDFSR_ConflictSizeInMb gets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_ConflictSizeInMb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ConflictSizeInMb")
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

// SetDS_msDFSR_ContentSetGuid sets the value of DS_msDFSR_ContentSetGuid for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_ContentSetGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_ContentSetGuid", (value))
}

// GetDS_msDFSR_ContentSetGuid gets the value of DS_msDFSR_ContentSetGuid for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_ContentSetGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ContentSetGuid")
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

// SetDS_msDFSR_DeletedPath sets the value of DS_msDFSR_DeletedPath for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_DeletedPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DeletedPath", (value))
}

// GetDS_msDFSR_DeletedPath gets the value of DS_msDFSR_DeletedPath for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_DeletedPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DeletedPath")
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

// SetDS_msDFSR_DeletedSizeInMb sets the value of DS_msDFSR_DeletedSizeInMb for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_DeletedSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_DeletedSizeInMb", (value))
}

// GetDS_msDFSR_DeletedSizeInMb gets the value of DS_msDFSR_DeletedSizeInMb for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_DeletedSizeInMb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DeletedSizeInMb")
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

// SetDS_msDFSR_DfsLinkTarget sets the value of DS_msDFSR_DfsLinkTarget for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_DfsLinkTarget(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DfsLinkTarget", (value))
}

// GetDS_msDFSR_DfsLinkTarget gets the value of DS_msDFSR_DfsLinkTarget for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_DfsLinkTarget() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DfsLinkTarget")
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

// SetDS_msDFSR_Enabled sets the value of DS_msDFSR_Enabled for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_Enabled(value bool) (err error) {
	return instance.SetProperty("DS_msDFSR_Enabled", (value))
}

// GetDS_msDFSR_Enabled gets the value of DS_msDFSR_Enabled for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Enabled")
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

// SetDS_msDFSR_Extension sets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_Extension(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Extension", (value))
}

// GetDS_msDFSR_Extension gets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_Extension() (value Uint8Array, err error) {
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
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_Flags(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Flags", (value))
}

// GetDS_msDFSR_Flags gets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_Flags() (value int32, err error) {
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

// SetDS_msDFSR_MaxAgeInCacheInMin sets the value of DS_msDFSR_MaxAgeInCacheInMin for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_MaxAgeInCacheInMin(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_MaxAgeInCacheInMin", (value))
}

// GetDS_msDFSR_MaxAgeInCacheInMin gets the value of DS_msDFSR_MaxAgeInCacheInMin for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_MaxAgeInCacheInMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_MaxAgeInCacheInMin")
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

// SetDS_msDFSR_MinDurationCacheInMin sets the value of DS_msDFSR_MinDurationCacheInMin for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_MinDurationCacheInMin(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_MinDurationCacheInMin", (value))
}

// GetDS_msDFSR_MinDurationCacheInMin gets the value of DS_msDFSR_MinDurationCacheInMin for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_MinDurationCacheInMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_MinDurationCacheInMin")
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

// SetDS_msDFSR_OnDemandExclusionDirectoryFilter sets the value of DS_msDFSR_OnDemandExclusionDirectoryFilter for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionDirectoryFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionDirectoryFilter gets the value of DS_msDFSR_OnDemandExclusionDirectoryFilter for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_OnDemandExclusionDirectoryFilter")
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

// SetDS_msDFSR_OnDemandExclusionFileFilter sets the value of DS_msDFSR_OnDemandExclusionFileFilter for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_OnDemandExclusionFileFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionFileFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionFileFilter gets the value of DS_msDFSR_OnDemandExclusionFileFilter for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_OnDemandExclusionFileFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_OnDemandExclusionFileFilter")
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

// SetDS_msDFSR_Options sets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_Options(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options", (value))
}

// GetDS_msDFSR_Options gets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_Options() (value int32, err error) {
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
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_Options2(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options2", (value))
}

// GetDS_msDFSR_Options2 gets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_Options2() (value int32, err error) {
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

// SetDS_msDFSR_ReadOnly sets the value of DS_msDFSR_ReadOnly for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_ReadOnly(value bool) (err error) {
	return instance.SetProperty("DS_msDFSR_ReadOnly", (value))
}

// GetDS_msDFSR_ReadOnly gets the value of DS_msDFSR_ReadOnly for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_ReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ReadOnly")
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

// SetDS_msDFSR_ReplicationGroupGuid sets the value of DS_msDFSR_ReplicationGroupGuid for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_ReplicationGroupGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_ReplicationGroupGuid", (value))
}

// GetDS_msDFSR_ReplicationGroupGuid gets the value of DS_msDFSR_ReplicationGroupGuid for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_ReplicationGroupGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ReplicationGroupGuid")
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

// SetDS_msDFSR_RootFence sets the value of DS_msDFSR_RootFence for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_RootFence(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_RootFence", (value))
}

// GetDS_msDFSR_RootFence gets the value of DS_msDFSR_RootFence for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_RootFence() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_RootFence")
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

// SetDS_msDFSR_RootPath sets the value of DS_msDFSR_RootPath for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_RootPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_RootPath", (value))
}

// GetDS_msDFSR_RootPath gets the value of DS_msDFSR_RootPath for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_RootPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_RootPath")
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

// SetDS_msDFSR_RootSizeInMb sets the value of DS_msDFSR_RootSizeInMb for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_RootSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_RootSizeInMb", (value))
}

// GetDS_msDFSR_RootSizeInMb gets the value of DS_msDFSR_RootSizeInMb for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_RootSizeInMb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_RootSizeInMb")
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

// SetDS_msDFSR_StagingCleanupTriggerInPercent sets the value of DS_msDFSR_StagingCleanupTriggerInPercent for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_StagingCleanupTriggerInPercent(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingCleanupTriggerInPercent", (value))
}

// GetDS_msDFSR_StagingCleanupTriggerInPercent gets the value of DS_msDFSR_StagingCleanupTriggerInPercent for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_StagingCleanupTriggerInPercent() (value int32, err error) {
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

// SetDS_msDFSR_StagingPath sets the value of DS_msDFSR_StagingPath for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_StagingPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingPath", (value))
}

// GetDS_msDFSR_StagingPath gets the value of DS_msDFSR_StagingPath for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_StagingPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_StagingPath")
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

// SetDS_msDFSR_StagingSizeInMb sets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_subscription) SetPropertyDS_msDFSR_StagingSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingSizeInMb", (value))
}

// GetDS_msDFSR_StagingSizeInMb gets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_subscription) GetPropertyDS_msDFSR_StagingSizeInMb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_StagingSizeInMb")
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
