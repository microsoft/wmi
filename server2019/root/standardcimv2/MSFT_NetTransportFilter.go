// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetTransportFilter struct
type MSFT_NetTransportFilter struct {
	CIM_FilterEntryBase

	//
	DestinationPrefix string

	//
	LocalPortEnd uint16

	//
	LocalPortStart uint16

	//
	Protocol uint16

	//
	RemotePortEnd uint16

	//
	RemotePortStart uint16

	//
	SettingName string
}

// SetDestinationPrefix sets the value of DestinationPrefix for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyDestinationPrefix(value string) (err error) {
	return instance.SetProperty("DestinationPrefix", value)
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyDestinationPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPrefix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPortEnd sets the value of LocalPortEnd for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyLocalPortEnd(value uint16) (err error) {
	return instance.SetProperty("LocalPortEnd", value)
}

// GetLocalPortEnd gets the value of LocalPortEnd for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyLocalPortEnd() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalPortEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPortStart sets the value of LocalPortStart for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyLocalPortStart(value uint16) (err error) {
	return instance.SetProperty("LocalPortStart", value)
}

// GetLocalPortStart gets the value of LocalPortStart for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyLocalPortStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalPortStart")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyProtocol(value uint16) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyProtocol() (value uint16, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePortEnd sets the value of RemotePortEnd for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyRemotePortEnd(value uint16) (err error) {
	return instance.SetProperty("RemotePortEnd", value)
}

// GetRemotePortEnd gets the value of RemotePortEnd for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyRemotePortEnd() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemotePortEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePortStart sets the value of RemotePortStart for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyRemotePortStart(value uint16) (err error) {
	return instance.SetProperty("RemotePortStart", value)
}

// GetRemotePortStart gets the value of RemotePortStart for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyRemotePortStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemotePortStart")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingName sets the value of SettingName for the instance
func (instance *MSFT_NetTransportFilter) SetPropertySettingName(value string) (err error) {
	return instance.SetProperty("SettingName", value)
}

// GetSettingName gets the value of SettingName for the instance
func (instance *MSFT_NetTransportFilter) GetPropertySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("SettingName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
