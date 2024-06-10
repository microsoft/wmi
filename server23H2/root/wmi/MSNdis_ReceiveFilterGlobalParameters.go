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

// MSNdis_ReceiveFilterGlobalParameters struct
type MSNdis_ReceiveFilterGlobalParameters struct {
	*MSNdis

	//
	EnabledFilterTypes uint32

	//
	EnabledQueueTypes uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader
}

func NewMSNdis_ReceiveFilterGlobalParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterGlobalParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterGlobalParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterGlobalParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterGlobalParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterGlobalParameters{
		MSNdis: tmp,
	}
	return
}

// SetEnabledFilterTypes sets the value of EnabledFilterTypes for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) SetPropertyEnabledFilterTypes(value uint32) (err error) {
	return instance.SetProperty("EnabledFilterTypes", (value))
}

// GetEnabledFilterTypes gets the value of EnabledFilterTypes for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) GetPropertyEnabledFilterTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledFilterTypes")
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

// SetEnabledQueueTypes sets the value of EnabledQueueTypes for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) SetPropertyEnabledQueueTypes(value uint32) (err error) {
	return instance.SetProperty("EnabledQueueTypes", (value))
}

// GetEnabledQueueTypes gets the value of EnabledQueueTypes for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) GetPropertyEnabledQueueTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledQueueTypes")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) GetPropertyFlags() (value uint32, err error) {
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
func (instance *MSNdis_ReceiveFilterGlobalParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterGlobalParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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
