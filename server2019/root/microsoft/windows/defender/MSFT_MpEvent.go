// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpEvent struct
type MSFT_MpEvent struct {
	cim.WmiInstance

	//
	AdditionalData uint32

	//
	CategoryDiscriminant uint32

	//
	ComputerNotificationsValue uint32

	//
	NotificationTime string

	//
	ScanNotificationsValue uint32

	//
	SignatureNotificationsValue uint32

	//
	ThreatNotificationsValue uint32
}

// SetAdditionalData sets the value of AdditionalData for the instance
func (instance *MSFT_MpEvent) SetPropertyAdditionalData(value uint32) (err error) {
	return instance.SetProperty("AdditionalData", value)
}

// GetAdditionalData gets the value of AdditionalData for the instance
func (instance *MSFT_MpEvent) GetPropertyAdditionalData() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdditionalData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCategoryDiscriminant sets the value of CategoryDiscriminant for the instance
func (instance *MSFT_MpEvent) SetPropertyCategoryDiscriminant(value uint32) (err error) {
	return instance.SetProperty("CategoryDiscriminant", value)
}

// GetCategoryDiscriminant gets the value of CategoryDiscriminant for the instance
func (instance *MSFT_MpEvent) GetPropertyCategoryDiscriminant() (value uint32, err error) {
	retValue, err := instance.GetProperty("CategoryDiscriminant")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerNotificationsValue sets the value of ComputerNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyComputerNotificationsValue(value uint32) (err error) {
	return instance.SetProperty("ComputerNotificationsValue", value)
}

// GetComputerNotificationsValue gets the value of ComputerNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyComputerNotificationsValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("ComputerNotificationsValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationTime sets the value of NotificationTime for the instance
func (instance *MSFT_MpEvent) SetPropertyNotificationTime(value string) (err error) {
	return instance.SetProperty("NotificationTime", value)
}

// GetNotificationTime gets the value of NotificationTime for the instance
func (instance *MSFT_MpEvent) GetPropertyNotificationTime() (value string, err error) {
	retValue, err := instance.GetProperty("NotificationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScanNotificationsValue sets the value of ScanNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyScanNotificationsValue(value uint32) (err error) {
	return instance.SetProperty("ScanNotificationsValue", value)
}

// GetScanNotificationsValue gets the value of ScanNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyScanNotificationsValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScanNotificationsValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignatureNotificationsValue sets the value of SignatureNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertySignatureNotificationsValue(value uint32) (err error) {
	return instance.SetProperty("SignatureNotificationsValue", value)
}

// GetSignatureNotificationsValue gets the value of SignatureNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertySignatureNotificationsValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureNotificationsValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatNotificationsValue sets the value of ThreatNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyThreatNotificationsValue(value uint32) (err error) {
	return instance.SetProperty("ThreatNotificationsValue", value)
}

// GetThreatNotificationsValue gets the value of ThreatNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyThreatNotificationsValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreatNotificationsValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
