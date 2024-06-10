// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Client
//
// ////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PS_VpnConnectionTriggerDnsConfiguration struct
type PS_VpnConnectionTriggerDnsConfiguration struct {
	*cim.WmiInstance

	//
	DnsIPAddress []string

	//
	DnsSuffix string
}

func NewPS_VpnConnectionTriggerDnsConfigurationEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnectionTriggerDnsConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionTriggerDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnConnectionTriggerDnsConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnConnectionTriggerDnsConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionTriggerDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetDnsIPAddress sets the value of DnsIPAddress for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) SetPropertyDnsIPAddress(value []string) (err error) {
	return instance.SetProperty("DnsIPAddress", (value))
}

// GetDnsIPAddress gets the value of DnsIPAddress for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) GetPropertyDnsIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsIPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *PS_VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

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
