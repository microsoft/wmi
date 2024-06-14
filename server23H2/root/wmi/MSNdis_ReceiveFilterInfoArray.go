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

// MSNdis_ReceiveFilterInfoArray struct
type MSNdis_ReceiveFilterInfoArray struct {
	*MSNdis

	//
	ElementSize uint32

	//
	Filter []MSNdis_ReceiveFilterInfo

	//
	FirstElementOffset uint32

	//
	Header MSNdis_ObjectHeader

	//
	NumElements uint32

	//
	QueueId uint32
}

func NewMSNdis_ReceiveFilterInfoArrayEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterInfoArray, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterInfoArray{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterInfoArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterInfoArray, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterInfoArray{
		MSNdis: tmp,
	}
	return
}

// SetElementSize sets the value of ElementSize for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyElementSize(value uint32) (err error) {
	return instance.SetProperty("ElementSize", (value))
}

// GetElementSize gets the value of ElementSize for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyElementSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ElementSize")
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

// SetFilter sets the value of Filter for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyFilter(value []MSNdis_ReceiveFilterInfo) (err error) {
	return instance.SetProperty("Filter", (value))
}

// GetFilter gets the value of Filter for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyFilter() (value []MSNdis_ReceiveFilterInfo, err error) {
	retValue, err := instance.GetProperty("Filter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSNdis_ReceiveFilterInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSNdis_ReceiveFilterInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSNdis_ReceiveFilterInfo(valuetmp))
	}

	return
}

// SetFirstElementOffset sets the value of FirstElementOffset for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyFirstElementOffset(value uint32) (err error) {
	return instance.SetProperty("FirstElementOffset", (value))
}

// GetFirstElementOffset gets the value of FirstElementOffset for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyFirstElementOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirstElementOffset")
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
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetNumElements sets the value of NumElements for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyNumElements(value uint32) (err error) {
	return instance.SetProperty("NumElements", (value))
}

// GetNumElements gets the value of NumElements for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyNumElements() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumElements")
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

// SetQueueId sets the value of QueueId for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) SetPropertyQueueId(value uint32) (err error) {
	return instance.SetProperty("QueueId", (value))
}

// GetQueueId gets the value of QueueId for the instance
func (instance *MSNdis_ReceiveFilterInfoArray) GetPropertyQueueId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueId")
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
