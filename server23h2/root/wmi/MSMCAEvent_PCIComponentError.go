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

// MSMCAEvent_PCIComponentError struct
type MSMCAEvent_PCIComponentError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	//
	Cpu uint32

	//
	ErrorSeverity PCIComponentError_ErrorSeverity

	//
	InstanceName string

	//
	LogToEventlog uint32

	//
	PCI_COMP_ERROR_STATUS uint64

	//
	PCI_COMP_INFO_BusNumber uint8

	//
	PCI_COMP_INFO_ClassCodeBaseClass uint8

	//
	PCI_COMP_INFO_ClassCodeInterface uint8

	//
	PCI_COMP_INFO_ClassCodeSubClass uint8

	//
	PCI_COMP_INFO_DeviceId uint16

	//
	PCI_COMP_INFO_DeviceNumber uint8

	//
	PCI_COMP_INFO_FunctionNumber uint8

	//
	PCI_COMP_INFO_SegmentNumber uint8

	//
	PCI_COMP_INFO_VendorId uint16

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	Size uint32

	//
	Type PCIComponentError_Type

	//
	VALIDATION_BITS uint64
}

func NewMSMCAEvent_PCIComponentErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_PCIComponentError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_PCIComponentError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_PCIComponentErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_PCIComponentError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_PCIComponentError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyAdditionalErrors() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyErrorSeverity(value PCIComponentError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyErrorSeverity() (value PCIComponentError_ErrorSeverity, err error) {
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

	value = PCIComponentError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyLogToEventlog() (value uint32, err error) {
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

// SetPCI_COMP_ERROR_STATUS sets the value of PCI_COMP_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_ERROR_STATUS(value uint64) (err error) {
	return instance.SetProperty("PCI_COMP_ERROR_STATUS", (value))
}

// GetPCI_COMP_ERROR_STATUS gets the value of PCI_COMP_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_ERROR_STATUS() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_ERROR_STATUS")
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

// SetPCI_COMP_INFO_BusNumber sets the value of PCI_COMP_INFO_BusNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_BusNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_BusNumber", (value))
}

// GetPCI_COMP_INFO_BusNumber gets the value of PCI_COMP_INFO_BusNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_BusNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_BusNumber")
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

// SetPCI_COMP_INFO_ClassCodeBaseClass sets the value of PCI_COMP_INFO_ClassCodeBaseClass for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_ClassCodeBaseClass(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_ClassCodeBaseClass", (value))
}

// GetPCI_COMP_INFO_ClassCodeBaseClass gets the value of PCI_COMP_INFO_ClassCodeBaseClass for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_ClassCodeBaseClass() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_ClassCodeBaseClass")
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

// SetPCI_COMP_INFO_ClassCodeInterface sets the value of PCI_COMP_INFO_ClassCodeInterface for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_ClassCodeInterface(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_ClassCodeInterface", (value))
}

// GetPCI_COMP_INFO_ClassCodeInterface gets the value of PCI_COMP_INFO_ClassCodeInterface for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_ClassCodeInterface() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_ClassCodeInterface")
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

// SetPCI_COMP_INFO_ClassCodeSubClass sets the value of PCI_COMP_INFO_ClassCodeSubClass for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_ClassCodeSubClass(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_ClassCodeSubClass", (value))
}

// GetPCI_COMP_INFO_ClassCodeSubClass gets the value of PCI_COMP_INFO_ClassCodeSubClass for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_ClassCodeSubClass() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_ClassCodeSubClass")
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

// SetPCI_COMP_INFO_DeviceId sets the value of PCI_COMP_INFO_DeviceId for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_DeviceId(value uint16) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_DeviceId", (value))
}

// GetPCI_COMP_INFO_DeviceId gets the value of PCI_COMP_INFO_DeviceId for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_DeviceId() (value uint16, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_DeviceId")
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

// SetPCI_COMP_INFO_DeviceNumber sets the value of PCI_COMP_INFO_DeviceNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_DeviceNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_DeviceNumber", (value))
}

// GetPCI_COMP_INFO_DeviceNumber gets the value of PCI_COMP_INFO_DeviceNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_DeviceNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_DeviceNumber")
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

// SetPCI_COMP_INFO_FunctionNumber sets the value of PCI_COMP_INFO_FunctionNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_FunctionNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_FunctionNumber", (value))
}

// GetPCI_COMP_INFO_FunctionNumber gets the value of PCI_COMP_INFO_FunctionNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_FunctionNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_FunctionNumber")
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

// SetPCI_COMP_INFO_SegmentNumber sets the value of PCI_COMP_INFO_SegmentNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_SegmentNumber(value uint8) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_SegmentNumber", (value))
}

// GetPCI_COMP_INFO_SegmentNumber gets the value of PCI_COMP_INFO_SegmentNumber for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_SegmentNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_SegmentNumber")
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

// SetPCI_COMP_INFO_VendorId sets the value of PCI_COMP_INFO_VendorId for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyPCI_COMP_INFO_VendorId(value uint16) (err error) {
	return instance.SetProperty("PCI_COMP_INFO_VendorId", (value))
}

// GetPCI_COMP_INFO_VendorId gets the value of PCI_COMP_INFO_VendorId for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyPCI_COMP_INFO_VendorId() (value uint16, err error) {
	retValue, err := instance.GetProperty("PCI_COMP_INFO_VendorId")
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyRecordId() (value uint64, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertySize() (value uint32, err error) {
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
func (instance *MSMCAEvent_PCIComponentError) SetPropertyType(value PCIComponentError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyType() (value PCIComponentError_Type, err error) {
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

	value = PCIComponentError_Type(valuetmp)

	return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PCIComponentError) SetPropertyVALIDATION_BITS(value uint64) (err error) {
	return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PCIComponentError) GetPropertyVALIDATION_BITS() (value uint64, err error) {
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
