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

// MDM_EnterpriseDataProtection_Settings01 struct
type MDM_EnterpriseDataProtection_Settings01 struct {
	cim.WmiInstance

	//
	AllowAzureRMSForEDP int32

	//
	AllowUserDecryption int32

	//
	DataRecoveryCertificate string

	//
	EDPEnforcementLevel int32

	//
	EDPShowIcons int32

	//
	EnterpriseProtectedDomainNames string

	//
	InstanceID string

	//
	ParentID string

	//
	RevokeOnUnenroll int32

	//
	RMSTemplateIDForEDP string

	//
	SMBAutoEncryptedFileExtensions string
}

// SetAllowAzureRMSForEDP sets the value of AllowAzureRMSForEDP for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyAllowAzureRMSForEDP(value int32) (err error) {
	return instance.SetProperty("AllowAzureRMSForEDP", value)
}

// GetAllowAzureRMSForEDP gets the value of AllowAzureRMSForEDP for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyAllowAzureRMSForEDP() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAzureRMSForEDP")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowUserDecryption sets the value of AllowUserDecryption for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyAllowUserDecryption(value int32) (err error) {
	return instance.SetProperty("AllowUserDecryption", value)
}

// GetAllowUserDecryption gets the value of AllowUserDecryption for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyAllowUserDecryption() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUserDecryption")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataRecoveryCertificate sets the value of DataRecoveryCertificate for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyDataRecoveryCertificate(value string) (err error) {
	return instance.SetProperty("DataRecoveryCertificate", value)
}

// GetDataRecoveryCertificate gets the value of DataRecoveryCertificate for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyDataRecoveryCertificate() (value string, err error) {
	retValue, err := instance.GetProperty("DataRecoveryCertificate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEDPEnforcementLevel sets the value of EDPEnforcementLevel for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyEDPEnforcementLevel(value int32) (err error) {
	return instance.SetProperty("EDPEnforcementLevel", value)
}

// GetEDPEnforcementLevel gets the value of EDPEnforcementLevel for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyEDPEnforcementLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("EDPEnforcementLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEDPShowIcons sets the value of EDPShowIcons for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyEDPShowIcons(value int32) (err error) {
	return instance.SetProperty("EDPShowIcons", value)
}

// GetEDPShowIcons gets the value of EDPShowIcons for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyEDPShowIcons() (value int32, err error) {
	retValue, err := instance.GetProperty("EDPShowIcons")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseProtectedDomainNames sets the value of EnterpriseProtectedDomainNames for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyEnterpriseProtectedDomainNames(value string) (err error) {
	return instance.SetProperty("EnterpriseProtectedDomainNames", value)
}

// GetEnterpriseProtectedDomainNames gets the value of EnterpriseProtectedDomainNames for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyEnterpriseProtectedDomainNames() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseProtectedDomainNames")
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
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyParentID() (value string, err error) {
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

// SetRevokeOnUnenroll sets the value of RevokeOnUnenroll for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyRevokeOnUnenroll(value int32) (err error) {
	return instance.SetProperty("RevokeOnUnenroll", value)
}

// GetRevokeOnUnenroll gets the value of RevokeOnUnenroll for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyRevokeOnUnenroll() (value int32, err error) {
	retValue, err := instance.GetProperty("RevokeOnUnenroll")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRMSTemplateIDForEDP sets the value of RMSTemplateIDForEDP for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertyRMSTemplateIDForEDP(value string) (err error) {
	return instance.SetProperty("RMSTemplateIDForEDP", value)
}

// GetRMSTemplateIDForEDP gets the value of RMSTemplateIDForEDP for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertyRMSTemplateIDForEDP() (value string, err error) {
	retValue, err := instance.GetProperty("RMSTemplateIDForEDP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSMBAutoEncryptedFileExtensions sets the value of SMBAutoEncryptedFileExtensions for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) SetPropertySMBAutoEncryptedFileExtensions(value string) (err error) {
	return instance.SetProperty("SMBAutoEncryptedFileExtensions", value)
}

// GetSMBAutoEncryptedFileExtensions gets the value of SMBAutoEncryptedFileExtensions for the instance
func (instance *MDM_EnterpriseDataProtection_Settings01) GetPropertySMBAutoEncryptedFileExtensions() (value string, err error) {
	retValue, err := instance.GetProperty("SMBAutoEncryptedFileExtensions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
