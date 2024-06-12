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

// MSNdis_ReceiveFilterParameters struct
type MSNdis_ReceiveFilterParameters struct {
	*MSNdis

	//
	FieldParameters []MSNdis_ReceiveFilterFieldParameters

	//
	FieldParametersArrayElementSize uint32

	//
	FieldParametersArrayNumElements uint32

	//
	FieldParametersArrayOffset uint32

	//
	FilterId uint32

	//
	FilterType uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	QueueId uint32

	//
	RequestedFilterIdBitCount uint32
}

func NewMSNdis_ReceiveFilterParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterParameters{
		MSNdis: tmp,
	}
	return
}

// SetFieldParameters sets the value of FieldParameters for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFieldParameters(value []MSNdis_ReceiveFilterFieldParameters) (err error) {
	return instance.SetProperty("FieldParameters", (value))
}

// GetFieldParameters gets the value of FieldParameters for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFieldParameters() (value []MSNdis_ReceiveFilterFieldParameters, err error) {
	retValue, err := instance.GetProperty("FieldParameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSNdis_ReceiveFilterFieldParameters)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSNdis_ReceiveFilterFieldParameters is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSNdis_ReceiveFilterFieldParameters(valuetmp))
	}

	return
}

// SetFieldParametersArrayElementSize sets the value of FieldParametersArrayElementSize for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFieldParametersArrayElementSize(value uint32) (err error) {
	return instance.SetProperty("FieldParametersArrayElementSize", (value))
}

// GetFieldParametersArrayElementSize gets the value of FieldParametersArrayElementSize for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFieldParametersArrayElementSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("FieldParametersArrayElementSize")
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

// SetFieldParametersArrayNumElements sets the value of FieldParametersArrayNumElements for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFieldParametersArrayNumElements(value uint32) (err error) {
	return instance.SetProperty("FieldParametersArrayNumElements", (value))
}

// GetFieldParametersArrayNumElements gets the value of FieldParametersArrayNumElements for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFieldParametersArrayNumElements() (value uint32, err error) {
	retValue, err := instance.GetProperty("FieldParametersArrayNumElements")
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

// SetFieldParametersArrayOffset sets the value of FieldParametersArrayOffset for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFieldParametersArrayOffset(value uint32) (err error) {
	return instance.SetProperty("FieldParametersArrayOffset", (value))
}

// GetFieldParametersArrayOffset gets the value of FieldParametersArrayOffset for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFieldParametersArrayOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("FieldParametersArrayOffset")
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

// SetFilterId sets the value of FilterId for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFilterId(value uint32) (err error) {
	return instance.SetProperty("FilterId", (value))
}

// GetFilterId gets the value of FilterId for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFilterId() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterId")
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

// SetFilterType sets the value of FilterType for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFilterType(value uint32) (err error) {
	return instance.SetProperty("FilterType", (value))
}

// GetFilterType gets the value of FilterType for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFilterType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterType")
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
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyFlags() (value uint32, err error) {
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
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetQueueId sets the value of QueueId for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyQueueId(value uint32) (err error) {
	return instance.SetProperty("QueueId", (value))
}

// GetQueueId gets the value of QueueId for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyQueueId() (value uint32, err error) {
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

// SetRequestedFilterIdBitCount sets the value of RequestedFilterIdBitCount for the instance
func (instance *MSNdis_ReceiveFilterParameters) SetPropertyRequestedFilterIdBitCount(value uint32) (err error) {
	return instance.SetProperty("RequestedFilterIdBitCount", (value))
}

// GetRequestedFilterIdBitCount gets the value of RequestedFilterIdBitCount for the instance
func (instance *MSNdis_ReceiveFilterParameters) GetPropertyRequestedFilterIdBitCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestedFilterIdBitCount")
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
