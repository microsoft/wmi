// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_DNSClientServerAddress struct
type MSFT_DNSClientServerAddress struct {
	CIM_RemoteServiceAccessPoint

	// 747
	AddressFamily uint16

	// 656
	InterfaceAlias string

	// 655
	InterfaceIndex uint32

	// 746
	ServerAddresses []string
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", value)
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", value)
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerAddresses sets the value of ServerAddresses for the instance
func (instance *MSFT_DNSClientServerAddress) SetPropertyServerAddresses(value []string) (err error) {
	return instance.SetProperty("ServerAddresses", value)
}

// GetServerAddresses gets the value of ServerAddresses for the instance
func (instance *MSFT_DNSClientServerAddress) GetPropertyServerAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("ServerAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
