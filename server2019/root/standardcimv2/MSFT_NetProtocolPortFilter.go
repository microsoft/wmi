// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetProtocolPortFilter struct
type MSFT_NetProtocolPortFilter struct {
	CIM_FilterEntryBase

	//
	DynamicTransport uint32

	//
	IcmpType []string

	//
	LocalPort []string

	//
	Protocol string

	//
	RemotePort []string
}

// SetDynamicTransport sets the value of DynamicTransport for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyDynamicTransport(value uint32) (err error) {
	return instance.SetProperty("DynamicTransport", value)
}

// GetDynamicTransport gets the value of DynamicTransport for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyDynamicTransport() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicTransport")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIcmpType sets the value of IcmpType for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyIcmpType(value []string) (err error) {
	return instance.SetProperty("IcmpType", value)
}

// GetIcmpType gets the value of IcmpType for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyIcmpType() (value []string, err error) {
	retValue, err := instance.GetProperty("IcmpType")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPort sets the value of LocalPort for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyLocalPort(value []string) (err error) {
	return instance.SetProperty("LocalPort", value)
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyLocalPort() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalPort")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetProtocolPortFilter) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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
func (instance *MSFT_NetProtocolPortFilter) SetPropertyRemotePort(value []string) (err error) {
	return instance.SetProperty("RemotePort", value)
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *MSFT_NetProtocolPortFilter) GetPropertyRemotePort() (value []string, err error) {
	retValue, err := instance.GetProperty("RemotePort")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
