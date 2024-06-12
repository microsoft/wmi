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

// Win32_TerminalServiceSettingError struct
type Win32_TerminalServiceSettingError struct {
	*__ExtendedStatus

	//
	TerminalServiceMode int32
}

func NewWin32_TerminalServiceSettingErrorEx1(instance *cim.WmiInstance) (newInstance *Win32_TerminalServiceSettingError, err error) {
	tmp, err := New__ExtendedStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalServiceSettingError{
		__ExtendedStatus: tmp,
	}
	return
}

func NewWin32_TerminalServiceSettingErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TerminalServiceSettingError, err error) {
	tmp, err := New__ExtendedStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalServiceSettingError{
		__ExtendedStatus: tmp,
	}
	return
}

// SetTerminalServiceMode sets the value of TerminalServiceMode for the instance
func (instance *Win32_TerminalServiceSettingError) SetPropertyTerminalServiceMode(value int32) (err error) {
	return instance.SetProperty("TerminalServiceMode", (value))
}

// GetTerminalServiceMode gets the value of TerminalServiceMode for the instance
func (instance *Win32_TerminalServiceSettingError) GetPropertyTerminalServiceMode() (value int32, err error) {
	retValue, err := instance.GetProperty("TerminalServiceMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
