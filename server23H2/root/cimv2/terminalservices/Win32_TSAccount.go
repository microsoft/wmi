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

// Win32_TSAccount struct
type Win32_TSAccount struct {
	*Win32_TerminalSetting

	//
	AccountName string

	//
	AuditFail uint32

	//
	AuditSuccess uint32

	//
	PermissionsAllowed uint32

	//
	PermissionsDenied uint32

	//
	SID string
}

func NewWin32_TSAccountEx1(instance *cim.WmiInstance) (newInstance *Win32_TSAccount, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSAccount{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSAccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSAccount, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSAccount{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *Win32_TSAccount) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", (value))
}

// GetAccountName gets the value of AccountName for the instance
func (instance *Win32_TSAccount) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
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

// SetAuditFail sets the value of AuditFail for the instance
func (instance *Win32_TSAccount) SetPropertyAuditFail(value uint32) (err error) {
	return instance.SetProperty("AuditFail", (value))
}

// GetAuditFail gets the value of AuditFail for the instance
func (instance *Win32_TSAccount) GetPropertyAuditFail() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuditFail")
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

// SetAuditSuccess sets the value of AuditSuccess for the instance
func (instance *Win32_TSAccount) SetPropertyAuditSuccess(value uint32) (err error) {
	return instance.SetProperty("AuditSuccess", (value))
}

// GetAuditSuccess gets the value of AuditSuccess for the instance
func (instance *Win32_TSAccount) GetPropertyAuditSuccess() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuditSuccess")
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

// SetPermissionsAllowed sets the value of PermissionsAllowed for the instance
func (instance *Win32_TSAccount) SetPropertyPermissionsAllowed(value uint32) (err error) {
	return instance.SetProperty("PermissionsAllowed", (value))
}

// GetPermissionsAllowed gets the value of PermissionsAllowed for the instance
func (instance *Win32_TSAccount) GetPropertyPermissionsAllowed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PermissionsAllowed")
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

// SetPermissionsDenied sets the value of PermissionsDenied for the instance
func (instance *Win32_TSAccount) SetPropertyPermissionsDenied(value uint32) (err error) {
	return instance.SetProperty("PermissionsDenied", (value))
}

// GetPermissionsDenied gets the value of PermissionsDenied for the instance
func (instance *Win32_TSAccount) GetPropertyPermissionsDenied() (value uint32, err error) {
	retValue, err := instance.GetProperty("PermissionsDenied")
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

// SetSID sets the value of SID for the instance
func (instance *Win32_TSAccount) SetPropertySID(value string) (err error) {
	return instance.SetProperty("SID", (value))
}

// GetSID gets the value of SID for the instance
func (instance *Win32_TSAccount) GetPropertySID() (value string, err error) {
	retValue, err := instance.GetProperty("SID")
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

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSAccount) Delete() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Delete")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Allow" type="bool "></param>
// <param name="PermissionMask" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSAccount) ModifyPermissions( /* IN */ PermissionMask uint32,
	/* IN */ Allow bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ModifyPermissions", PermissionMask, Allow)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="PermissionMask" type="uint32 "></param>
// <param name="Success" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSAccount) ModifyAuditPermissions( /* IN */ PermissionMask uint32,
	/* IN */ Success bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ModifyAuditPermissions", PermissionMask, Success)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
