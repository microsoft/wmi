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

// MSMCAEvent_CPUError struct
type MSMCAEvent_CPUError struct {
	*WMIEvent

	//
	Active bool

	//
	AdditionalErrors uint32

	//
	BusSev uint32

	//
	BusType uint32

	//
	CacheMesi uint32

	//
	CacheOp uint32

	//
	Cpu uint32

	//
	ErrorSeverity CPUError_ErrorSeverity

	//
	InstanceName string

	//
	Level uint32

	//
	LogToEventlog uint32

	//
	MajorErrorType CPUError_MajorErrorType

	//
	MSArrayId uint32

	//
	MSIndex uint32

	//
	MSOp uint32

	//
	MSSid uint32

	//
	RawRecord []uint8

	//
	RecordId uint64

	//
	RegFileId uint32

	//
	RegFileOp uint32

	//
	Size uint32

	//
	TLBOp uint32

	//
	Type CPUError_Type
}

func NewMSMCAEvent_CPUErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_CPUError, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_CPUError{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_CPUErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_CPUError, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_CPUError{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyActive() (value bool, err error) {
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
func (instance *MSMCAEvent_CPUError) SetPropertyAdditionalErrors(value uint32) (err error) {
	return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyAdditionalErrors() (value uint32, err error) {
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

// SetBusSev sets the value of BusSev for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyBusSev(value uint32) (err error) {
	return instance.SetProperty("BusSev", (value))
}

// GetBusSev gets the value of BusSev for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyBusSev() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusSev")
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

// SetBusType sets the value of BusType for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyBusType(value uint32) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyBusType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusType")
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

// SetCacheMesi sets the value of CacheMesi for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyCacheMesi(value uint32) (err error) {
	return instance.SetProperty("CacheMesi", (value))
}

// GetCacheMesi gets the value of CacheMesi for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyCacheMesi() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheMesi")
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

// SetCacheOp sets the value of CacheOp for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyCacheOp(value uint32) (err error) {
	return instance.SetProperty("CacheOp", (value))
}

// GetCacheOp gets the value of CacheOp for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyCacheOp() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheOp")
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
func (instance *MSMCAEvent_CPUError) SetPropertyCpu(value uint32) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyCpu() (value uint32, err error) {
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
func (instance *MSMCAEvent_CPUError) SetPropertyErrorSeverity(value CPUError_ErrorSeverity) (err error) {
	return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyErrorSeverity() (value CPUError_ErrorSeverity, err error) {
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

	value = CPUError_ErrorSeverity(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyInstanceName() (value string, err error) {
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

// SetLevel sets the value of Level for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyLevel(value uint32) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Level")
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

// SetLogToEventlog sets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyLogToEventlog() (value uint32, err error) {
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

// SetMajorErrorType sets the value of MajorErrorType for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyMajorErrorType(value CPUError_MajorErrorType) (err error) {
	return instance.SetProperty("MajorErrorType", (value))
}

// GetMajorErrorType gets the value of MajorErrorType for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyMajorErrorType() (value CPUError_MajorErrorType, err error) {
	retValue, err := instance.GetProperty("MajorErrorType")
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

	value = CPUError_MajorErrorType(valuetmp)

	return
}

// SetMSArrayId sets the value of MSArrayId for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyMSArrayId(value uint32) (err error) {
	return instance.SetProperty("MSArrayId", (value))
}

// GetMSArrayId gets the value of MSArrayId for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyMSArrayId() (value uint32, err error) {
	retValue, err := instance.GetProperty("MSArrayId")
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

// SetMSIndex sets the value of MSIndex for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyMSIndex(value uint32) (err error) {
	return instance.SetProperty("MSIndex", (value))
}

// GetMSIndex gets the value of MSIndex for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyMSIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("MSIndex")
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

// SetMSOp sets the value of MSOp for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyMSOp(value uint32) (err error) {
	return instance.SetProperty("MSOp", (value))
}

// GetMSOp gets the value of MSOp for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyMSOp() (value uint32, err error) {
	retValue, err := instance.GetProperty("MSOp")
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

// SetMSSid sets the value of MSSid for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyMSSid(value uint32) (err error) {
	return instance.SetProperty("MSSid", (value))
}

// GetMSSid gets the value of MSSid for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyMSSid() (value uint32, err error) {
	retValue, err := instance.GetProperty("MSSid")
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
func (instance *MSMCAEvent_CPUError) SetPropertyRawRecord(value []uint8) (err error) {
	return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyRawRecord() (value []uint8, err error) {
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
func (instance *MSMCAEvent_CPUError) SetPropertyRecordId(value uint64) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyRecordId() (value uint64, err error) {
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

// SetRegFileId sets the value of RegFileId for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyRegFileId(value uint32) (err error) {
	return instance.SetProperty("RegFileId", (value))
}

// GetRegFileId gets the value of RegFileId for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyRegFileId() (value uint32, err error) {
	retValue, err := instance.GetProperty("RegFileId")
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

// SetRegFileOp sets the value of RegFileOp for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyRegFileOp(value uint32) (err error) {
	return instance.SetProperty("RegFileOp", (value))
}

// GetRegFileOp gets the value of RegFileOp for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyRegFileOp() (value uint32, err error) {
	retValue, err := instance.GetProperty("RegFileOp")
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

// SetSize sets the value of Size for the instance
func (instance *MSMCAEvent_CPUError) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_CPUError) GetPropertySize() (value uint32, err error) {
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

// SetTLBOp sets the value of TLBOp for the instance
func (instance *MSMCAEvent_CPUError) SetPropertyTLBOp(value uint32) (err error) {
	return instance.SetProperty("TLBOp", (value))
}

// GetTLBOp gets the value of TLBOp for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyTLBOp() (value uint32, err error) {
	retValue, err := instance.GetProperty("TLBOp")
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
func (instance *MSMCAEvent_CPUError) SetPropertyType(value CPUError_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_CPUError) GetPropertyType() (value CPUError_Type, err error) {
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

	value = CPUError_Type(valuetmp)

	return
}
