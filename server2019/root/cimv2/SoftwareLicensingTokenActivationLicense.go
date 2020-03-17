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

// SoftwareLicensingTokenActivationLicense struct
type SoftwareLicensingTokenActivationLicense struct {
	cim.WmiInstance

	//
	AdditionalInfo string

	//
	AuthorizationStatus uint32

	//
	Description string

	//
	ExpirationDate string

	//
	ID string

	//
	ILID string

	//
	ILVID uint32
}

// SetAdditionalInfo sets the value of AdditionalInfo for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyAdditionalInfo(value string) (err error) {
	return instance.SetProperty("AdditionalInfo", value)
}

// GetAdditionalInfo gets the value of AdditionalInfo for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyAdditionalInfo() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthorizationStatus sets the value of AuthorizationStatus for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyAuthorizationStatus(value uint32) (err error) {
	return instance.SetProperty("AuthorizationStatus", value)
}

// GetAuthorizationStatus gets the value of AuthorizationStatus for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyAuthorizationStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthorizationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpirationDate sets the value of ExpirationDate for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyExpirationDate(value string) (err error) {
	return instance.SetProperty("ExpirationDate", value)
}

// GetExpirationDate gets the value of ExpirationDate for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyExpirationDate() (value string, err error) {
	retValue, err := instance.GetProperty("ExpirationDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetILID sets the value of ILID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyILID(value string) (err error) {
	return instance.SetProperty("ILID", value)
}

// GetILID gets the value of ILID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyILID() (value string, err error) {
	retValue, err := instance.GetProperty("ILID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetILVID sets the value of ILVID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) SetPropertyILVID(value uint32) (err error) {
	return instance.SetProperty("ILVID", value)
}

// GetILVID gets the value of ILVID for the instance
func (instance *SoftwareLicensingTokenActivationLicense) GetPropertyILVID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ILVID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SoftwareLicensingTokenActivationLicense) Uninstall() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Uninstall")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
