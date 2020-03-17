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

// CIM_SupportAccess struct
type CIM_SupportAccess struct {
	cim.WmiInstance

	//
	CommunicationInfo string

	//
	CommunicationMode uint16

	//
	Description string

	//
	Locale string

	//
	SupportAccessId string
}

// SetCommunicationInfo sets the value of CommunicationInfo for the instance
func (instance *CIM_SupportAccess) SetPropertyCommunicationInfo(value string) (err error) {
	return instance.SetProperty("CommunicationInfo", value)
}

// GetCommunicationInfo gets the value of CommunicationInfo for the instance
func (instance *CIM_SupportAccess) GetPropertyCommunicationInfo() (value string, err error) {
	retValue, err := instance.GetProperty("CommunicationInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommunicationMode sets the value of CommunicationMode for the instance
func (instance *CIM_SupportAccess) SetPropertyCommunicationMode(value uint16) (err error) {
	return instance.SetProperty("CommunicationMode", value)
}

// GetCommunicationMode gets the value of CommunicationMode for the instance
func (instance *CIM_SupportAccess) GetPropertyCommunicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("CommunicationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_SupportAccess) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_SupportAccess) GetPropertyDescription() (value string, err error) {
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

// SetLocale sets the value of Locale for the instance
func (instance *CIM_SupportAccess) SetPropertyLocale(value string) (err error) {
	return instance.SetProperty("Locale", value)
}

// GetLocale gets the value of Locale for the instance
func (instance *CIM_SupportAccess) GetPropertyLocale() (value string, err error) {
	retValue, err := instance.GetProperty("Locale")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportAccessId sets the value of SupportAccessId for the instance
func (instance *CIM_SupportAccess) SetPropertySupportAccessId(value string) (err error) {
	return instance.SetProperty("SupportAccessId", value)
}

// GetSupportAccessId gets the value of SupportAccessId for the instance
func (instance *CIM_SupportAccess) GetPropertySupportAccessId() (value string, err error) {
	retValue, err := instance.GetProperty("SupportAccessId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
