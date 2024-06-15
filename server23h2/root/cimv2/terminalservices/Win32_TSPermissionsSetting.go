// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSPermissionsSetting struct
type Win32_TSPermissionsSetting struct {
	*Win32_TerminalSetting

	//
	DenyAdminPermissionForCustomization uint32

	//
	PolicySourceDenyAdminPermissionForCustomization uint32

	//
	StringSecurityDescriptor string
}

func NewWin32_TSPermissionsSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSPermissionsSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPermissionsSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSPermissionsSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSPermissionsSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPermissionsSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetDenyAdminPermissionForCustomization sets the value of DenyAdminPermissionForCustomization for the instance
func (instance *Win32_TSPermissionsSetting) SetPropertyDenyAdminPermissionForCustomization(value uint32) (err error) {
	return instance.SetProperty("DenyAdminPermissionForCustomization", (value))
}

// GetDenyAdminPermissionForCustomization gets the value of DenyAdminPermissionForCustomization for the instance
func (instance *Win32_TSPermissionsSetting) GetPropertyDenyAdminPermissionForCustomization() (value uint32, err error) {
	retValue, err := instance.GetProperty("DenyAdminPermissionForCustomization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceDenyAdminPermissionForCustomization sets the value of PolicySourceDenyAdminPermissionForCustomization for the instance
func (instance *Win32_TSPermissionsSetting) SetPropertyPolicySourceDenyAdminPermissionForCustomization(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDenyAdminPermissionForCustomization", (value))
}

// GetPolicySourceDenyAdminPermissionForCustomization gets the value of PolicySourceDenyAdminPermissionForCustomization for the instance
func (instance *Win32_TSPermissionsSetting) GetPropertyPolicySourceDenyAdminPermissionForCustomization() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDenyAdminPermissionForCustomization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetStringSecurityDescriptor sets the value of StringSecurityDescriptor for the instance
func (instance *Win32_TSPermissionsSetting) SetPropertyStringSecurityDescriptor(value string) (err error) {
	return instance.SetProperty("StringSecurityDescriptor", (value))
}

// GetStringSecurityDescriptor gets the value of StringSecurityDescriptor for the instance
func (instance *Win32_TSPermissionsSetting) GetPropertyStringSecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("StringSecurityDescriptor")
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

// <param name="AccountName" type="string "></param>
// <param name="PermissionPreSet" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSPermissionsSetting) AddAccount( /* IN */ AccountName string,
	/* IN */ PermissionPreSet uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddAccount", AccountName, PermissionPreSet)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSPermissionsSetting) RestoreDefaults() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestoreDefaults")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
