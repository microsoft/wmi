// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventPacketCaptureTarget struct
type MSFT_NetEventPacketCaptureTarget struct {
	CIM_LogicalElement

	//
	CaptureStatus uint32

	//
	Id string

	//
	ProviderName string
}

// SetCaptureStatus sets the value of CaptureStatus for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) SetPropertyCaptureStatus(value uint32) (err error) {
	return instance.SetProperty("CaptureStatus", value)
}

// GetCaptureStatus gets the value of CaptureStatus for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) GetPropertyCaptureStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("CaptureStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *MSFT_NetEventPacketCaptureTarget) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
