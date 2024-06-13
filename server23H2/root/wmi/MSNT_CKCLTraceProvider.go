// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNT_CKCLTraceProvider struct
type MSNT_CKCLTraceProvider struct {
	*EventTrace

	//
	Flags CKCLTraceProvider_Flags
}

func NewMSNT_CKCLTraceProviderEx1(instance *cim.WmiInstance) (newInstance *MSNT_CKCLTraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNT_CKCLTraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewMSNT_CKCLTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNT_CKCLTraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNT_CKCLTraceProvider{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNT_CKCLTraceProvider) SetPropertyFlags(value CKCLTraceProvider_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNT_CKCLTraceProvider) GetPropertyFlags() (value CKCLTraceProvider_Flags, err error) {
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

	value = CKCLTraceProvider_Flags(valuetmp)

	return
}
