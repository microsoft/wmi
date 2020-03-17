// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetConnectionProfile struct
type MSFT_NetConnectionProfile struct {
	MSFT_NetSettingData

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	IPv4Connectivity uint32

	//
	IPv6Connectivity uint32

	//
	Name string

	//
	NetworkCategory uint32
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", value)
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyInterfaceAlias() (value string, err error) {
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
func (instance *MSFT_NetConnectionProfile) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyInterfaceIndex() (value uint32, err error) {
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

// SetIPv4Connectivity sets the value of IPv4Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyIPv4Connectivity(value uint32) (err error) {
	return instance.SetProperty("IPv4Connectivity", value)
}

// GetIPv4Connectivity gets the value of IPv4Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyIPv4Connectivity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4Connectivity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Connectivity sets the value of IPv6Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyIPv6Connectivity(value uint32) (err error) {
	return instance.SetProperty("IPv6Connectivity", value)
}

// GetIPv6Connectivity gets the value of IPv6Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyIPv6Connectivity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6Connectivity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkCategory sets the value of NetworkCategory for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyNetworkCategory(value uint32) (err error) {
	return instance.SetProperty("NetworkCategory", value)
}

// GetNetworkCategory gets the value of NetworkCategory for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyNetworkCategory() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkCategory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
