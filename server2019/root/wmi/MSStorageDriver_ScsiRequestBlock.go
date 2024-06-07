// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSStorageDriver_ScsiRequestBlock struct
type MSStorageDriver_ScsiRequestBlock struct { 
	*cim.WmiInstance

	// CDB
	cdb []uint8

	// CDB Length
	cdbLength uint8

	// Data Buffer Pointer
	dataBuffer uint64

	// Data Transfer Length
	dataTransferLength uint32

	// Function
	function uint8

	// Internal Status
	internalStatus uint32

	// Length
	length uint16

	// LUN
	lun uint8

	// Next SRB Pointer
	nextSRB uint64

	// Original Request Pointer
	originalRequest uint64

	// Path ID
	pathID uint8

	// Queue Action
	queueAction uint8

	// Queue Tag
	queueTag uint8

	// Reserved (only available in Win64)
	reserved uint32

	// SCSI Status
	scsiStatus uint8

	// Sense Info Buffer Pointer
	senseInfoBuffer uint64

	// Sense Info Buffer Length
	senseInfoBufferLength uint8

	// SRB Extension Pointer
	srbExtension uint64

	// SRB Flags
	srbFlags uint32

	// SRB Status
	srbStatus uint8

	// Target ID
	targetID uint8

	// Time Out Value
	timeOutValue uint32
}

	func NewMSStorageDriver_ScsiRequestBlockEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_ScsiRequestBlock, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ScsiRequestBlock {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSStorageDriver_ScsiRequestBlockEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSStorageDriver_ScsiRequestBlock, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ScsiRequestBlock {
	WmiInstance: tmp,
	}
	return
	}
	

// Setcdb sets the value of cdb for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertycdb(value []uint8) (err error) { 
    return instance.SetProperty("cdb", (value))
}

// Getcdb gets the value of cdb for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertycdb()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("cdb")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetcdbLength sets the value of cdbLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertycdbLength(value uint8) (err error) { 
    return instance.SetProperty("cdbLength", (value))
}

// GetcdbLength gets the value of cdbLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertycdbLength()(value uint8, err error) { 
    retValue, err := instance.GetProperty("cdbLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetdataBuffer sets the value of dataBuffer for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertydataBuffer(value uint64) (err error) { 
    return instance.SetProperty("dataBuffer", (value))
}

// GetdataBuffer gets the value of dataBuffer for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertydataBuffer()(value uint64, err error) { 
    retValue, err := instance.GetProperty("dataBuffer")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetdataTransferLength sets the value of dataTransferLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertydataTransferLength(value uint32) (err error) { 
    return instance.SetProperty("dataTransferLength", (value))
}

// GetdataTransferLength gets the value of dataTransferLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertydataTransferLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("dataTransferLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// Setfunction sets the value of function for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyfunction(value uint8) (err error) { 
    return instance.SetProperty("function", (value))
}

// Getfunction gets the value of function for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyfunction()(value uint8, err error) { 
    retValue, err := instance.GetProperty("function")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetinternalStatus sets the value of internalStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyinternalStatus(value uint32) (err error) { 
    return instance.SetProperty("internalStatus", (value))
}

// GetinternalStatus gets the value of internalStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyinternalStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("internalStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// Setlength sets the value of length for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertylength(value uint16) (err error) { 
    return instance.SetProperty("length", (value))
}

// Getlength gets the value of length for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertylength()(value uint16, err error) { 
    retValue, err := instance.GetProperty("length")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// Setlun sets the value of lun for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertylun(value uint8) (err error) { 
    return instance.SetProperty("lun", (value))
}

// Getlun gets the value of lun for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertylun()(value uint8, err error) { 
    retValue, err := instance.GetProperty("lun")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetnextSRB sets the value of nextSRB for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertynextSRB(value uint64) (err error) { 
    return instance.SetProperty("nextSRB", (value))
}

// GetnextSRB gets the value of nextSRB for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertynextSRB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("nextSRB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetoriginalRequest sets the value of originalRequest for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyoriginalRequest(value uint64) (err error) { 
    return instance.SetProperty("originalRequest", (value))
}

// GetoriginalRequest gets the value of originalRequest for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyoriginalRequest()(value uint64, err error) { 
    retValue, err := instance.GetProperty("originalRequest")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetpathID sets the value of pathID for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertypathID(value uint8) (err error) { 
    return instance.SetProperty("pathID", (value))
}

// GetpathID gets the value of pathID for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertypathID()(value uint8, err error) { 
    retValue, err := instance.GetProperty("pathID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetqueueAction sets the value of queueAction for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyqueueAction(value uint8) (err error) { 
    return instance.SetProperty("queueAction", (value))
}

// GetqueueAction gets the value of queueAction for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyqueueAction()(value uint8, err error) { 
    retValue, err := instance.GetProperty("queueAction")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetqueueTag sets the value of queueTag for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyqueueTag(value uint8) (err error) { 
    return instance.SetProperty("queueTag", (value))
}

// GetqueueTag gets the value of queueTag for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyqueueTag()(value uint8, err error) { 
    retValue, err := instance.GetProperty("queueTag")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// Setreserved sets the value of reserved for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyreserved(value uint32) (err error) { 
    return instance.SetProperty("reserved", (value))
}

// Getreserved gets the value of reserved for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyreserved()(value uint32, err error) { 
    retValue, err := instance.GetProperty("reserved")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetscsiStatus sets the value of scsiStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertyscsiStatus(value uint8) (err error) { 
    return instance.SetProperty("scsiStatus", (value))
}

// GetscsiStatus gets the value of scsiStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertyscsiStatus()(value uint8, err error) { 
    retValue, err := instance.GetProperty("scsiStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetsenseInfoBuffer sets the value of senseInfoBuffer for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertysenseInfoBuffer(value uint64) (err error) { 
    return instance.SetProperty("senseInfoBuffer", (value))
}

// GetsenseInfoBuffer gets the value of senseInfoBuffer for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertysenseInfoBuffer()(value uint64, err error) { 
    retValue, err := instance.GetProperty("senseInfoBuffer")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetsenseInfoBufferLength sets the value of senseInfoBufferLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertysenseInfoBufferLength(value uint8) (err error) { 
    return instance.SetProperty("senseInfoBufferLength", (value))
}

// GetsenseInfoBufferLength gets the value of senseInfoBufferLength for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertysenseInfoBufferLength()(value uint8, err error) { 
    retValue, err := instance.GetProperty("senseInfoBufferLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetsrbExtension sets the value of srbExtension for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertysrbExtension(value uint64) (err error) { 
    return instance.SetProperty("srbExtension", (value))
}

// GetsrbExtension gets the value of srbExtension for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertysrbExtension()(value uint64, err error) { 
    retValue, err := instance.GetProperty("srbExtension")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetsrbFlags sets the value of srbFlags for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertysrbFlags(value uint32) (err error) { 
    return instance.SetProperty("srbFlags", (value))
}

// GetsrbFlags gets the value of srbFlags for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertysrbFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("srbFlags")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetsrbStatus sets the value of srbStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertysrbStatus(value uint8) (err error) { 
    return instance.SetProperty("srbStatus", (value))
}

// GetsrbStatus gets the value of srbStatus for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertysrbStatus()(value uint8, err error) { 
    retValue, err := instance.GetProperty("srbStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SettargetID sets the value of targetID for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertytargetID(value uint8) (err error) { 
    return instance.SetProperty("targetID", (value))
}

// GettargetID gets the value of targetID for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertytargetID()(value uint8, err error) { 
    retValue, err := instance.GetProperty("targetID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SettimeOutValue sets the value of timeOutValue for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) SetPropertytimeOutValue(value uint32) (err error) { 
    return instance.SetProperty("timeOutValue", (value))
}

// GettimeOutValue gets the value of timeOutValue for the instance
func (instance *MSStorageDriver_ScsiRequestBlock) GetPropertytimeOutValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("timeOutValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

