// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_BitLocker struct
type MDM_BitLocker struct {
	*cim.WmiInstance

	//
	AllowStandardUserEncryption int32

	//
	AllowWarningForOtherDiskEncryption int32

	//
	ConfigureRecoveryPasswordRotation int32

	//
	EncryptionMethodByDriveType string

	//
	FixedDrivesEncryptionType string

	//
	FixedDrivesRecoveryOptions string

	//
	FixedDrivesRequireEncryption string

	//
	IdentificationField string

	//
	InstanceID string

	//
	ParentID string

	//
	RemovableDrivesConfigureBDE string

	//
	RemovableDrivesEncryptionType string

	//
	RemovableDrivesExcludedFromEncryption string

	//
	RemovableDrivesRequireEncryption string

	//
	RequireDeviceEncryption int32

	//
	RequireStorageCardEncryption int32

	//
	SystemDrivesDisallowStandardUsersCanChangePIN string

	//
	SystemDrivesEnablePrebootInputProtectorsOnSlates string

	//
	SystemDrivesEnablePreBootPinExceptionOnDECapableDevice string

	//
	SystemDrivesEncryptionType string

	//
	SystemDrivesEnhancedPIN string

	//
	SystemDrivesMinimumPINLength string

	//
	SystemDrivesRecoveryMessage string

	//
	SystemDrivesRecoveryOptions string

	//
	SystemDrivesRequireStartupAuthentication string
}

func NewMDM_BitLockerEx1(instance *cim.WmiInstance) (newInstance *MDM_BitLocker, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_BitLocker{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_BitLockerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_BitLocker, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_BitLocker{
		WmiInstance: tmp,
	}
	return
}

// SetAllowStandardUserEncryption sets the value of AllowStandardUserEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyAllowStandardUserEncryption(value int32) (err error) {
	return instance.SetProperty("AllowStandardUserEncryption", (value))
}

// GetAllowStandardUserEncryption gets the value of AllowStandardUserEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyAllowStandardUserEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStandardUserEncryption")
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

// SetAllowWarningForOtherDiskEncryption sets the value of AllowWarningForOtherDiskEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyAllowWarningForOtherDiskEncryption(value int32) (err error) {
	return instance.SetProperty("AllowWarningForOtherDiskEncryption", (value))
}

// GetAllowWarningForOtherDiskEncryption gets the value of AllowWarningForOtherDiskEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyAllowWarningForOtherDiskEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWarningForOtherDiskEncryption")
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

// SetConfigureRecoveryPasswordRotation sets the value of ConfigureRecoveryPasswordRotation for the instance
func (instance *MDM_BitLocker) SetPropertyConfigureRecoveryPasswordRotation(value int32) (err error) {
	return instance.SetProperty("ConfigureRecoveryPasswordRotation", (value))
}

// GetConfigureRecoveryPasswordRotation gets the value of ConfigureRecoveryPasswordRotation for the instance
func (instance *MDM_BitLocker) GetPropertyConfigureRecoveryPasswordRotation() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureRecoveryPasswordRotation")
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

// SetEncryptionMethodByDriveType sets the value of EncryptionMethodByDriveType for the instance
func (instance *MDM_BitLocker) SetPropertyEncryptionMethodByDriveType(value string) (err error) {
	return instance.SetProperty("EncryptionMethodByDriveType", (value))
}

// GetEncryptionMethodByDriveType gets the value of EncryptionMethodByDriveType for the instance
func (instance *MDM_BitLocker) GetPropertyEncryptionMethodByDriveType() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionMethodByDriveType")
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

// SetFixedDrivesEncryptionType sets the value of FixedDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) SetPropertyFixedDrivesEncryptionType(value string) (err error) {
	return instance.SetProperty("FixedDrivesEncryptionType", (value))
}

// GetFixedDrivesEncryptionType gets the value of FixedDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) GetPropertyFixedDrivesEncryptionType() (value string, err error) {
	retValue, err := instance.GetProperty("FixedDrivesEncryptionType")
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

// SetFixedDrivesRecoveryOptions sets the value of FixedDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) SetPropertyFixedDrivesRecoveryOptions(value string) (err error) {
	return instance.SetProperty("FixedDrivesRecoveryOptions", (value))
}

// GetFixedDrivesRecoveryOptions gets the value of FixedDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) GetPropertyFixedDrivesRecoveryOptions() (value string, err error) {
	retValue, err := instance.GetProperty("FixedDrivesRecoveryOptions")
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

// SetFixedDrivesRequireEncryption sets the value of FixedDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyFixedDrivesRequireEncryption(value string) (err error) {
	return instance.SetProperty("FixedDrivesRequireEncryption", (value))
}

// GetFixedDrivesRequireEncryption gets the value of FixedDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyFixedDrivesRequireEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("FixedDrivesRequireEncryption")
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

// SetIdentificationField sets the value of IdentificationField for the instance
func (instance *MDM_BitLocker) SetPropertyIdentificationField(value string) (err error) {
	return instance.SetProperty("IdentificationField", (value))
}

// GetIdentificationField gets the value of IdentificationField for the instance
func (instance *MDM_BitLocker) GetPropertyIdentificationField() (value string, err error) {
	retValue, err := instance.GetProperty("IdentificationField")
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
func (instance *MDM_BitLocker) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_BitLocker) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_BitLocker) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_BitLocker) GetPropertyParentID() (value string, err error) {
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

// SetRemovableDrivesConfigureBDE sets the value of RemovableDrivesConfigureBDE for the instance
func (instance *MDM_BitLocker) SetPropertyRemovableDrivesConfigureBDE(value string) (err error) {
	return instance.SetProperty("RemovableDrivesConfigureBDE", (value))
}

// GetRemovableDrivesConfigureBDE gets the value of RemovableDrivesConfigureBDE for the instance
func (instance *MDM_BitLocker) GetPropertyRemovableDrivesConfigureBDE() (value string, err error) {
	retValue, err := instance.GetProperty("RemovableDrivesConfigureBDE")
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

// SetRemovableDrivesEncryptionType sets the value of RemovableDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) SetPropertyRemovableDrivesEncryptionType(value string) (err error) {
	return instance.SetProperty("RemovableDrivesEncryptionType", (value))
}

// GetRemovableDrivesEncryptionType gets the value of RemovableDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) GetPropertyRemovableDrivesEncryptionType() (value string, err error) {
	retValue, err := instance.GetProperty("RemovableDrivesEncryptionType")
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

// SetRemovableDrivesExcludedFromEncryption sets the value of RemovableDrivesExcludedFromEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRemovableDrivesExcludedFromEncryption(value string) (err error) {
	return instance.SetProperty("RemovableDrivesExcludedFromEncryption", (value))
}

// GetRemovableDrivesExcludedFromEncryption gets the value of RemovableDrivesExcludedFromEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRemovableDrivesExcludedFromEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("RemovableDrivesExcludedFromEncryption")
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

// SetRemovableDrivesRequireEncryption sets the value of RemovableDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRemovableDrivesRequireEncryption(value string) (err error) {
	return instance.SetProperty("RemovableDrivesRequireEncryption", (value))
}

// GetRemovableDrivesRequireEncryption gets the value of RemovableDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRemovableDrivesRequireEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("RemovableDrivesRequireEncryption")
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

// SetRequireDeviceEncryption sets the value of RequireDeviceEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRequireDeviceEncryption(value int32) (err error) {
	return instance.SetProperty("RequireDeviceEncryption", (value))
}

// GetRequireDeviceEncryption gets the value of RequireDeviceEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRequireDeviceEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireDeviceEncryption")
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

// SetRequireStorageCardEncryption sets the value of RequireStorageCardEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRequireStorageCardEncryption(value int32) (err error) {
	return instance.SetProperty("RequireStorageCardEncryption", (value))
}

// GetRequireStorageCardEncryption gets the value of RequireStorageCardEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRequireStorageCardEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireStorageCardEncryption")
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

// SetSystemDrivesDisallowStandardUsersCanChangePIN sets the value of SystemDrivesDisallowStandardUsersCanChangePIN for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesDisallowStandardUsersCanChangePIN(value string) (err error) {
	return instance.SetProperty("SystemDrivesDisallowStandardUsersCanChangePIN", (value))
}

// GetSystemDrivesDisallowStandardUsersCanChangePIN gets the value of SystemDrivesDisallowStandardUsersCanChangePIN for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesDisallowStandardUsersCanChangePIN() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesDisallowStandardUsersCanChangePIN")
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

// SetSystemDrivesEnablePrebootInputProtectorsOnSlates sets the value of SystemDrivesEnablePrebootInputProtectorsOnSlates for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesEnablePrebootInputProtectorsOnSlates(value string) (err error) {
	return instance.SetProperty("SystemDrivesEnablePrebootInputProtectorsOnSlates", (value))
}

// GetSystemDrivesEnablePrebootInputProtectorsOnSlates gets the value of SystemDrivesEnablePrebootInputProtectorsOnSlates for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesEnablePrebootInputProtectorsOnSlates() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesEnablePrebootInputProtectorsOnSlates")
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

// SetSystemDrivesEnablePreBootPinExceptionOnDECapableDevice sets the value of SystemDrivesEnablePreBootPinExceptionOnDECapableDevice for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesEnablePreBootPinExceptionOnDECapableDevice(value string) (err error) {
	return instance.SetProperty("SystemDrivesEnablePreBootPinExceptionOnDECapableDevice", (value))
}

// GetSystemDrivesEnablePreBootPinExceptionOnDECapableDevice gets the value of SystemDrivesEnablePreBootPinExceptionOnDECapableDevice for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesEnablePreBootPinExceptionOnDECapableDevice() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesEnablePreBootPinExceptionOnDECapableDevice")
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

// SetSystemDrivesEncryptionType sets the value of SystemDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesEncryptionType(value string) (err error) {
	return instance.SetProperty("SystemDrivesEncryptionType", (value))
}

// GetSystemDrivesEncryptionType gets the value of SystemDrivesEncryptionType for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesEncryptionType() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesEncryptionType")
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

// SetSystemDrivesEnhancedPIN sets the value of SystemDrivesEnhancedPIN for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesEnhancedPIN(value string) (err error) {
	return instance.SetProperty("SystemDrivesEnhancedPIN", (value))
}

// GetSystemDrivesEnhancedPIN gets the value of SystemDrivesEnhancedPIN for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesEnhancedPIN() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesEnhancedPIN")
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

// SetSystemDrivesMinimumPINLength sets the value of SystemDrivesMinimumPINLength for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesMinimumPINLength(value string) (err error) {
	return instance.SetProperty("SystemDrivesMinimumPINLength", (value))
}

// GetSystemDrivesMinimumPINLength gets the value of SystemDrivesMinimumPINLength for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesMinimumPINLength() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesMinimumPINLength")
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

// SetSystemDrivesRecoveryMessage sets the value of SystemDrivesRecoveryMessage for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRecoveryMessage(value string) (err error) {
	return instance.SetProperty("SystemDrivesRecoveryMessage", (value))
}

// GetSystemDrivesRecoveryMessage gets the value of SystemDrivesRecoveryMessage for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRecoveryMessage() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRecoveryMessage")
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

// SetSystemDrivesRecoveryOptions sets the value of SystemDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRecoveryOptions(value string) (err error) {
	return instance.SetProperty("SystemDrivesRecoveryOptions", (value))
}

// GetSystemDrivesRecoveryOptions gets the value of SystemDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRecoveryOptions() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRecoveryOptions")
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

// SetSystemDrivesRequireStartupAuthentication sets the value of SystemDrivesRequireStartupAuthentication for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRequireStartupAuthentication(value string) (err error) {
	return instance.SetProperty("SystemDrivesRequireStartupAuthentication", (value))
}

// GetSystemDrivesRequireStartupAuthentication gets the value of SystemDrivesRequireStartupAuthentication for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRequireStartupAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRequireStartupAuthentication")
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

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_BitLocker) RotateRecoveryPasswordsMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RotateRecoveryPasswordsMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
