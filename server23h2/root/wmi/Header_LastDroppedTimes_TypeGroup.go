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

// Header_LastDroppedTimes_TypeGroup struct
type Header_LastDroppedTimes_TypeGroup struct {
	*EventTraceEvent

	//
	Padding uint32

	//
	TimeStamp []uint64

	//
	TimeStampCount uint32
}

func NewHeader_LastDroppedTimes_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_LastDroppedTimes_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Header_LastDroppedTimes_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

func NewHeader_LastDroppedTimes_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Header_LastDroppedTimes_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Header_LastDroppedTimes_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

// SetPadding sets the value of Padding for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) SetPropertyPadding(value uint32) (err error) {
	return instance.SetProperty("Padding", (value))
}

// GetPadding gets the value of Padding for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) GetPropertyPadding() (value uint32, err error) {
	retValue, err := instance.GetProperty("Padding")
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

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) SetPropertyTimeStamp(value []uint64) (err error) {
	return instance.SetProperty("TimeStamp", (value))
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) GetPropertyTimeStamp() (value []uint64, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
	}

	return
}

// SetTimeStampCount sets the value of TimeStampCount for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) SetPropertyTimeStampCount(value uint32) (err error) {
	return instance.SetProperty("TimeStampCount", (value))
}

// GetTimeStampCount gets the value of TimeStampCount for the instance
func (instance *Header_LastDroppedTimes_TypeGroup) GetPropertyTimeStampCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeStampCount")
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
