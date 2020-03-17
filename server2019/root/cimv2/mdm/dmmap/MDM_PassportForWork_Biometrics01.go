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

// MDM_PassportForWork_Biometrics01 struct
type MDM_PassportForWork_Biometrics01 struct {
	cim.WmiInstance

	//
	FacialFeaturesUseEnhancedAntiSpoofing bool

	//
	InstanceID string

	//
	ParentID string

	//
	UseBiometrics bool
}

// SetFacialFeaturesUseEnhancedAntiSpoofing sets the value of FacialFeaturesUseEnhancedAntiSpoofing for the instance
func (instance *MDM_PassportForWork_Biometrics01) SetPropertyFacialFeaturesUseEnhancedAntiSpoofing(value bool) (err error) {
	return instance.SetProperty("FacialFeaturesUseEnhancedAntiSpoofing", value)
}

// GetFacialFeaturesUseEnhancedAntiSpoofing gets the value of FacialFeaturesUseEnhancedAntiSpoofing for the instance
func (instance *MDM_PassportForWork_Biometrics01) GetPropertyFacialFeaturesUseEnhancedAntiSpoofing() (value bool, err error) {
	retValue, err := instance.GetProperty("FacialFeaturesUseEnhancedAntiSpoofing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_Biometrics01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_Biometrics01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_PassportForWork_Biometrics01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork_Biometrics01) GetPropertyParentID() (value string, err error) {
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

// SetUseBiometrics sets the value of UseBiometrics for the instance
func (instance *MDM_PassportForWork_Biometrics01) SetPropertyUseBiometrics(value bool) (err error) {
	return instance.SetProperty("UseBiometrics", value)
}

// GetUseBiometrics gets the value of UseBiometrics for the instance
func (instance *MDM_PassportForWork_Biometrics01) GetPropertyUseBiometrics() (value bool, err error) {
	retValue, err := instance.GetProperty("UseBiometrics")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
