// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbServerNetworkInterface struct
type MSFT_SmbServerNetworkInterface struct {
	cim.WmiInstance

	//
	FriendlyName string

	//
	InterfaceIndex uint32

	//
	IpAddress string

	//
	LinkSpeed uint64

	//
	RdmaCapable bool

	//
	RssCapable bool

	//
	ScopeName string
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyInterfaceIndex() (value uint32, err error) {
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

// SetIpAddress sets the value of IpAddress for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyIpAddress(value string) (err error) {
	return instance.SetProperty("IpAddress", value)
}

// GetIpAddress gets the value of IpAddress for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IpAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkSpeed sets the value of LinkSpeed for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("LinkSpeed", value)
}

// GetLinkSpeed gets the value of LinkSpeed for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("LinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaCapable sets the value of RdmaCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyRdmaCapable(value bool) (err error) {
	return instance.SetProperty("RdmaCapable", value)
}

// GetRdmaCapable gets the value of RdmaCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyRdmaCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("RdmaCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRssCapable sets the value of RssCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyRssCapable(value bool) (err error) {
	return instance.SetProperty("RssCapable", value)
}

// GetRssCapable gets the value of RssCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyRssCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("RssCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", value)
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
