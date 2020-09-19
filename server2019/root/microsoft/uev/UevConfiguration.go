// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// UevConfiguration struct
type UevConfiguration struct {
	*cim.WmiInstance
}

func NewUevConfigurationEx1(instance *cim.WmiInstance) (newInstance *UevConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &UevConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewUevConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UevConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UevConfiguration{
		WmiInstance: tmp,
	}
	return
}

// Enable UEV.
func (instance *UevConfiguration) Enable() (err error) {
	_, err = instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	return

}

// Disable UEV.
func (instance *UevConfiguration) Disable() (err error) {
	_, err = instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	return

}

// Check if UEV is enabled.

// <param name="ReturnValue" type="bool "></param>
func (instance *UevConfiguration) IsEnabled() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsEnabled")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

// Check if UEV is in reboot required state.

// <param name="ReturnValue" type="bool "></param>
func (instance *UevConfiguration) IsRebootRequired() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsRebootRequired")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
