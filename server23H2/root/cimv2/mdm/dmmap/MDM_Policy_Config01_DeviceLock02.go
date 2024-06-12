// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_Policy_Config01_DeviceLock02 struct
type MDM_Policy_Config01_DeviceLock02 struct {
	*cim.WmiInstance

	//
	AllowScreenTimeoutWhileLockedUserConfig int32

	//
	AllowSimpleDevicePassword int32

	//
	AlphanumericDevicePasswordRequired int32

	//
	DevicePasswordEnabled int32

	//
	DevicePasswordExpiration int32

	//
	DevicePasswordHistory int32

	//
	EnforceLockScreenAndLogonImage string

	//
	EnforceLockScreenProvider string

	//
	InstanceID string

	//
	MaxDevicePasswordFailedAttempts int32

	//
	MaxInactivityTimeDeviceLock int32

	//
	MinDevicePasswordComplexCharacters int32

	//
	MinDevicePasswordLength int32

	//
	MinimumPasswordAge int32

	//
	ParentID string

	//
	PreventEnablingLockScreenCamera string

	//
	PreventLockScreenSlideShow string

	//
	ScreenTimeoutWhileLocked int32
}

func NewMDM_Policy_Config01_DeviceLock02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_DeviceLock02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceLock02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_DeviceLock02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_DeviceLock02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceLock02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowScreenTimeoutWhileLockedUserConfig sets the value of AllowScreenTimeoutWhileLockedUserConfig for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyAllowScreenTimeoutWhileLockedUserConfig(value int32) (err error) {
	return instance.SetProperty("AllowScreenTimeoutWhileLockedUserConfig", (value))
}

// GetAllowScreenTimeoutWhileLockedUserConfig gets the value of AllowScreenTimeoutWhileLockedUserConfig for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyAllowScreenTimeoutWhileLockedUserConfig() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowScreenTimeoutWhileLockedUserConfig")
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

// SetAllowSimpleDevicePassword sets the value of AllowSimpleDevicePassword for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyAllowSimpleDevicePassword(value int32) (err error) {
	return instance.SetProperty("AllowSimpleDevicePassword", (value))
}

// GetAllowSimpleDevicePassword gets the value of AllowSimpleDevicePassword for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyAllowSimpleDevicePassword() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSimpleDevicePassword")
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

// SetAlphanumericDevicePasswordRequired sets the value of AlphanumericDevicePasswordRequired for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyAlphanumericDevicePasswordRequired(value int32) (err error) {
	return instance.SetProperty("AlphanumericDevicePasswordRequired", (value))
}

// GetAlphanumericDevicePasswordRequired gets the value of AlphanumericDevicePasswordRequired for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyAlphanumericDevicePasswordRequired() (value int32, err error) {
	retValue, err := instance.GetProperty("AlphanumericDevicePasswordRequired")
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

// SetDevicePasswordEnabled sets the value of DevicePasswordEnabled for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyDevicePasswordEnabled(value int32) (err error) {
	return instance.SetProperty("DevicePasswordEnabled", (value))
}

// GetDevicePasswordEnabled gets the value of DevicePasswordEnabled for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyDevicePasswordEnabled() (value int32, err error) {
	retValue, err := instance.GetProperty("DevicePasswordEnabled")
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

// SetDevicePasswordExpiration sets the value of DevicePasswordExpiration for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyDevicePasswordExpiration(value int32) (err error) {
	return instance.SetProperty("DevicePasswordExpiration", (value))
}

// GetDevicePasswordExpiration gets the value of DevicePasswordExpiration for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyDevicePasswordExpiration() (value int32, err error) {
	retValue, err := instance.GetProperty("DevicePasswordExpiration")
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

// SetDevicePasswordHistory sets the value of DevicePasswordHistory for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyDevicePasswordHistory(value int32) (err error) {
	return instance.SetProperty("DevicePasswordHistory", (value))
}

// GetDevicePasswordHistory gets the value of DevicePasswordHistory for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyDevicePasswordHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("DevicePasswordHistory")
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

// SetEnforceLockScreenAndLogonImage sets the value of EnforceLockScreenAndLogonImage for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyEnforceLockScreenAndLogonImage(value string) (err error) {
	return instance.SetProperty("EnforceLockScreenAndLogonImage", (value))
}

// GetEnforceLockScreenAndLogonImage gets the value of EnforceLockScreenAndLogonImage for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyEnforceLockScreenAndLogonImage() (value string, err error) {
	retValue, err := instance.GetProperty("EnforceLockScreenAndLogonImage")
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

// SetEnforceLockScreenProvider sets the value of EnforceLockScreenProvider for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyEnforceLockScreenProvider(value string) (err error) {
	return instance.SetProperty("EnforceLockScreenProvider", (value))
}

// GetEnforceLockScreenProvider gets the value of EnforceLockScreenProvider for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyEnforceLockScreenProvider() (value string, err error) {
	retValue, err := instance.GetProperty("EnforceLockScreenProvider")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyInstanceID() (value string, err error) {
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

// SetMaxDevicePasswordFailedAttempts sets the value of MaxDevicePasswordFailedAttempts for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyMaxDevicePasswordFailedAttempts(value int32) (err error) {
	return instance.SetProperty("MaxDevicePasswordFailedAttempts", (value))
}

// GetMaxDevicePasswordFailedAttempts gets the value of MaxDevicePasswordFailedAttempts for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyMaxDevicePasswordFailedAttempts() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxDevicePasswordFailedAttempts")
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

// SetMaxInactivityTimeDeviceLock sets the value of MaxInactivityTimeDeviceLock for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyMaxInactivityTimeDeviceLock(value int32) (err error) {
	return instance.SetProperty("MaxInactivityTimeDeviceLock", (value))
}

// GetMaxInactivityTimeDeviceLock gets the value of MaxInactivityTimeDeviceLock for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyMaxInactivityTimeDeviceLock() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxInactivityTimeDeviceLock")
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

// SetMinDevicePasswordComplexCharacters sets the value of MinDevicePasswordComplexCharacters for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyMinDevicePasswordComplexCharacters(value int32) (err error) {
	return instance.SetProperty("MinDevicePasswordComplexCharacters", (value))
}

// GetMinDevicePasswordComplexCharacters gets the value of MinDevicePasswordComplexCharacters for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyMinDevicePasswordComplexCharacters() (value int32, err error) {
	retValue, err := instance.GetProperty("MinDevicePasswordComplexCharacters")
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

// SetMinDevicePasswordLength sets the value of MinDevicePasswordLength for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyMinDevicePasswordLength(value int32) (err error) {
	return instance.SetProperty("MinDevicePasswordLength", (value))
}

// GetMinDevicePasswordLength gets the value of MinDevicePasswordLength for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyMinDevicePasswordLength() (value int32, err error) {
	retValue, err := instance.GetProperty("MinDevicePasswordLength")
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

// SetMinimumPasswordAge sets the value of MinimumPasswordAge for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyMinimumPasswordAge(value int32) (err error) {
	return instance.SetProperty("MinimumPasswordAge", (value))
}

// GetMinimumPasswordAge gets the value of MinimumPasswordAge for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyMinimumPasswordAge() (value int32, err error) {
	retValue, err := instance.GetProperty("MinimumPasswordAge")
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
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyParentID() (value string, err error) {
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

// SetPreventEnablingLockScreenCamera sets the value of PreventEnablingLockScreenCamera for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyPreventEnablingLockScreenCamera(value string) (err error) {
	return instance.SetProperty("PreventEnablingLockScreenCamera", (value))
}

// GetPreventEnablingLockScreenCamera gets the value of PreventEnablingLockScreenCamera for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyPreventEnablingLockScreenCamera() (value string, err error) {
	retValue, err := instance.GetProperty("PreventEnablingLockScreenCamera")
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

// SetPreventLockScreenSlideShow sets the value of PreventLockScreenSlideShow for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyPreventLockScreenSlideShow(value string) (err error) {
	return instance.SetProperty("PreventLockScreenSlideShow", (value))
}

// GetPreventLockScreenSlideShow gets the value of PreventLockScreenSlideShow for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyPreventLockScreenSlideShow() (value string, err error) {
	retValue, err := instance.GetProperty("PreventLockScreenSlideShow")
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

// SetScreenTimeoutWhileLocked sets the value of ScreenTimeoutWhileLocked for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) SetPropertyScreenTimeoutWhileLocked(value int32) (err error) {
	return instance.SetProperty("ScreenTimeoutWhileLocked", (value))
}

// GetScreenTimeoutWhileLocked gets the value of ScreenTimeoutWhileLocked for the instance
func (instance *MDM_Policy_Config01_DeviceLock02) GetPropertyScreenTimeoutWhileLocked() (value int32, err error) {
	retValue, err := instance.GetProperty("ScreenTimeoutWhileLocked")
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
