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

// MDM_VPNv2_CryptographySuite03 struct
type MDM_VPNv2_CryptographySuite03 struct {
	cim.WmiInstance

	//
	AuthenticationTransformConstants string

	//
	CipherTransformConstants string

	//
	DHGroup string

	//
	EncryptionMethod string

	//
	InstanceID string

	//
	IntegrityCheckMethod string

	//
	ParentID string

	//
	PfsGroup string
}

// SetAuthenticationTransformConstants sets the value of AuthenticationTransformConstants for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyAuthenticationTransformConstants(value string) (err error) {
	return instance.SetProperty("AuthenticationTransformConstants", value)
}

// GetAuthenticationTransformConstants gets the value of AuthenticationTransformConstants for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyAuthenticationTransformConstants() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationTransformConstants")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCipherTransformConstants sets the value of CipherTransformConstants for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyCipherTransformConstants(value string) (err error) {
	return instance.SetProperty("CipherTransformConstants", value)
}

// GetCipherTransformConstants gets the value of CipherTransformConstants for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyCipherTransformConstants() (value string, err error) {
	retValue, err := instance.GetProperty("CipherTransformConstants")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDHGroup sets the value of DHGroup for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyDHGroup(value string) (err error) {
	return instance.SetProperty("DHGroup", value)
}

// GetDHGroup gets the value of DHGroup for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyDHGroup() (value string, err error) {
	retValue, err := instance.GetProperty("DHGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyEncryptionMethod(value string) (err error) {
	return instance.SetProperty("EncryptionMethod", value)
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyEncryptionMethod() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionMethod")
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
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyInstanceID() (value string, err error) {
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

// SetIntegrityCheckMethod sets the value of IntegrityCheckMethod for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyIntegrityCheckMethod(value string) (err error) {
	return instance.SetProperty("IntegrityCheckMethod", value)
}

// GetIntegrityCheckMethod gets the value of IntegrityCheckMethod for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyIntegrityCheckMethod() (value string, err error) {
	retValue, err := instance.GetProperty("IntegrityCheckMethod")
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
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyParentID() (value string, err error) {
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

// SetPfsGroup sets the value of PfsGroup for the instance
func (instance *MDM_VPNv2_CryptographySuite03) SetPropertyPfsGroup(value string) (err error) {
	return instance.SetProperty("PfsGroup", value)
}

// GetPfsGroup gets the value of PfsGroup for the instance
func (instance *MDM_VPNv2_CryptographySuite03) GetPropertyPfsGroup() (value string, err error) {
	retValue, err := instance.GetProperty("PfsGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
