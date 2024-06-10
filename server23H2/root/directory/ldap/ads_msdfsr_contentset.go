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

// ads_msdfsr_contentset struct
type ads_msdfsr_contentset struct {
	*ds_top

	//
	DS_msDFSR_ConflictSizeInMb int64

	//
	DS_msDFSR_DefaultCompressionExclusionFilter string

	//
	DS_msDFSR_DeletedSizeInMb int64

	//
	DS_msDFSR_DfsPath string

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
	DS_msDFSR_Priority int32

	//
	DS_msDFSR_RootSizeInMb int64

	//
	DS_msDFSR_StagingSizeInMb int64
}

func Newads_msdfsr_contentsetEx1(instance *cim.WmiInstance) (newInstance *ads_msdfsr_contentset, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_contentset{
		ds_top: tmp,
	}
	return
}

func Newads_msdfsr_contentsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfsr_contentset, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_contentset{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFSR_ConflictSizeInMb sets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_ConflictSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_ConflictSizeInMb", (value))
}

// GetDS_msDFSR_ConflictSizeInMb gets the value of DS_msDFSR_ConflictSizeInMb for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_ConflictSizeInMb() (value int64, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_DefaultCompressionExclusionFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DefaultCompressionExclusionFilter", (value))
}

// GetDS_msDFSR_DefaultCompressionExclusionFilter gets the value of DS_msDFSR_DefaultCompressionExclusionFilter for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_DefaultCompressionExclusionFilter() (value string, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_DeletedSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_DeletedSizeInMb", (value))
}

// GetDS_msDFSR_DeletedSizeInMb gets the value of DS_msDFSR_DeletedSizeInMb for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_DeletedSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_DfsPath sets the value of DS_msDFSR_DfsPath for the instance
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_DfsPath(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DfsPath", (value))
}

// GetDS_msDFSR_DfsPath gets the value of DS_msDFSR_DfsPath for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_DfsPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DfsPath")
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

// SetDS_msDFSR_DirectoryFilter sets the value of DS_msDFSR_DirectoryFilter for the instance
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_DirectoryFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_DirectoryFilter", (value))
}

// GetDS_msDFSR_DirectoryFilter gets the value of DS_msDFSR_DirectoryFilter for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_DirectoryFilter() (value string, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_Extension(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Extension", (value))
}

// GetDS_msDFSR_Extension gets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_Extension() (value Uint8Array, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_FileFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_FileFilter", (value))
}

// GetDS_msDFSR_FileFilter gets the value of DS_msDFSR_FileFilter for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_FileFilter() (value string, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_Flags(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Flags", (value))
}

// GetDS_msDFSR_Flags gets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_Flags() (value int32, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionDirectoryFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionDirectoryFilter gets the value of DS_msDFSR_OnDemandExclusionDirectoryFilter for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_OnDemandExclusionDirectoryFilter() (value string, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_OnDemandExclusionFileFilter(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_OnDemandExclusionFileFilter", (value))
}

// GetDS_msDFSR_OnDemandExclusionFileFilter gets the value of DS_msDFSR_OnDemandExclusionFileFilter for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_OnDemandExclusionFileFilter() (value string, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_Options(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options", (value))
}

// GetDS_msDFSR_Options gets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_Options() (value int32, err error) {
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_Options2(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options2", (value))
}

// GetDS_msDFSR_Options2 gets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_Options2() (value int32, err error) {
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

// SetDS_msDFSR_Priority sets the value of DS_msDFSR_Priority for the instance
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_Priority(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Priority", (value))
}

// GetDS_msDFSR_Priority gets the value of DS_msDFSR_Priority for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_Priority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Priority")
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
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_RootSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_RootSizeInMb", (value))
}

// GetDS_msDFSR_RootSizeInMb gets the value of DS_msDFSR_RootSizeInMb for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_RootSizeInMb() (value int64, err error) {
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

// SetDS_msDFSR_StagingSizeInMb sets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_contentset) SetPropertyDS_msDFSR_StagingSizeInMb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_StagingSizeInMb", (value))
}

// GetDS_msDFSR_StagingSizeInMb gets the value of DS_msDFSR_StagingSizeInMb for the instance
func (instance *ads_msdfsr_contentset) GetPropertyDS_msDFSR_StagingSizeInMb() (value int64, err error) {
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
