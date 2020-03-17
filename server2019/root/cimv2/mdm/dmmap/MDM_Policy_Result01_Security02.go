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

// MDM_Policy_Result01_Security02 struct
type MDM_Policy_Result01_Security02 struct {
	cim.WmiInstance

	//
	AllowAddProvisioningPackage int32

	//
	AllowRemoveProvisioningPackage int32

	//
	ClearTPMIfNotReady int32

	//
	ConfigureWindowsPasswords int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreventAutomaticDeviceEncryptionForAzureADJoinedDevices int32

	//
	RecoveryEnvironmentAuthentication int32

	//
	RequireDeviceEncryption int32

	//
	RequireProvisioningPackageSignature int32

	//
	RequireRetrieveHealthCertificateOnBoot int32
}

// SetAllowAddProvisioningPackage sets the value of AllowAddProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyAllowAddProvisioningPackage(value int32) (err error) {
	return instance.SetProperty("AllowAddProvisioningPackage", value)
}

// GetAllowAddProvisioningPackage gets the value of AllowAddProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyAllowAddProvisioningPackage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAddProvisioningPackage")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowRemoveProvisioningPackage sets the value of AllowRemoveProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyAllowRemoveProvisioningPackage(value int32) (err error) {
	return instance.SetProperty("AllowRemoveProvisioningPackage", value)
}

// GetAllowRemoveProvisioningPackage gets the value of AllowRemoveProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyAllowRemoveProvisioningPackage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowRemoveProvisioningPackage")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClearTPMIfNotReady sets the value of ClearTPMIfNotReady for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyClearTPMIfNotReady(value int32) (err error) {
	return instance.SetProperty("ClearTPMIfNotReady", value)
}

// GetClearTPMIfNotReady gets the value of ClearTPMIfNotReady for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyClearTPMIfNotReady() (value int32, err error) {
	retValue, err := instance.GetProperty("ClearTPMIfNotReady")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureWindowsPasswords sets the value of ConfigureWindowsPasswords for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyConfigureWindowsPasswords(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsPasswords", value)
}

// GetConfigureWindowsPasswords gets the value of ConfigureWindowsPasswords for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyConfigureWindowsPasswords() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsPasswords")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_Security02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyParentID() (value string, err error) {
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

// SetPreventAutomaticDeviceEncryptionForAzureADJoinedDevices sets the value of PreventAutomaticDeviceEncryptionForAzureADJoinedDevices for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyPreventAutomaticDeviceEncryptionForAzureADJoinedDevices(value int32) (err error) {
	return instance.SetProperty("PreventAutomaticDeviceEncryptionForAzureADJoinedDevices", value)
}

// GetPreventAutomaticDeviceEncryptionForAzureADJoinedDevices gets the value of PreventAutomaticDeviceEncryptionForAzureADJoinedDevices for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyPreventAutomaticDeviceEncryptionForAzureADJoinedDevices() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventAutomaticDeviceEncryptionForAzureADJoinedDevices")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryEnvironmentAuthentication sets the value of RecoveryEnvironmentAuthentication for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRecoveryEnvironmentAuthentication(value int32) (err error) {
	return instance.SetProperty("RecoveryEnvironmentAuthentication", value)
}

// GetRecoveryEnvironmentAuthentication gets the value of RecoveryEnvironmentAuthentication for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRecoveryEnvironmentAuthentication() (value int32, err error) {
	retValue, err := instance.GetProperty("RecoveryEnvironmentAuthentication")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireDeviceEncryption sets the value of RequireDeviceEncryption for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireDeviceEncryption(value int32) (err error) {
	return instance.SetProperty("RequireDeviceEncryption", value)
}

// GetRequireDeviceEncryption gets the value of RequireDeviceEncryption for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireDeviceEncryption() (value int32, err error) {
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

// SetRequireProvisioningPackageSignature sets the value of RequireProvisioningPackageSignature for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireProvisioningPackageSignature(value int32) (err error) {
	return instance.SetProperty("RequireProvisioningPackageSignature", value)
}

// GetRequireProvisioningPackageSignature gets the value of RequireProvisioningPackageSignature for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireProvisioningPackageSignature() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireProvisioningPackageSignature")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireRetrieveHealthCertificateOnBoot sets the value of RequireRetrieveHealthCertificateOnBoot for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireRetrieveHealthCertificateOnBoot(value int32) (err error) {
	return instance.SetProperty("RequireRetrieveHealthCertificateOnBoot", value)
}

// GetRequireRetrieveHealthCertificateOnBoot gets the value of RequireRetrieveHealthCertificateOnBoot for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireRetrieveHealthCertificateOnBoot() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireRetrieveHealthCertificateOnBoot")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
