// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SystemEnclosure struct
type Win32_SystemEnclosure struct {
	CIM_Chassis

	//
	SecurityStatus uint16

	//
	SMBIOSAssetTag string
}

// SetSecurityStatus sets the value of SecurityStatus for the instance
func (instance *Win32_SystemEnclosure) SetPropertySecurityStatus(value uint16) (err error) {
	return instance.SetProperty("SecurityStatus", value)
}

// GetSecurityStatus gets the value of SecurityStatus for the instance
func (instance *Win32_SystemEnclosure) GetPropertySecurityStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("SecurityStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSMBIOSAssetTag sets the value of SMBIOSAssetTag for the instance
func (instance *Win32_SystemEnclosure) SetPropertySMBIOSAssetTag(value string) (err error) {
	return instance.SetProperty("SMBIOSAssetTag", value)
}

// GetSMBIOSAssetTag gets the value of SMBIOSAssetTag for the instance
func (instance *Win32_SystemEnclosure) GetPropertySMBIOSAssetTag() (value string, err error) {
	retValue, err := instance.GetProperty("SMBIOSAssetTag")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
