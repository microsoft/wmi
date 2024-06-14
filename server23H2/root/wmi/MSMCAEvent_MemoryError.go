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

// MSMCAEvent_MemoryError struct
type MSMCAEvent_MemoryError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	//
	BUS_SPECIFIC_DATA uint64

	//
	Cpu uint32

	//
	ErrorSeverity MemoryError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	MEM_BANK uint16

	//
	MEM_BIT_POSITION uint16

	//
	MEM_CARD uint16

	//
	MEM_COLUMN uint16

	//
	MEM_ERROR_STATUS uint64

	//
	MEM_MODULE uint16

	//
	MEM_NODE uint16

	//
	MEM_PHYSICAL_ADDR uint64

	//
	MEM_PHYSICAL_MASK uint64

	//
	MEM_ROW uint16

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	REQUESTOR_ID uint64

	//
	RESPONDER_ID uint64

	//
	Size uint32

	//
	TARGET_ID uint64

	//
	Type MemoryError_Type

	//
	VALIDATION_BITS uint64

	//
	xMEM_DEVICE uint16
}

func NewMSMCAEvent_MemoryErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_MemoryError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_MemoryError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_MemoryErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_MemoryError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_MemoryError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_MemoryError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyAdditionalErrors() (value uint32, err error) {
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

// SetBUS_SPECIFIC_DATA sets the value of BUS_SPECIFIC_DATA for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyBUS_SPECIFIC_DATA(value uint64) (err error) {
	return instance.SetProperty("BUS_SPECIFIC_DATA", (value))
}

// GetBUS_SPECIFIC_DATA gets the value of BUS_SPECIFIC_DATA for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyBUS_SPECIFIC_DATA() (value uint64, err error) {
	retValue, err := instance.GetProperty("BUS_SPECIFIC_DATA")
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
func (instance *MSMCAEvent_MemoryError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_MemoryError) SetPropertyErrorSeverity(value MemoryError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyErrorSeverity() (value MemoryError_ErrorSeverity, err error) {
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

	value = MemoryError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSMCAEvent_MemoryError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyLogToEventlog() (value uint32, err error) {
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

// SetMEM_BANK sets the value of MEM_BANK for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_BANK(value uint16) (err error) {
	return instance.SetProperty("MEM_BANK", (value))
}

// GetMEM_BANK gets the value of MEM_BANK for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_BANK() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_BANK")
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

// SetMEM_BIT_POSITION sets the value of MEM_BIT_POSITION for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_BIT_POSITION(value uint16) (err error) {
	return instance.SetProperty("MEM_BIT_POSITION", (value))
}

// GetMEM_BIT_POSITION gets the value of MEM_BIT_POSITION for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_BIT_POSITION() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_BIT_POSITION")
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

// SetMEM_CARD sets the value of MEM_CARD for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_CARD(value uint16) (err error) {
	return instance.SetProperty("MEM_CARD", (value))
}

// GetMEM_CARD gets the value of MEM_CARD for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_CARD() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_CARD")
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

// SetMEM_COLUMN sets the value of MEM_COLUMN for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_COLUMN(value uint16) (err error) {
	return instance.SetProperty("MEM_COLUMN", (value))
}

// GetMEM_COLUMN gets the value of MEM_COLUMN for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_COLUMN() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_COLUMN")
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

// SetMEM_ERROR_STATUS sets the value of MEM_ERROR_STATUS for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_ERROR_STATUS(value uint64) (err error) {
	return instance.SetProperty("MEM_ERROR_STATUS", (value))
}

// GetMEM_ERROR_STATUS gets the value of MEM_ERROR_STATUS for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_ERROR_STATUS() (value uint64, err error) {
	retValue, err := instance.GetProperty("MEM_ERROR_STATUS")
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

// SetMEM_MODULE sets the value of MEM_MODULE for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_MODULE(value uint16) (err error) {
	return instance.SetProperty("MEM_MODULE", (value))
}

// GetMEM_MODULE gets the value of MEM_MODULE for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_MODULE() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_MODULE")
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

// SetMEM_NODE sets the value of MEM_NODE for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_NODE(value uint16) (err error) {
	return instance.SetProperty("MEM_NODE", (value))
}

// GetMEM_NODE gets the value of MEM_NODE for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_NODE() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_NODE")
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

// SetMEM_PHYSICAL_ADDR sets the value of MEM_PHYSICAL_ADDR for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_PHYSICAL_ADDR(value uint64) (err error) {
	return instance.SetProperty("MEM_PHYSICAL_ADDR", (value))
}

// GetMEM_PHYSICAL_ADDR gets the value of MEM_PHYSICAL_ADDR for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_PHYSICAL_ADDR() (value uint64, err error) {
	retValue, err := instance.GetProperty("MEM_PHYSICAL_ADDR")
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

// SetMEM_PHYSICAL_MASK sets the value of MEM_PHYSICAL_MASK for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_PHYSICAL_MASK(value uint64) (err error) {
	return instance.SetProperty("MEM_PHYSICAL_MASK", (value))
}

// GetMEM_PHYSICAL_MASK gets the value of MEM_PHYSICAL_MASK for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_PHYSICAL_MASK() (value uint64, err error) {
	retValue, err := instance.GetProperty("MEM_PHYSICAL_MASK")
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

// SetMEM_ROW sets the value of MEM_ROW for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyMEM_ROW(value uint16) (err error) {
	return instance.SetProperty("MEM_ROW", (value))
}

// GetMEM_ROW gets the value of MEM_ROW for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyMEM_ROW() (value uint16, err error) {
	retValue, err := instance.GetProperty("MEM_ROW")
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

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_MemoryError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyRecordId() (value uint64, err error) {
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

// SetREQUESTOR_ID sets the value of REQUESTOR_ID for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyREQUESTOR_ID(value uint64) (err error) {
	return instance.SetProperty("REQUESTOR_ID", (value))
}

// GetREQUESTOR_ID gets the value of REQUESTOR_ID for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyREQUESTOR_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("REQUESTOR_ID")
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

// SetRESPONDER_ID sets the value of RESPONDER_ID for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyRESPONDER_ID(value uint64) (err error) {
	return instance.SetProperty("RESPONDER_ID", (value))
}

// GetRESPONDER_ID gets the value of RESPONDER_ID for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyRESPONDER_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("RESPONDER_ID")
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
func (instance *MSMCAEvent_MemoryError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertySize() (value uint32, err error) {
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

// SetTARGET_ID sets the value of TARGET_ID for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyTARGET_ID(value uint64) (err error) {
	return instance.SetProperty("TARGET_ID", (value))
}

// GetTARGET_ID gets the value of TARGET_ID for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyTARGET_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("TARGET_ID")
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

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyType(value MemoryError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyType() (value MemoryError_Type, err error) {
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

	value = MemoryError_Type(valuetmp)

	return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyVALIDATION_BITS(value uint64) (err error) {
	return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyVALIDATION_BITS() (value uint64, err error) {
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

// SetxMEM_DEVICE sets the value of xMEM_DEVICE for the instance
func (instance *MSMCAEvent_MemoryError) SetPropertyxMEM_DEVICE(value uint16) (err error) {
	return instance.SetProperty("xMEM_DEVICE", (value))
}

// GetxMEM_DEVICE gets the value of xMEM_DEVICE for the instance
func (instance *MSMCAEvent_MemoryError) GetPropertyxMEM_DEVICE() (value uint16, err error) {
	retValue, err := instance.GetProperty("xMEM_DEVICE")
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
