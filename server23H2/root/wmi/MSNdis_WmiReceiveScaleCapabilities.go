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

// MSNdis_WmiReceiveScaleCapabilities struct
type MSNdis_WmiReceiveScaleCapabilities struct {
	*MSNdis

	//
	CapabilitiesFlags uint32

	//
	Header MSNdis_ObjectHeader

	//
	NumberOfInterruptMessages uint32

	//
	NumberOfReceiveQueues uint32
}

func NewMSNdis_WmiReceiveScaleCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiReceiveScaleCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiReceiveScaleCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiReceiveScaleCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiReceiveScaleCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiReceiveScaleCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetCapabilitiesFlags sets the value of CapabilitiesFlags for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) SetPropertyCapabilitiesFlags(value uint32) (err error) {
	return instance.SetProperty("CapabilitiesFlags", (value))
}

// GetCapabilitiesFlags gets the value of CapabilitiesFlags for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) GetPropertyCapabilitiesFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("CapabilitiesFlags")
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
func (instance *MSNdis_WmiReceiveScaleCapabilities) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetNumberOfInterruptMessages sets the value of NumberOfInterruptMessages for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) SetPropertyNumberOfInterruptMessages(value uint32) (err error) {
	return instance.SetProperty("NumberOfInterruptMessages", (value))
}

// GetNumberOfInterruptMessages gets the value of NumberOfInterruptMessages for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) GetPropertyNumberOfInterruptMessages() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInterruptMessages")
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

// SetNumberOfReceiveQueues sets the value of NumberOfReceiveQueues for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) SetPropertyNumberOfReceiveQueues(value uint32) (err error) {
	return instance.SetProperty("NumberOfReceiveQueues", (value))
}

// GetNumberOfReceiveQueues gets the value of NumberOfReceiveQueues for the instance
func (instance *MSNdis_WmiReceiveScaleCapabilities) GetPropertyNumberOfReceiveQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfReceiveQueues")
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
