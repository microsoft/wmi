// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetTCPConnection struct
type MSFT_NetTCPConnection struct {
	MSFT_NetTransportConnection

	//
	AppliedSetting uint8

	//
	OffloadState uint8

	//
	RemoteAddress string

	//
	RemotePort uint16

	//
	State uint8
}

// SetAppliedSetting sets the value of AppliedSetting for the instance
func (instance *MSFT_NetTCPConnection) SetPropertyAppliedSetting(value uint8) (err error) {
	return instance.SetProperty("AppliedSetting", value)
}

// GetAppliedSetting gets the value of AppliedSetting for the instance
func (instance *MSFT_NetTCPConnection) GetPropertyAppliedSetting() (value uint8, err error) {
	retValue, err := instance.GetProperty("AppliedSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOffloadState sets the value of OffloadState for the instance
func (instance *MSFT_NetTCPConnection) SetPropertyOffloadState(value uint8) (err error) {
	return instance.SetProperty("OffloadState", value)
}

// GetOffloadState gets the value of OffloadState for the instance
func (instance *MSFT_NetTCPConnection) GetPropertyOffloadState() (value uint8, err error) {
	retValue, err := instance.GetProperty("OffloadState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteAddress sets the value of RemoteAddress for the instance
func (instance *MSFT_NetTCPConnection) SetPropertyRemoteAddress(value string) (err error) {
	return instance.SetProperty("RemoteAddress", value)
}

// GetRemoteAddress gets the value of RemoteAddress for the instance
func (instance *MSFT_NetTCPConnection) GetPropertyRemoteAddress() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePort sets the value of RemotePort for the instance
func (instance *MSFT_NetTCPConnection) SetPropertyRemotePort(value uint16) (err error) {
	return instance.SetProperty("RemotePort", value)
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *MSFT_NetTCPConnection) GetPropertyRemotePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemotePort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetTCPConnection) SetPropertyState(value uint8) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetTCPConnection) GetPropertyState() (value uint8, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
