// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NtlmClientInitialize_Start struct
type NtlmClientInitialize_Start struct {
	*NtlmClientInitialize

	// In-Context
	InContext uint32

	// Stage Hint
	StageHint uint32
}

func NewNtlmClientInitialize_StartEx1(instance *cim.WmiInstance) (newInstance *NtlmClientInitialize_Start, err error) {
	tmp, err := NewNtlmClientInitializeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmClientInitialize_Start{
		NtlmClientInitialize: tmp,
	}
	return
}

func NewNtlmClientInitialize_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmClientInitialize_Start, err error) {
	tmp, err := NewNtlmClientInitializeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmClientInitialize_Start{
		NtlmClientInitialize: tmp,
	}
	return
}

// SetInContext sets the value of InContext for the instance
func (instance *NtlmClientInitialize_Start) SetPropertyInContext(value uint32) (err error) {
	return instance.SetProperty("InContext", (value))
}

// GetInContext gets the value of InContext for the instance
func (instance *NtlmClientInitialize_Start) GetPropertyInContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("InContext")
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

// SetStageHint sets the value of StageHint for the instance
func (instance *NtlmClientInitialize_Start) SetPropertyStageHint(value uint32) (err error) {
	return instance.SetProperty("StageHint", (value))
}

// GetStageHint gets the value of StageHint for the instance
func (instance *NtlmClientInitialize_Start) GetPropertyStageHint() (value uint32, err error) {
	retValue, err := instance.GetProperty("StageHint")
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
