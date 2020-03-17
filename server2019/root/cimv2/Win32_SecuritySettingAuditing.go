// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SecuritySettingAuditing struct
type Win32_SecuritySettingAuditing struct {
	cim.WmiInstance

	//
	AuditedAccessMask uint32

	//
	GuidInheritedObjectType string

	//
	GuidObjectType string

	//
	Inheritance uint32

	//
	SecuritySetting Win32_SecuritySetting

	//
	Trustee Win32_SID

	//
	Type uint32
}

// SetAuditedAccessMask sets the value of AuditedAccessMask for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyAuditedAccessMask(value uint32) (err error) {
	return instance.SetProperty("AuditedAccessMask", value)
}

// GetAuditedAccessMask gets the value of AuditedAccessMask for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyAuditedAccessMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuditedAccessMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuidInheritedObjectType sets the value of GuidInheritedObjectType for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyGuidInheritedObjectType(value string) (err error) {
	return instance.SetProperty("GuidInheritedObjectType", value)
}

// GetGuidInheritedObjectType gets the value of GuidInheritedObjectType for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyGuidInheritedObjectType() (value string, err error) {
	retValue, err := instance.GetProperty("GuidInheritedObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuidObjectType sets the value of GuidObjectType for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyGuidObjectType(value string) (err error) {
	return instance.SetProperty("GuidObjectType", value)
}

// GetGuidObjectType gets the value of GuidObjectType for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyGuidObjectType() (value string, err error) {
	retValue, err := instance.GetProperty("GuidObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInheritance sets the value of Inheritance for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyInheritance(value uint32) (err error) {
	return instance.SetProperty("Inheritance", value)
}

// GetInheritance gets the value of Inheritance for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyInheritance() (value uint32, err error) {
	retValue, err := instance.GetProperty("Inheritance")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecuritySetting sets the value of SecuritySetting for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertySecuritySetting(value Win32_SecuritySetting) (err error) {
	return instance.SetProperty("SecuritySetting", value)
}

// GetSecuritySetting gets the value of SecuritySetting for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertySecuritySetting() (value Win32_SecuritySetting, err error) {
	retValue, err := instance.GetProperty("SecuritySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SecuritySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustee sets the value of Trustee for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyTrustee(value Win32_SID) (err error) {
	return instance.SetProperty("Trustee", value)
}

// GetTrustee gets the value of Trustee for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyTrustee() (value Win32_SID, err error) {
	retValue, err := instance.GetProperty("Trustee")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SID)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *Win32_SecuritySettingAuditing) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Win32_SecuritySettingAuditing) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
