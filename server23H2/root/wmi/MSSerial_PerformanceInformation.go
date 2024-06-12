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

// MSSerial_PerformanceInformation struct
type MSSerial_PerformanceInformation struct {
	*MSSerial

	//
	Active bool

	//
	BufferOverrunErrorCount uint32

	//
	FrameErrorCount uint32

	//
	InstanceName string

	//
	ParityErrorCount uint32

	//
	ReceivedCount uint32

	//
	SerialOverrunErrorCount uint32

	//
	TransmittedCount uint32
}

func NewMSSerial_PerformanceInformationEx1(instance *cim.WmiInstance) (newInstance *MSSerial_PerformanceInformation, err error) {
	tmp, err := NewMSSerialEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSerial_PerformanceInformation{
		MSSerial: tmp,
	}
	return
}

func NewMSSerial_PerformanceInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSerial_PerformanceInformation, err error) {
	tmp, err := NewMSSerialEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSerial_PerformanceInformation{
		MSSerial: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBufferOverrunErrorCount sets the value of BufferOverrunErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyBufferOverrunErrorCount(value uint32) (err error) {
	return instance.SetProperty("BufferOverrunErrorCount", (value))
}

// GetBufferOverrunErrorCount gets the value of BufferOverrunErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyBufferOverrunErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferOverrunErrorCount")
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

// SetFrameErrorCount sets the value of FrameErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyFrameErrorCount(value uint32) (err error) {
	return instance.SetProperty("FrameErrorCount", (value))
}

// GetFrameErrorCount gets the value of FrameErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyFrameErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FrameErrorCount")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetParityErrorCount sets the value of ParityErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyParityErrorCount(value uint32) (err error) {
	return instance.SetProperty("ParityErrorCount", (value))
}

// GetParityErrorCount gets the value of ParityErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyParityErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParityErrorCount")
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

// SetReceivedCount sets the value of ReceivedCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyReceivedCount(value uint32) (err error) {
	return instance.SetProperty("ReceivedCount", (value))
}

// GetReceivedCount gets the value of ReceivedCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyReceivedCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceivedCount")
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

// SetSerialOverrunErrorCount sets the value of SerialOverrunErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertySerialOverrunErrorCount(value uint32) (err error) {
	return instance.SetProperty("SerialOverrunErrorCount", (value))
}

// GetSerialOverrunErrorCount gets the value of SerialOverrunErrorCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertySerialOverrunErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SerialOverrunErrorCount")
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

// SetTransmittedCount sets the value of TransmittedCount for the instance
func (instance *MSSerial_PerformanceInformation) SetPropertyTransmittedCount(value uint32) (err error) {
	return instance.SetProperty("TransmittedCount", (value))
}

// GetTransmittedCount gets the value of TransmittedCount for the instance
func (instance *MSSerial_PerformanceInformation) GetPropertyTransmittedCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransmittedCount")
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
