// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_BitLocker struct
type MDM_BitLocker struct {
	cim.WmiInstance

	//
	AllowStandardUserEncryption int32

	//
	AllowWarningForOtherDiskEncryption int32

	//
	ConfigureRecoveryPasswordRotation int32

	//
	EncryptionMethodByDriveType string

	//
	FixedDrivesRecoveryOptions string

	//
	FixedDrivesRequireEncryption string

	//
	InstanceID string

	//
	ParentID string

	//
	RemovableDrivesRequireEncryption string

	//
	RequireDeviceEncryption int32

	//
	RequireStorageCardEncryption int32

	//
	SystemDrivesMinimumPINLength string

	//
	SystemDrivesRecoveryMessage string

	//
	SystemDrivesRecoveryOptions string

	//
	SystemDrivesRequireStartupAuthentication string
}

// SetAllowStandardUserEncryption sets the value of AllowStandardUserEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyAllowStandardUserEncryption(value int32) (err error) {
	return instance.SetProperty("AllowStandardUserEncryption", value)
}

// GetAllowStandardUserEncryption gets the value of AllowStandardUserEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyAllowStandardUserEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStandardUserEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWarningForOtherDiskEncryption sets the value of AllowWarningForOtherDiskEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyAllowWarningForOtherDiskEncryption(value int32) (err error) {
	return instance.SetProperty("AllowWarningForOtherDiskEncryption", value)
}

// GetAllowWarningForOtherDiskEncryption gets the value of AllowWarningForOtherDiskEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyAllowWarningForOtherDiskEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWarningForOtherDiskEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureRecoveryPasswordRotation sets the value of ConfigureRecoveryPasswordRotation for the instance
func (instance *MDM_BitLocker) SetPropertyConfigureRecoveryPasswordRotation(value int32) (err error) {
	return instance.SetProperty("ConfigureRecoveryPasswordRotation", value)
}

// GetConfigureRecoveryPasswordRotation gets the value of ConfigureRecoveryPasswordRotation for the instance
func (instance *MDM_BitLocker) GetPropertyConfigureRecoveryPasswordRotation() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureRecoveryPasswordRotation")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionMethodByDriveType sets the value of EncryptionMethodByDriveType for the instance
func (instance *MDM_BitLocker) SetPropertyEncryptionMethodByDriveType(value string) (err error) {
	return instance.SetProperty("EncryptionMethodByDriveType", value)
}

// GetEncryptionMethodByDriveType gets the value of EncryptionMethodByDriveType for the instance
func (instance *MDM_BitLocker) GetPropertyEncryptionMethodByDriveType() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionMethodByDriveType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFixedDrivesRecoveryOptions sets the value of FixedDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) SetPropertyFixedDrivesRecoveryOptions(value string) (err error) {
	return instance.SetProperty("FixedDrivesRecoveryOptions", value)
}

// GetFixedDrivesRecoveryOptions gets the value of FixedDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) GetPropertyFixedDrivesRecoveryOptions() (value string, err error) {
	retValue, err := instance.GetProperty("FixedDrivesRecoveryOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFixedDrivesRequireEncryption sets the value of FixedDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyFixedDrivesRequireEncryption(value string) (err error) {
	return instance.SetProperty("FixedDrivesRequireEncryption", value)
}

// GetFixedDrivesRequireEncryption gets the value of FixedDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyFixedDrivesRequireEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("FixedDrivesRequireEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_BitLocker) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_BitLocker) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_BitLocker) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_BitLocker) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovableDrivesRequireEncryption sets the value of RemovableDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRemovableDrivesRequireEncryption(value string) (err error) {
	return instance.SetProperty("RemovableDrivesRequireEncryption", value)
}

// GetRemovableDrivesRequireEncryption gets the value of RemovableDrivesRequireEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRemovableDrivesRequireEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("RemovableDrivesRequireEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireDeviceEncryption sets the value of RequireDeviceEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRequireDeviceEncryption(value int32) (err error) {
	return instance.SetProperty("RequireDeviceEncryption", value)
}

// GetRequireDeviceEncryption gets the value of RequireDeviceEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRequireDeviceEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireDeviceEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireStorageCardEncryption sets the value of RequireStorageCardEncryption for the instance
func (instance *MDM_BitLocker) SetPropertyRequireStorageCardEncryption(value int32) (err error) {
	return instance.SetProperty("RequireStorageCardEncryption", value)
}

// GetRequireStorageCardEncryption gets the value of RequireStorageCardEncryption for the instance
func (instance *MDM_BitLocker) GetPropertyRequireStorageCardEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireStorageCardEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemDrivesMinimumPINLength sets the value of SystemDrivesMinimumPINLength for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesMinimumPINLength(value string) (err error) {
	return instance.SetProperty("SystemDrivesMinimumPINLength", value)
}

// GetSystemDrivesMinimumPINLength gets the value of SystemDrivesMinimumPINLength for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesMinimumPINLength() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesMinimumPINLength")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemDrivesRecoveryMessage sets the value of SystemDrivesRecoveryMessage for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRecoveryMessage(value string) (err error) {
	return instance.SetProperty("SystemDrivesRecoveryMessage", value)
}

// GetSystemDrivesRecoveryMessage gets the value of SystemDrivesRecoveryMessage for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRecoveryMessage() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRecoveryMessage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemDrivesRecoveryOptions sets the value of SystemDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRecoveryOptions(value string) (err error) {
	return instance.SetProperty("SystemDrivesRecoveryOptions", value)
}

// GetSystemDrivesRecoveryOptions gets the value of SystemDrivesRecoveryOptions for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRecoveryOptions() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRecoveryOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemDrivesRequireStartupAuthentication sets the value of SystemDrivesRequireStartupAuthentication for the instance
func (instance *MDM_BitLocker) SetPropertySystemDrivesRequireStartupAuthentication(value string) (err error) {
	return instance.SetProperty("SystemDrivesRequireStartupAuthentication", value)
}

// GetSystemDrivesRequireStartupAuthentication gets the value of SystemDrivesRequireStartupAuthentication for the instance
func (instance *MDM_BitLocker) GetPropertySystemDrivesRequireStartupAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDrivesRequireStartupAuthentication")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
