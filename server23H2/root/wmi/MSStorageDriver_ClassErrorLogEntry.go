// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSStorageDriver_ClassErrorLogEntry struct
type MSStorageDriver_ClassErrorLogEntry struct {
	*cim.WmiInstance

	// Error Paging
	errorPaging bool

	// Error Reserved
	errorReserved uint8

	// Error Retried
	errorRetried bool

	// Error Unhandled
	errorUnhandled bool

	// Event Time
	eventTime string

	// Port Number
	portNumber uint32

	// Reserved
	reserved []uint8

	// Sense Data
	senseData MSStorageDriver_SenseData

	// SCSI Request Block
	srb MSStorageDriver_ScsiRequestBlock

	// Tick Count
	tickCount uint64
}

func NewMSStorageDriver_ClassErrorLogEntryEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_ClassErrorLogEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_ClassErrorLogEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSStorageDriver_ClassErrorLogEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver_ClassErrorLogEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_ClassErrorLogEntry{
		WmiInstance: tmp,
	}
	return
}

// SeterrorPaging sets the value of errorPaging for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyerrorPaging(value bool) (err error) {
	return instance.SetProperty("errorPaging", (value))
}

// GeterrorPaging gets the value of errorPaging for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyerrorPaging() (value bool, err error) {
	retValue, err := instance.GetProperty("errorPaging")
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

// SeterrorReserved sets the value of errorReserved for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyerrorReserved(value uint8) (err error) {
	return instance.SetProperty("errorReserved", (value))
}

// GeterrorReserved gets the value of errorReserved for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyerrorReserved() (value uint8, err error) {
	retValue, err := instance.GetProperty("errorReserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SeterrorRetried sets the value of errorRetried for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyerrorRetried(value bool) (err error) {
	return instance.SetProperty("errorRetried", (value))
}

// GeterrorRetried gets the value of errorRetried for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyerrorRetried() (value bool, err error) {
	retValue, err := instance.GetProperty("errorRetried")
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

// SeterrorUnhandled sets the value of errorUnhandled for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyerrorUnhandled(value bool) (err error) {
	return instance.SetProperty("errorUnhandled", (value))
}

// GeterrorUnhandled gets the value of errorUnhandled for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyerrorUnhandled() (value bool, err error) {
	retValue, err := instance.GetProperty("errorUnhandled")
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

// SeteventTime sets the value of eventTime for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyeventTime(value string) (err error) {
	return instance.SetProperty("eventTime", (value))
}

// GeteventTime gets the value of eventTime for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyeventTime() (value string, err error) {
	retValue, err := instance.GetProperty("eventTime")
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

// SetportNumber sets the value of portNumber for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyportNumber(value uint32) (err error) {
	return instance.SetProperty("portNumber", (value))
}

// GetportNumber gets the value of portNumber for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyportNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("portNumber")
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

// Setreserved sets the value of reserved for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertyreserved(value []uint8) (err error) {
	return instance.SetProperty("reserved", (value))
}

// Getreserved gets the value of reserved for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertyreserved() (value []uint8, err error) {
	retValue, err := instance.GetProperty("reserved")
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

// SetsenseData sets the value of senseData for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertysenseData(value MSStorageDriver_SenseData) (err error) {
	return instance.SetProperty("senseData", (value))
}

// GetsenseData gets the value of senseData for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertysenseData() (value MSStorageDriver_SenseData, err error) {
	retValue, err := instance.GetProperty("senseData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSStorageDriver_SenseData)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSStorageDriver_SenseData is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSStorageDriver_SenseData(valuetmp)

	return
}

// Setsrb sets the value of srb for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertysrb(value MSStorageDriver_ScsiRequestBlock) (err error) {
	return instance.SetProperty("srb", (value))
}

// Getsrb gets the value of srb for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertysrb() (value MSStorageDriver_ScsiRequestBlock, err error) {
	retValue, err := instance.GetProperty("srb")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSStorageDriver_ScsiRequestBlock)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSStorageDriver_ScsiRequestBlock is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSStorageDriver_ScsiRequestBlock(valuetmp)

	return
}

// SettickCount sets the value of tickCount for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) SetPropertytickCount(value uint64) (err error) {
	return instance.SetProperty("tickCount", (value))
}

// GettickCount gets the value of tickCount for the instance
func (instance *MSStorageDriver_ClassErrorLogEntry) GetPropertytickCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("tickCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
