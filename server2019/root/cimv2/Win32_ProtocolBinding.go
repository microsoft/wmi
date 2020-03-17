// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ProtocolBinding struct
type Win32_ProtocolBinding struct {
	cim.WmiInstance

	//
	Antecedent Win32_NetworkProtocol

	//
	Dependent Win32_SystemDriver

	//
	Device Win32_NetworkAdapter
}

// SetAntecedent sets the value of Antecedent for the instance
func (instance *Win32_ProtocolBinding) SetPropertyAntecedent(value Win32_NetworkProtocol) (err error) {
	return instance.SetProperty("Antecedent", value)
}

// GetAntecedent gets the value of Antecedent for the instance
func (instance *Win32_ProtocolBinding) GetPropertyAntecedent() (value Win32_NetworkProtocol, err error) {
	retValue, err := instance.GetProperty("Antecedent")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_NetworkProtocol)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependent sets the value of Dependent for the instance
func (instance *Win32_ProtocolBinding) SetPropertyDependent(value Win32_SystemDriver) (err error) {
	return instance.SetProperty("Dependent", value)
}

// GetDependent gets the value of Dependent for the instance
func (instance *Win32_ProtocolBinding) GetPropertyDependent() (value Win32_SystemDriver, err error) {
	retValue, err := instance.GetProperty("Dependent")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SystemDriver)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDevice sets the value of Device for the instance
func (instance *Win32_ProtocolBinding) SetPropertyDevice(value Win32_NetworkAdapter) (err error) {
	return instance.SetProperty("Device", value)
}

// GetDevice gets the value of Device for the instance
func (instance *Win32_ProtocolBinding) GetPropertyDevice() (value Win32_NetworkAdapter, err error) {
	retValue, err := instance.GetProperty("Device")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_NetworkAdapter)
	if !ok {
		// TODO: Set an error
	}
	return
}
