// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_SharedPC struct
type MDM_SharedPC struct {
	*cim.WmiInstance

	//
	AccountModel int32

	//
	DeletionPolicy int32

	//
	DiskLevelCaching int32

	//
	DiskLevelDeletion int32

	//
	EnableAccountManager bool

	//
	EnableSharedPCMode bool

	//
	InactiveThreshold int32

	//
	InstanceID string

	//
	KioskModeAUMID string

	//
	KioskModeUserTileDisplayText string

	//
	MaintenanceStartTime int32

	//
	MaxPageFileSizeMB int32

	//
	ParentID string

	//
	RestrictLocalStorage bool

	//
	SetEduPolicies bool

	//
	SetPowerPolicies bool

	//
	SignInOnResume bool

	//
	SleepTimeout int32
}

func NewMDM_SharedPCEx1(instance *cim.WmiInstance) (newInstance *MDM_SharedPC, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_SharedPC{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_SharedPCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_SharedPC, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_SharedPC{
		WmiInstance: tmp,
	}
	return
}

// SetAccountModel sets the value of AccountModel for the instance
func (instance *MDM_SharedPC) SetPropertyAccountModel(value int32) (err error) {
	return instance.SetProperty("AccountModel", (value))
}

// GetAccountModel gets the value of AccountModel for the instance
func (instance *MDM_SharedPC) GetPropertyAccountModel() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountModel")
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

// SetDeletionPolicy sets the value of DeletionPolicy for the instance
func (instance *MDM_SharedPC) SetPropertyDeletionPolicy(value int32) (err error) {
	return instance.SetProperty("DeletionPolicy", (value))
}

// GetDeletionPolicy gets the value of DeletionPolicy for the instance
func (instance *MDM_SharedPC) GetPropertyDeletionPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("DeletionPolicy")
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

// SetDiskLevelCaching sets the value of DiskLevelCaching for the instance
func (instance *MDM_SharedPC) SetPropertyDiskLevelCaching(value int32) (err error) {
	return instance.SetProperty("DiskLevelCaching", (value))
}

// GetDiskLevelCaching gets the value of DiskLevelCaching for the instance
func (instance *MDM_SharedPC) GetPropertyDiskLevelCaching() (value int32, err error) {
	retValue, err := instance.GetProperty("DiskLevelCaching")
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

// SetDiskLevelDeletion sets the value of DiskLevelDeletion for the instance
func (instance *MDM_SharedPC) SetPropertyDiskLevelDeletion(value int32) (err error) {
	return instance.SetProperty("DiskLevelDeletion", (value))
}

// GetDiskLevelDeletion gets the value of DiskLevelDeletion for the instance
func (instance *MDM_SharedPC) GetPropertyDiskLevelDeletion() (value int32, err error) {
	retValue, err := instance.GetProperty("DiskLevelDeletion")
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

// SetEnableAccountManager sets the value of EnableAccountManager for the instance
func (instance *MDM_SharedPC) SetPropertyEnableAccountManager(value bool) (err error) {
	return instance.SetProperty("EnableAccountManager", (value))
}

// GetEnableAccountManager gets the value of EnableAccountManager for the instance
func (instance *MDM_SharedPC) GetPropertyEnableAccountManager() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableAccountManager")
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

// SetEnableSharedPCMode sets the value of EnableSharedPCMode for the instance
func (instance *MDM_SharedPC) SetPropertyEnableSharedPCMode(value bool) (err error) {
	return instance.SetProperty("EnableSharedPCMode", (value))
}

// GetEnableSharedPCMode gets the value of EnableSharedPCMode for the instance
func (instance *MDM_SharedPC) GetPropertyEnableSharedPCMode() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSharedPCMode")
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

// SetInactiveThreshold sets the value of InactiveThreshold for the instance
func (instance *MDM_SharedPC) SetPropertyInactiveThreshold(value int32) (err error) {
	return instance.SetProperty("InactiveThreshold", (value))
}

// GetInactiveThreshold gets the value of InactiveThreshold for the instance
func (instance *MDM_SharedPC) GetPropertyInactiveThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("InactiveThreshold")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_SharedPC) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_SharedPC) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetKioskModeAUMID sets the value of KioskModeAUMID for the instance
func (instance *MDM_SharedPC) SetPropertyKioskModeAUMID(value string) (err error) {
	return instance.SetProperty("KioskModeAUMID", (value))
}

// GetKioskModeAUMID gets the value of KioskModeAUMID for the instance
func (instance *MDM_SharedPC) GetPropertyKioskModeAUMID() (value string, err error) {
	retValue, err := instance.GetProperty("KioskModeAUMID")
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

// SetKioskModeUserTileDisplayText sets the value of KioskModeUserTileDisplayText for the instance
func (instance *MDM_SharedPC) SetPropertyKioskModeUserTileDisplayText(value string) (err error) {
	return instance.SetProperty("KioskModeUserTileDisplayText", (value))
}

// GetKioskModeUserTileDisplayText gets the value of KioskModeUserTileDisplayText for the instance
func (instance *MDM_SharedPC) GetPropertyKioskModeUserTileDisplayText() (value string, err error) {
	retValue, err := instance.GetProperty("KioskModeUserTileDisplayText")
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

// SetMaintenanceStartTime sets the value of MaintenanceStartTime for the instance
func (instance *MDM_SharedPC) SetPropertyMaintenanceStartTime(value int32) (err error) {
	return instance.SetProperty("MaintenanceStartTime", (value))
}

// GetMaintenanceStartTime gets the value of MaintenanceStartTime for the instance
func (instance *MDM_SharedPC) GetPropertyMaintenanceStartTime() (value int32, err error) {
	retValue, err := instance.GetProperty("MaintenanceStartTime")
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

// SetMaxPageFileSizeMB sets the value of MaxPageFileSizeMB for the instance
func (instance *MDM_SharedPC) SetPropertyMaxPageFileSizeMB(value int32) (err error) {
	return instance.SetProperty("MaxPageFileSizeMB", (value))
}

// GetMaxPageFileSizeMB gets the value of MaxPageFileSizeMB for the instance
func (instance *MDM_SharedPC) GetPropertyMaxPageFileSizeMB() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxPageFileSizeMB")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_SharedPC) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_SharedPC) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetRestrictLocalStorage sets the value of RestrictLocalStorage for the instance
func (instance *MDM_SharedPC) SetPropertyRestrictLocalStorage(value bool) (err error) {
	return instance.SetProperty("RestrictLocalStorage", (value))
}

// GetRestrictLocalStorage gets the value of RestrictLocalStorage for the instance
func (instance *MDM_SharedPC) GetPropertyRestrictLocalStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("RestrictLocalStorage")
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

// SetSetEduPolicies sets the value of SetEduPolicies for the instance
func (instance *MDM_SharedPC) SetPropertySetEduPolicies(value bool) (err error) {
	return instance.SetProperty("SetEduPolicies", (value))
}

// GetSetEduPolicies gets the value of SetEduPolicies for the instance
func (instance *MDM_SharedPC) GetPropertySetEduPolicies() (value bool, err error) {
	retValue, err := instance.GetProperty("SetEduPolicies")
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

// SetSetPowerPolicies sets the value of SetPowerPolicies for the instance
func (instance *MDM_SharedPC) SetPropertySetPowerPolicies(value bool) (err error) {
	return instance.SetProperty("SetPowerPolicies", (value))
}

// GetSetPowerPolicies gets the value of SetPowerPolicies for the instance
func (instance *MDM_SharedPC) GetPropertySetPowerPolicies() (value bool, err error) {
	retValue, err := instance.GetProperty("SetPowerPolicies")
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

// SetSignInOnResume sets the value of SignInOnResume for the instance
func (instance *MDM_SharedPC) SetPropertySignInOnResume(value bool) (err error) {
	return instance.SetProperty("SignInOnResume", (value))
}

// GetSignInOnResume gets the value of SignInOnResume for the instance
func (instance *MDM_SharedPC) GetPropertySignInOnResume() (value bool, err error) {
	retValue, err := instance.GetProperty("SignInOnResume")
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

// SetSleepTimeout sets the value of SleepTimeout for the instance
func (instance *MDM_SharedPC) SetPropertySleepTimeout(value int32) (err error) {
	return instance.SetProperty("SleepTimeout", (value))
}

// GetSleepTimeout gets the value of SleepTimeout for the instance
func (instance *MDM_SharedPC) GetPropertySleepTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("SleepTimeout")
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
