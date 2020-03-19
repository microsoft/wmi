// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SecuritySetting struct
type Win32_SecuritySetting struct {
	*CIM_Setting

	//
	ControlFlags uint32
}

func NewWin32_SecuritySettingEx1(instance *cim.WmiInstance) (newInstance *Win32_SecuritySetting, err error) {
	tmp, err := NewCIM_SettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_SecuritySetting{
		CIM_Setting: tmp,
	}
	return
}

func NewWin32_SecuritySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SecuritySetting, err error) {
	tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SecuritySetting{
		CIM_Setting: tmp,
	}
	return
}

// SetControlFlags sets the value of ControlFlags for the instance
func (instance *Win32_SecuritySetting) SetPropertyControlFlags(value uint32) (err error) {
	return instance.SetProperty("ControlFlags", value)
}

// GetControlFlags gets the value of ControlFlags for the instance
func (instance *Win32_SecuritySetting) GetPropertyControlFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ControlFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SecuritySetting) GetSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SecuritySetting) SetSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
