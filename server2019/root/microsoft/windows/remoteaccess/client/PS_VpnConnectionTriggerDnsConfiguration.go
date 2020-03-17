// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnConnectionTriggerDnsConfiguration struct
type PS_VpnConnectionTriggerDnsConfiguration struct {
	cim.WmiInstance

	//
	DnsIPAddress []string

	//
	DnsSuffix string
}

// SetDnsIPAddress sets the value of DnsIPAddress for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) SetPropertyDnsIPAddress(value []string) (err error) {
	return instance.SetProperty("DnsIPAddress", value)
}

// GetDnsIPAddress gets the value of DnsIPAddress for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) GetPropertyDnsIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DnsIPAddress" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerDnsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerDnsConfiguration) Add( /* IN */ ConnectionName string,
	/* IN */ DnsSuffix string,
	/* IN */ DnsIPAddress []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerDnsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ConnectionName, DnsSuffix, DnsIPAddress, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DnsIPAddress" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="DnsSuffixSearchList" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerDnsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerDnsConfiguration) Set( /* IN */ ConnectionName string,
	/* IN */ DnsSuffix string,
	/* IN */ DnsIPAddress []string,
	/* IN */ DnsSuffixSearchList []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerDnsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ConnectionName, DnsSuffix, DnsIPAddress, DnsSuffixSearchList, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DnsSuffix" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerDnsConfiguration []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerDnsConfiguration) Remove( /* IN */ ConnectionName string,
	/* IN */ DnsSuffix []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput []VpnConnectionTriggerDnsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ConnectionName, DnsSuffix, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
