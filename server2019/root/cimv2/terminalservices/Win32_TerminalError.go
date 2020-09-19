// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TerminalError struct
type Win32_TerminalError struct {
	*__ExtendedStatus

	//
	TerminalName string
}

func NewWin32_TerminalErrorEx1(instance *cim.WmiInstance) (newInstance *Win32_TerminalError, err error) {
	tmp, err := New__ExtendedStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalError{
		__ExtendedStatus: tmp,
	}
	return
}

func NewWin32_TerminalErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TerminalError, err error) {
	tmp, err := New__ExtendedStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalError{
		__ExtendedStatus: tmp,
	}
	return
}

// SetTerminalName sets the value of TerminalName for the instance
func (instance *Win32_TerminalError) SetPropertyTerminalName(value string) (err error) {
	return instance.SetProperty("TerminalName", (value))
}

// GetTerminalName gets the value of TerminalName for the instance
func (instance *Win32_TerminalError) GetPropertyTerminalName() (value string, err error) {
	retValue, err := instance.GetProperty("TerminalName")
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
