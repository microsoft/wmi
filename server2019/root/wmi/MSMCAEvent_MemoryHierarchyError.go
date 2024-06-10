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

// MSMCAEvent_MemoryHierarchyError struct
type MSMCAEvent_MemoryHierarchyError struct {
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
	ErrorSeverity MemoryHierarchyError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	MemoryHierarchyLevel MemoryHierarchyError_MemoryHierarchyLevel

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	RequestType MemoryHierarchyError_RequestType

	//
	Size uint32

	//
	TransactionType MemoryHierarchyError_TransactionType

	//
	Type MemoryHierarchyError_Type
}

func NewMSMCAEvent_MemoryHierarchyErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_MemoryHierarchyError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_MemoryHierarchyError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_MemoryHierarchyErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_MemoryHierarchyError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_MemoryHierarchyError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyAdditionalErrors() (value uint32, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyAddress(value uint64) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyAddress() (value uint64, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyErrorSeverity(value MemoryHierarchyError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyErrorSeverity() (value MemoryHierarchyError_ErrorSeverity, err error) {
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

	value = MemoryHierarchyError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyLogToEventlog() (value uint32, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyMemoryHierarchyLevel(value MemoryHierarchyError_MemoryHierarchyLevel) (err error) {
	return instance.SetProperty("MemoryHierarchyLevel", (value))
}

// GetMemoryHierarchyLevel gets the value of MemoryHierarchyLevel for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyMemoryHierarchyLevel() (value MemoryHierarchyError_MemoryHierarchyLevel, err error) {
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

	value = MemoryHierarchyError_MemoryHierarchyLevel(valuetmp)

	return
}

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyRecordId() (value uint64, err error) {
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

// SetRequestType sets the value of RequestType for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyRequestType(value MemoryHierarchyError_RequestType) (err error) {
	return instance.SetProperty("RequestType", (value))
}

// GetRequestType gets the value of RequestType for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyRequestType() (value MemoryHierarchyError_RequestType, err error) {
	retValue, err := instance.GetProperty("RequestType")
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

	value = MemoryHierarchyError_RequestType(valuetmp)

	return
}

// SetSize sets the value of Size for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertySize() (value uint32, err error) {
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
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyTransactionType(value MemoryHierarchyError_TransactionType) (err error) {
	return instance.SetProperty("TransactionType", (value))
}

// GetTransactionType gets the value of TransactionType for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyTransactionType() (value MemoryHierarchyError_TransactionType, err error) {
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

	value = MemoryHierarchyError_TransactionType(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) SetPropertyType(value MemoryHierarchyError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_MemoryHierarchyError) GetPropertyType() (value MemoryHierarchyError_Type, err error) {
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

	value = MemoryHierarchyError_Type(valuetmp)

	return
}
