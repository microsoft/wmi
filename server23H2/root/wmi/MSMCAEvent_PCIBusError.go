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

// MSMCAEvent_PCIBusError struct
type MSMCAEvent_PCIBusError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	//
	Cpu uint32

	//
	ErrorSeverity PCIBusError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	PCI_BUS_ADDRESS uint64

	//
	PCI_BUS_CMD uint64

	//
	PCI_BUS_DATA uint64

	//
	PCI_BUS_ERROR_STATUS uint64

	//
	PCI_BUS_ERROR_TYPE uint16

	//
	PCI_BUS_ID_BusNumber uint8

	//
	PCI_BUS_ID_SegmentNumber uint8

	//
	PCI_BUS_REQUESTOR_ID uint64

	//
	PCI_BUS_RESPONDER_ID uint64

	//
	PCI_BUS_TARGET_ID uint64

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	Size uint32

	//
	Type PCIBusError_Type

	//
	VALIDATION_BITS uint64
}

func NewMSMCAEvent_PCIBusErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_PCIBusError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_PCIBusError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_PCIBusErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_PCIBusError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_PCIBusError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyAdditionalErrors() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyErrorSeverity(value PCIBusError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyErrorSeverity() (value PCIBusError_ErrorSeverity, err error) {
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

	value = PCIBusError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyLogToEventlog() (value uint32, err error) {
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

// SetPCI_BUS_ADDRESS sets the value of PCI_BUS_ADDRESS for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_ADDRESS(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_ADDRESS", (value))
}

// GetPCI_BUS_ADDRESS gets the value of PCI_BUS_ADDRESS for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_ADDRESS() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_ADDRESS")
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

// SetPCI_BUS_CMD sets the value of PCI_BUS_CMD for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_CMD(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_CMD", (value))
}

// GetPCI_BUS_CMD gets the value of PCI_BUS_CMD for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_CMD() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_CMD")
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

// SetPCI_BUS_DATA sets the value of PCI_BUS_DATA for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_DATA(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_DATA", (value))
}

// GetPCI_BUS_DATA gets the value of PCI_BUS_DATA for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_DATA() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_DATA")
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

// SetPCI_BUS_ERROR_STATUS sets the value of PCI_BUS_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_ERROR_STATUS(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_ERROR_STATUS", (value))
}

// GetPCI_BUS_ERROR_STATUS gets the value of PCI_BUS_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_ERROR_STATUS() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_ERROR_STATUS")
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

// SetPCI_BUS_ERROR_TYPE sets the value of PCI_BUS_ERROR_TYPE for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_ERROR_TYPE(value uint16) (err error) {
	return instance.SetProperty("PCI_BUS_ERROR_TYPE", (value))
}

// GetPCI_BUS_ERROR_TYPE gets the value of PCI_BUS_ERROR_TYPE for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_ERROR_TYPE() (value uint16, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_ERROR_TYPE")
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

// SetPCI_BUS_ID_BusNumber sets the value of PCI_BUS_ID_BusNumber for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_ID_BusNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_BUS_ID_BusNumber", (value))
}

// GetPCI_BUS_ID_BusNumber gets the value of PCI_BUS_ID_BusNumber for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_ID_BusNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_ID_BusNumber")
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

// SetPCI_BUS_ID_SegmentNumber sets the value of PCI_BUS_ID_SegmentNumber for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_ID_SegmentNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_BUS_ID_SegmentNumber", (value))
}

// GetPCI_BUS_ID_SegmentNumber gets the value of PCI_BUS_ID_SegmentNumber for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_ID_SegmentNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_ID_SegmentNumber")
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

// SetPCI_BUS_REQUESTOR_ID sets the value of PCI_BUS_REQUESTOR_ID for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_REQUESTOR_ID(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_REQUESTOR_ID", (value))
}

// GetPCI_BUS_REQUESTOR_ID gets the value of PCI_BUS_REQUESTOR_ID for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_REQUESTOR_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_REQUESTOR_ID")
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

// SetPCI_BUS_RESPONDER_ID sets the value of PCI_BUS_RESPONDER_ID for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_RESPONDER_ID(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_RESPONDER_ID", (value))
}

// GetPCI_BUS_RESPONDER_ID gets the value of PCI_BUS_RESPONDER_ID for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_RESPONDER_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_RESPONDER_ID")
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

// SetPCI_BUS_TARGET_ID sets the value of PCI_BUS_TARGET_ID for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyPCI_BUS_TARGET_ID(value uint64) (err error) {
	return instance.SetProperty("PCI_BUS_TARGET_ID", (value))
}

// GetPCI_BUS_TARGET_ID gets the value of PCI_BUS_TARGET_ID for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyPCI_BUS_TARGET_ID() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_BUS_TARGET_ID")
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

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyRecordId() (value uint64, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertySize() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIBusError) SetPropertyType(value PCIBusError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyType() (value PCIBusError_Type, err error) {
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

	value = PCIBusError_Type(valuetmp)

	return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PCIBusError) SetPropertyVALIDATION_BITS(value uint64) (err error) {
	return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PCIBusError) GetPropertyVALIDATION_BITS() (value uint64, err error) {
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
