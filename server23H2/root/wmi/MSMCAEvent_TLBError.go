// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSMCAEvent_TLBError struct
type MSMCAEvent_TLBError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	// The address at which the error occurred.
	Address uint64

	//
	Cpu uint32

	//
	ErrorSeverity TLBError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	MemoryHierarchyLevel TLBError_MemoryHierarchyLevel

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	Size uint32

	//
	TransactionType TLBError_TransactionType

	//
	Type TLBError_Type
}

func NewMSMCAEvent_TLBErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_TLBError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_TLBError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_TLBErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_TLBError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_TLBError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyActive() (value bool, err error) {
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

// SetAdditionalErrors sets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyAdditionalErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdditionalErrors")
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

// SetAddress sets the value of Address for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyAddress(value uint64) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetCpu sets the value of Cpu for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyCpu() (value uint32, err error) {
	retValue, err := instance.GetProperty("Cpu")
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

// SetErrorSeverity sets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyErrorSeverity(value TLBError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyErrorSeverity() (value TLBError_ErrorSeverity, err error) {
	retValue, err := instance.GetProperty("ErrorSeverity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TLBError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyInstanceName() (value string, err error) {
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

// SetLogToEventlog sets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyLogToEventlog() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogToEventlog")
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

// SetMemoryHierarchyLevel sets the value of MemoryHierarchyLevel for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyMemoryHierarchyLevel(value TLBError_MemoryHierarchyLevel) (err error) {
	return instance.SetProperty("MemoryHierarchyLevel", (value))
}

// GetMemoryHierarchyLevel gets the value of MemoryHierarchyLevel for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyMemoryHierarchyLevel() (value TLBError_MemoryHierarchyLevel, err error) {
	retValue, err := instance.GetProperty("MemoryHierarchyLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TLBError_MemoryHierarchyLevel(valuetmp)

	return
}

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyRawRecord() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RawRecord")
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

// SetRecordId sets the value of RecordId for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyRecordId() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecordId")
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

// SetSize sets the value of Size for the instance
func (instance *MSMCAEvent_TLBError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_TLBError) GetPropertySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetTransactionType sets the value of TransactionType for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyTransactionType(value TLBError_TransactionType) (err error) {
	return instance.SetProperty("TransactionType", (value))
}

// GetTransactionType gets the value of TransactionType for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyTransactionType() (value TLBError_TransactionType, err error) {
	retValue, err := instance.GetProperty("TransactionType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TLBError_TransactionType(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_TLBError) SetPropertyType(value TLBError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_TLBError) GetPropertyType() (value TLBError_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TLBError_Type(valuetmp)

	return
}
