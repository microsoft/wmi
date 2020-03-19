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

// Win32_JobObjectStatus struct
type Win32_JobObjectStatus struct {
	*__ExtendedStatus

	//
	AdditionalDescription string

	//
	Win32ErrorCode uint32
}

func NewWin32_JobObjectStatusEx1(instance *cim.WmiInstance) (newInstance *Win32_JobObjectStatus, err error) {
	tmp, err := New__ExtendedStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_JobObjectStatus{
		__ExtendedStatus: tmp,
	}
	return
}

func NewWin32_JobObjectStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_JobObjectStatus, err error) {
	tmp, err := New__ExtendedStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_JobObjectStatus{
		__ExtendedStatus: tmp,
	}
	return
}

// SetAdditionalDescription sets the value of AdditionalDescription for the instance
func (instance *Win32_JobObjectStatus) SetPropertyAdditionalDescription(value string) (err error) {
	return instance.SetProperty("AdditionalDescription", value)
}

// GetAdditionalDescription gets the value of AdditionalDescription for the instance
func (instance *Win32_JobObjectStatus) GetPropertyAdditionalDescription() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWin32ErrorCode sets the value of Win32ErrorCode for the instance
func (instance *Win32_JobObjectStatus) SetPropertyWin32ErrorCode(value uint32) (err error) {
	return instance.SetProperty("Win32ErrorCode", value)
}

// GetWin32ErrorCode gets the value of Win32ErrorCode for the instance
func (instance *Win32_JobObjectStatus) GetPropertyWin32ErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Win32ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
