// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSNetworkAdapterListSetting struct
type Win32_TSNetworkAdapterListSetting struct {
	*Win32_TerminalSetting

	//
	NetworkAdapterID string

	//
	NetworkAdapterIP string
}

func NewWin32_TSNetworkAdapterListSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSNetworkAdapterListSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSNetworkAdapterListSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSNetworkAdapterListSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSNetworkAdapterListSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSNetworkAdapterListSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetNetworkAdapterID sets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterListSetting) SetPropertyNetworkAdapterID(value string) (err error) {
	return instance.SetProperty("NetworkAdapterID", (value))
}

// GetNetworkAdapterID gets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterListSetting) GetPropertyNetworkAdapterID() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterID")
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

// SetNetworkAdapterIP sets the value of NetworkAdapterIP for the instance
func (instance *Win32_TSNetworkAdapterListSetting) SetPropertyNetworkAdapterIP(value string) (err error) {
	return instance.SetProperty("NetworkAdapterIP", (value))
}

// GetNetworkAdapterIP gets the value of NetworkAdapterIP for the instance
func (instance *Win32_TSNetworkAdapterListSetting) GetPropertyNetworkAdapterIP() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterIP")
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
