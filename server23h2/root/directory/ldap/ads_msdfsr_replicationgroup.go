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

// ads_msdfsr_replicationgroup struct
type ads_msdfsr_replicationgroup struct {
	*ds_top

	//
	DS_msDFSR_ConflictSizeInMb int64

	//
	DS_msDFSR_DefaultCompressionExclusionFilter string

	//
	DS_msDFSR_DeletedSizeInMb int64

	//
	DS_msDFSR_DirectoryFilter string

	//
	DS_msDFSR_Extension Uint8Array

	//
	DS_msDFSR_FileFilter string

	//
	DS_msDFSR_Flags int32

	//
	DS_msDFSR_OnDemandExclusionDirectoryFilter string

	//
	DS_msDFSR_OnDemandExclusionFileFilter string

	//
	DS_msDFSR_Options int32

	//
	DS_msDFSR_Options2 int32

	//
	DS_msDFSR_ReplicationGroupType int32

	//
	DS_msDFSR_RootSizeInMb int64

	//
	DS_msDFSR_Schedule Uint8Array

	//
	DS_msDFSR_StagingSizeInMb int64

	//
	DS_msDFSR_TombstoneExpiryInMin int32

	//
	DS_msDFSR_Version string
}

func Newads_msdfsr_replicationgroupEx1(instance *cim.WmiInstance) (newInstance *ads_msdfsr_replicationgroup, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_replicationgroup{
		ds_top: tmp,
	}
	return
}

func Newads_msdfsr_replicationgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfsr_replicationgroup, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_replicationgroup{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFSR_ConflictSizeInMb sets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_ConflictSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_ConflictSizeInMb", (value))
}

// GetDS_msDFSR_ConflictSizeInMb gets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_ConflictSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_DefaultCompressionExclusionFilter sets the value of DS_msDFSR_DefaultCompressionExclusionFilter for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_DefaultCompressionExclusionFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DefaultCompressionExclusionFilter", (value))
}

// GetDS_msDFSR_DefaultCompressionExclusionFilter gets the value of DS_msDFSR_DefaultCompressionExclusionFilter for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_DefaultCompressionExclusionFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DefaultCompressionExclusionFilter")
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
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_DeletedSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_DeletedSizeInMb", (value))
}

// GetDS_msDFSR_DeletedSizeInMb gets the value of DS_msDFSR_DeletedSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_DeletedSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_DirectoryFilter sets the value of DS_msDFSR_DirectoryFilter for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_DirectoryFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DirectoryFilter", (value))
}

// GetDS_msDFSR_DirectoryFilter gets the value of DS_msDFSR_DirectoryFilter for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_DirectoryFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DirectoryFilter")
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

// SetDS_msDFSR_Extension sets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Extension(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Extension", (value))
}

// GetDS_msDFSR_Extension gets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Extension() (value Uint8Array, err error) {
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

// SetDS_msDFSR_FileFilter sets the value of DS_msDFSR_FileFilter for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_FileFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_FileFilter", (value))
}

// GetDS_msDFSR_FileFilter gets the value of DS_msDFSR_FileFilter for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_FileFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_FileFilter")
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

// SetDS_msDFSR_Flags sets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Flags(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Flags", (value))
}

// GetDS_msDFSR_Flags gets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Flags() (value int32, err error) {
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

// SetDS_msDFSR_OnDemandExclusionDirectoryFilter sets the value of DS_msDFSR_OnDemandExclusionDirectoryFilter for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionDirectoryFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionDirectoryFilter gets the value of DS_msDFSR_OnDemandExclusionDirectoryFilter for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter() (value string, err error) {
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
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_OnDemandExclusionFileFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionFileFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionFileFilter gets the value of DS_msDFSR_OnDemandExclusionFileFilter for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_OnDemandExclusionFileFilter() (value string, err error) {
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
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Options(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options", (value))
}

// GetDS_msDFSR_Options gets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Options() (value int32, err error) {
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
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Options2(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options2", (value))
}

// GetDS_msDFSR_Options2 gets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Options2() (value int32, err error) {
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

// SetDS_msDFSR_ReplicationGroupType sets the value of DS_msDFSR_ReplicationGroupType for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_ReplicationGroupType(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_ReplicationGroupType", (value))
}

// GetDS_msDFSR_ReplicationGroupType gets the value of DS_msDFSR_ReplicationGroupType for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_ReplicationGroupType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ReplicationGroupType")
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

// SetDS_msDFSR_RootSizeInMb sets the value of DS_msDFSR_RootSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_RootSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_RootSizeInMb", (value))
}

// GetDS_msDFSR_RootSizeInMb gets the value of DS_msDFSR_RootSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_RootSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_Schedule sets the value of DS_msDFSR_Schedule for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Schedule", (value))
}

// GetDS_msDFSR_Schedule gets the value of DS_msDFSR_Schedule for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Schedule")
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

// SetDS_msDFSR_StagingSizeInMb sets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_StagingSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingSizeInMb", (value))
}

// GetDS_msDFSR_StagingSizeInMb gets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_StagingSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_TombstoneExpiryInMin sets the value of DS_msDFSR_TombstoneExpiryInMin for the instance
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_TombstoneExpiryInMin(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_TombstoneExpiryInMin", (value))
}

// GetDS_msDFSR_TombstoneExpiryInMin gets the value of DS_msDFSR_TombstoneExpiryInMin for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_TombstoneExpiryInMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_TombstoneExpiryInMin")
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
func (instance *ads_msdfsr_replicationgroup) SetPropertyDS_msDFSR_Version(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_Version", (value))
}

// GetDS_msDFSR_Version gets the value of DS_msDFSR_Version for the instance
func (instance *ads_msdfsr_replicationgroup) GetPropertyDS_msDFSR_Version() (value string, err error) {
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
