// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn struct
type MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn struct {
	MSFT_NetAdapterPowerManagement_WakePattern

	//
	DestinationAddress string

	//
	DestinationPort uint16

	//
	SourceAddress string

	//
	SourcePort uint16
}

// SetDestinationAddress sets the value of DestinationAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertyDestinationAddress(value string) (err error) {
	return instance.SetProperty("DestinationAddress", value)
}

// GetDestinationAddress gets the value of DestinationAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertyDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationPort sets the value of DestinationPort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertyDestinationPort(value uint16) (err error) {
	return instance.SetProperty("DestinationPort", value)
}

// GetDestinationPort gets the value of DestinationPort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertyDestinationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("DestinationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceAddress sets the value of SourceAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertySourceAddress(value string) (err error) {
	return instance.SetProperty("SourceAddress", value)
}

// GetSourceAddress gets the value of SourceAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertySourceAddress() (value string, err error) {
	retValue, err := instance.GetProperty("SourceAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourcePort sets the value of SourcePort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertySourcePort(value uint16) (err error) {
	return instance.SetProperty("SourcePort", value)
}

// GetSourcePort gets the value of SourcePort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertySourcePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("SourcePort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
