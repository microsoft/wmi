// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterPowerManagement_Offload_Arp struct
type MSFT_NetAdapterPowerManagement_Offload_Arp struct {
	*MSFT_NetAdapterPowerManagement_Offload

	//
	HostIPv4Address string

	//
	MACAddress string

	//
	RemoteIPv4Address string
}

func NewMSFT_NetAdapterPowerManagement_Offload_ArpEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_Offload_Arp, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_Arp{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_Offload_ArpEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_Offload_Arp, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_Arp{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

// SetHostIPv4Address sets the value of HostIPv4Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) SetPropertyHostIPv4Address(value string) (err error) {
	return instance.SetProperty("HostIPv4Address", value)
}

// GetHostIPv4Address gets the value of HostIPv4Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) GetPropertyHostIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("HostIPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMACAddress sets the value of MACAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", value)
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) GetPropertyMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MACAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteIPv4Address sets the value of RemoteIPv4Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) SetPropertyRemoteIPv4Address(value string) (err error) {
	return instance.SetProperty("RemoteIPv4Address", value)
}

// GetRemoteIPv4Address gets the value of RemoteIPv4Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_Arp) GetPropertyRemoteIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteIPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
