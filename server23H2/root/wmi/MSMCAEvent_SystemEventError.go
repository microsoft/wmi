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

// MSMCAEvent_SystemEventError struct
type MSMCAEvent_SystemEventError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	//
	Cpu uint32

	//
	ErrorSeverity SystemEventError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	SEL_DATA1 uint8

	//
	SEL_DATA2 uint8

	//
	SEL_DATA3 uint8

	//
	SEL_EVENT_DIR_TYPE uint8

	//
	SEL_EVM_REV uint8

	//
	SEL_GENERATOR_ID uint16

	//
	SEL_RECORD_ID uint16

	//
	SEL_RECORD_TYPE uint8

	//
	SEL_SENSOR_NUM uint8

	//
	SEL_SENSOR_TYPE uint8

	//
	SEL_TIME_STAMP uint64

	//
	Size uint32

	//
	Type SystemEventError_Type

	//
	VALIDATION_BITS uint64
}

func NewMSMCAEvent_SystemEventErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_SystemEventError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_SystemEventError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_SystemEventErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_SystemEventError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_SystemEventError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_SystemEventError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyAdditionalErrors() (value uint32, err error) {
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

// SetCpu sets the value of Cpu for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_SystemEventError) SetPropertyErrorSeverity(value SystemEventError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyErrorSeverity() (value SystemEventError_ErrorSeverity, err error) {
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

	value = SystemEventError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSMCAEvent_SystemEventError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyLogToEventlog() (value uint32, err error) {
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

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_SystemEventError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyRecordId() (value uint64, err error) {
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

// SetSEL_DATA1 sets the value of SEL_DATA1 for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_DATA1(value uint8) (err error) {
	return instance.SetProperty("SEL_DATA1", (value))
}

// GetSEL_DATA1 gets the value of SEL_DATA1 for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_DATA1() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_DATA1")
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

// SetSEL_DATA2 sets the value of SEL_DATA2 for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_DATA2(value uint8) (err error) {
	return instance.SetProperty("SEL_DATA2", (value))
}

// GetSEL_DATA2 gets the value of SEL_DATA2 for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_DATA2() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_DATA2")
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

// SetSEL_DATA3 sets the value of SEL_DATA3 for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_DATA3(value uint8) (err error) {
	return instance.SetProperty("SEL_DATA3", (value))
}

// GetSEL_DATA3 gets the value of SEL_DATA3 for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_DATA3() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_DATA3")
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

// SetSEL_EVENT_DIR_TYPE sets the value of SEL_EVENT_DIR_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_EVENT_DIR_TYPE(value uint8) (err error) {
	return instance.SetProperty("SEL_EVENT_DIR_TYPE", (value))
}

// GetSEL_EVENT_DIR_TYPE gets the value of SEL_EVENT_DIR_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_EVENT_DIR_TYPE() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_EVENT_DIR_TYPE")
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

// SetSEL_EVM_REV sets the value of SEL_EVM_REV for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_EVM_REV(value uint8) (err error) {
	return instance.SetProperty("SEL_EVM_REV", (value))
}

// GetSEL_EVM_REV gets the value of SEL_EVM_REV for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_EVM_REV() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_EVM_REV")
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

// SetSEL_GENERATOR_ID sets the value of SEL_GENERATOR_ID for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_GENERATOR_ID(value uint16) (err error) {
	return instance.SetProperty("SEL_GENERATOR_ID", (value))
}

// GetSEL_GENERATOR_ID gets the value of SEL_GENERATOR_ID for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_GENERATOR_ID() (value uint16, err error) {
	retValue, err := instance.GetProperty("SEL_GENERATOR_ID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSEL_RECORD_ID sets the value of SEL_RECORD_ID for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_RECORD_ID(value uint16) (err error) {
	return instance.SetProperty("SEL_RECORD_ID", (value))
}

// GetSEL_RECORD_ID gets the value of SEL_RECORD_ID for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_RECORD_ID() (value uint16, err error) {
	retValue, err := instance.GetProperty("SEL_RECORD_ID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSEL_RECORD_TYPE sets the value of SEL_RECORD_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_RECORD_TYPE(value uint8) (err error) {
	return instance.SetProperty("SEL_RECORD_TYPE", (value))
}

// GetSEL_RECORD_TYPE gets the value of SEL_RECORD_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_RECORD_TYPE() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_RECORD_TYPE")
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

// SetSEL_SENSOR_NUM sets the value of SEL_SENSOR_NUM for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_SENSOR_NUM(value uint8) (err error) {
	return instance.SetProperty("SEL_SENSOR_NUM", (value))
}

// GetSEL_SENSOR_NUM gets the value of SEL_SENSOR_NUM for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_SENSOR_NUM() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_SENSOR_NUM")
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

// SetSEL_SENSOR_TYPE sets the value of SEL_SENSOR_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_SENSOR_TYPE(value uint8) (err error) {
	return instance.SetProperty("SEL_SENSOR_TYPE", (value))
}

// GetSEL_SENSOR_TYPE gets the value of SEL_SENSOR_TYPE for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_SENSOR_TYPE() (value uint8, err error) {
	retValue, err := instance.GetProperty("SEL_SENSOR_TYPE")
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

// SetSEL_TIME_STAMP sets the value of SEL_TIME_STAMP for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertySEL_TIME_STAMP(value uint64) (err error) {
	return instance.SetProperty("SEL_TIME_STAMP", (value))
}

// GetSEL_TIME_STAMP gets the value of SEL_TIME_STAMP for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySEL_TIME_STAMP() (value uint64, err error) {
	retValue, err := instance.GetProperty("SEL_TIME_STAMP")
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
func (instance *MSMCAEvent_SystemEventError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertySize() (value uint32, err error) {
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

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyType(value SystemEventError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyType() (value SystemEventError_Type, err error) {
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

	value = SystemEventError_Type(valuetmp)

	return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_SystemEventError) SetPropertyVALIDATION_BITS(value uint64) (err error) {
	return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_SystemEventError) GetPropertyVALIDATION_BITS() (value uint64, err error) {
	retValue, err := instance.GetProperty("VALIDATION_BITS")
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
