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

// MSNdis_ReceiveFilterFieldParameters struct
type MSNdis_ReceiveFilterFieldParameters struct {
	*MSNdis

	//
	FieldByteArrayValue []uint8

	//
	Flags uint32

	//
	FrameHeader uint32

	//
	Header MSNdis_ObjectHeader

	//
	MacHeaderField uint32

	//
	ReceiveFilterTest uint32

	//
	ResultByteArrayValue []uint8
}

func NewMSNdis_ReceiveFilterFieldParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterFieldParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterFieldParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterFieldParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterFieldParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterFieldParameters{
		MSNdis: tmp,
	}
	return
}

// SetFieldByteArrayValue sets the value of FieldByteArrayValue for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyFieldByteArrayValue(value []uint8) (err error) {
	return instance.SetProperty("FieldByteArrayValue", (value))
}

// GetFieldByteArrayValue gets the value of FieldByteArrayValue for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyFieldByteArrayValue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FieldByteArrayValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyFlags() (value uint32, err error) {
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

// SetFrameHeader sets the value of FrameHeader for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyFrameHeader(value uint32) (err error) {
	return instance.SetProperty("FrameHeader", (value))
}

// GetFrameHeader gets the value of FrameHeader for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyFrameHeader() (value uint32, err error) {
	retValue, err := instance.GetProperty("FrameHeader")
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
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetMacHeaderField sets the value of MacHeaderField for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyMacHeaderField(value uint32) (err error) {
	return instance.SetProperty("MacHeaderField", (value))
}

// GetMacHeaderField gets the value of MacHeaderField for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyMacHeaderField() (value uint32, err error) {
	retValue, err := instance.GetProperty("MacHeaderField")
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

// SetReceiveFilterTest sets the value of ReceiveFilterTest for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyReceiveFilterTest(value uint32) (err error) {
	return instance.SetProperty("ReceiveFilterTest", (value))
}

// GetReceiveFilterTest gets the value of ReceiveFilterTest for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyReceiveFilterTest() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveFilterTest")
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

// SetResultByteArrayValue sets the value of ResultByteArrayValue for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) SetPropertyResultByteArrayValue(value []uint8) (err error) {
	return instance.SetProperty("ResultByteArrayValue", (value))
}

// GetResultByteArrayValue gets the value of ResultByteArrayValue for the instance
func (instance *MSNdis_ReceiveFilterFieldParameters) GetPropertyResultByteArrayValue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ResultByteArrayValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
