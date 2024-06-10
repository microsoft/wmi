// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_InterruptModerationParameters struct
type MSNdis_InterruptModerationParameters struct {
	*MSNdis

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	InterruptModeration uint32
}

func NewMSNdis_InterruptModerationParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_InterruptModerationParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_InterruptModerationParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_InterruptModerationParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_InterruptModerationParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_InterruptModerationParameters{
		MSNdis: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_InterruptModerationParameters) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_InterruptModerationParameters) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_InterruptModerationParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_InterruptModerationParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetInterruptModeration sets the value of InterruptModeration for the instance
func (instance *MSNdis_InterruptModerationParameters) SetPropertyInterruptModeration(value uint32) (err error) {
	return instance.SetProperty("InterruptModeration", (value))
}

// GetInterruptModeration gets the value of InterruptModeration for the instance
func (instance *MSNdis_InterruptModerationParameters) GetPropertyInterruptModeration() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterruptModeration")
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
