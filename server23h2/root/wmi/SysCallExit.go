// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SysCallExit struct
type SysCallExit struct {
	*PerfInfo_V2

	//
	SysCallNtStatus uint32
}

func NewSysCallExitEx1(instance *cim.WmiInstance) (newInstance *SysCallExit, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SysCallExit{
		PerfInfo_V2: tmp,
	}
	return
}

func NewSysCallExitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SysCallExit, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SysCallExit{
		PerfInfo_V2: tmp,
	}
	return
}

// SetSysCallNtStatus sets the value of SysCallNtStatus for the instance
func (instance *SysCallExit) SetPropertySysCallNtStatus(value uint32) (err error) {
	return instance.SetProperty("SysCallNtStatus", (value))
}

// GetSysCallNtStatus gets the value of SysCallNtStatus for the instance
func (instance *SysCallExit) GetPropertySysCallNtStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("SysCallNtStatus")
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
