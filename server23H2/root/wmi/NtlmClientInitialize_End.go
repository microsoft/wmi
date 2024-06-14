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

// NtlmClientInitialize_End struct
type NtlmClientInitialize_End struct {
	*NtlmClientInitialize

	// In-Context
	InContext uint32

	// Out-Context
	OutContext uint32

	// Stage Hint
	StageHint uint32

	// Status
	Status uint32
}

func NewNtlmClientInitialize_EndEx1(instance *cim.WmiInstance) (newInstance *NtlmClientInitialize_End, err error) {
	tmp, err := NewNtlmClientInitializeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmClientInitialize_End{
		NtlmClientInitialize: tmp,
	}
	return
}

func NewNtlmClientInitialize_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmClientInitialize_End, err error) {
	tmp, err := NewNtlmClientInitializeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmClientInitialize_End{
		NtlmClientInitialize: tmp,
	}
	return
}

// SetInContext sets the value of InContext for the instance
func (instance *NtlmClientInitialize_End) SetPropertyInContext(value uint32) (err error) {
	return instance.SetProperty("InContext", (value))
}

// GetInContext gets the value of InContext for the instance
func (instance *NtlmClientInitialize_End) GetPropertyInContext() (value uint32, err error) {
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

// SetOutContext sets the value of OutContext for the instance
func (instance *NtlmClientInitialize_End) SetPropertyOutContext(value uint32) (err error) {
	return instance.SetProperty("OutContext", (value))
}

// GetOutContext gets the value of OutContext for the instance
func (instance *NtlmClientInitialize_End) GetPropertyOutContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutContext")
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
func (instance *NtlmClientInitialize_End) SetPropertyStageHint(value uint32) (err error) {
	return instance.SetProperty("StageHint", (value))
}

// GetStageHint gets the value of StageHint for the instance
func (instance *NtlmClientInitialize_End) GetPropertyStageHint() (value uint32, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *NtlmClientInitialize_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *NtlmClientInitialize_End) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
