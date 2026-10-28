// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WlanDlgTraceGuid struct
type WlanDlgTraceGuid struct {
	*EventTrace

	// Enable Flags
	Flags WlanDlgTraceGuid_Flags
}

func NewWlanDlgTraceGuidEx1(instance *cim.WmiInstance) (newInstance *WlanDlgTraceGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WlanDlgTraceGuid{
		EventTrace: tmp,
	}
	return
}

func NewWlanDlgTraceGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WlanDlgTraceGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WlanDlgTraceGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *WlanDlgTraceGuid) SetPropertyFlags(value WlanDlgTraceGuid_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WlanDlgTraceGuid) GetPropertyFlags() (value WlanDlgTraceGuid_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
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

	value = WlanDlgTraceGuid_Flags(valuetmp)

	return
}
