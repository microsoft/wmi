// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetNatExternalAddress struct
type MSFT_NetNatExternalAddress struct {
	MSFT_NetSettingData

	//
	Active uint8

	//
	ExternalAddressID uint32

	//
	IPAddress string

	//
	NatName string

	//
	PortEnd uint16

	//
	PortStart uint16
}

// SetActive sets the value of Active for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyActive(value uint8) (err error) {
	return instance.SetProperty("Active", value)
}

// GetActive gets the value of Active for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyActive() (value uint8, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExternalAddressID sets the value of ExternalAddressID for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyExternalAddressID(value uint32) (err error) {
	return instance.SetProperty("ExternalAddressID", value)
}

// GetExternalAddressID gets the value of ExternalAddressID for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyExternalAddressID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExternalAddressID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNatName sets the value of NatName for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyNatName(value string) (err error) {
	return instance.SetProperty("NatName", value)
}

// GetNatName gets the value of NatName for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyNatName() (value string, err error) {
	retValue, err := instance.GetProperty("NatName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortEnd sets the value of PortEnd for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyPortEnd(value uint16) (err error) {
	return instance.SetProperty("PortEnd", value)
}

// GetPortEnd gets the value of PortEnd for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyPortEnd() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortStart sets the value of PortStart for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyPortStart(value uint16) (err error) {
	return instance.SetProperty("PortStart", value)
}

// GetPortStart gets the value of PortStart for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyPortStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortStart")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
