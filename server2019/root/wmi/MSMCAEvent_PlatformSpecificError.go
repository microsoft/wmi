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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSMCAEvent_PlatformSpecificError struct
type MSMCAEvent_PlatformSpecificError struct { 
	*WMIEvent

	// 
	Active bool

	// 
	AdditionalErrors uint32

	// 
	Cpu uint32

	// 
	ErrorSeverity PlatformSpecificError_ErrorSeverity

	// 
	InstanceName string

	// 
	LogToEventlog uint32

	// 
	OEM_COMPONENT_ID []uint8

	// 
	PLATFORM_BUS_SPECIFIC_DATA uint64

	// 
	PLATFORM_ERROR_STATUS uint64

	// 
	PLATFORM_REQUESTOR_ID uint64

	// 
	PLATFORM_RESPONDER_ID uint64

	// 
	PLATFORM_TARGET_ID uint64

	// 
	RawRecord []uint8

	// 
	RecordId uint64

	// 
	Size uint32

	// 
	Type PlatformSpecificError_Type

	// 
	VALIDATION_BITS uint64
}

	func NewMSMCAEvent_PlatformSpecificErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_PlatformSpecificError, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_PlatformSpecificError {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewMSMCAEvent_PlatformSpecificErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSMCAEvent_PlatformSpecificError, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_PlatformSpecificError {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetAdditionalErrors sets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyAdditionalErrors(value uint32) (err error) { 
    return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyAdditionalErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AdditionalErrors")
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

// SetCpu sets the value of Cpu for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyCpu(value uint32) (err error) { 
    return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyCpu()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Cpu")
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

// SetErrorSeverity sets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyErrorSeverity(value PlatformSpecificError_ErrorSeverity) (err error) { 
    return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyErrorSeverity()(value PlatformSpecificError_ErrorSeverity, err error) { 
    retValue, err := instance.GetProperty("ErrorSeverity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PlatformSpecificError_ErrorSeverity(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetLogToEventlog sets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyLogToEventlog(value uint32) (err error) { 
    return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyLogToEventlog()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogToEventlog")
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

// SetOEM_COMPONENT_ID sets the value of OEM_COMPONENT_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyOEM_COMPONENT_ID(value []uint8) (err error) { 
    return instance.SetProperty("OEM_COMPONENT_ID", (value))
}

// GetOEM_COMPONENT_ID gets the value of OEM_COMPONENT_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyOEM_COMPONENT_ID()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("OEM_COMPONENT_ID")
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

// SetPLATFORM_BUS_SPECIFIC_DATA sets the value of PLATFORM_BUS_SPECIFIC_DATA for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyPLATFORM_BUS_SPECIFIC_DATA(value uint64) (err error) { 
    return instance.SetProperty("PLATFORM_BUS_SPECIFIC_DATA", (value))
}

// GetPLATFORM_BUS_SPECIFIC_DATA gets the value of PLATFORM_BUS_SPECIFIC_DATA for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyPLATFORM_BUS_SPECIFIC_DATA()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PLATFORM_BUS_SPECIFIC_DATA")
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

// SetPLATFORM_ERROR_STATUS sets the value of PLATFORM_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyPLATFORM_ERROR_STATUS(value uint64) (err error) { 
    return instance.SetProperty("PLATFORM_ERROR_STATUS", (value))
}

// GetPLATFORM_ERROR_STATUS gets the value of PLATFORM_ERROR_STATUS for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyPLATFORM_ERROR_STATUS()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PLATFORM_ERROR_STATUS")
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

// SetPLATFORM_REQUESTOR_ID sets the value of PLATFORM_REQUESTOR_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyPLATFORM_REQUESTOR_ID(value uint64) (err error) { 
    return instance.SetProperty("PLATFORM_REQUESTOR_ID", (value))
}

// GetPLATFORM_REQUESTOR_ID gets the value of PLATFORM_REQUESTOR_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyPLATFORM_REQUESTOR_ID()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PLATFORM_REQUESTOR_ID")
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

// SetPLATFORM_RESPONDER_ID sets the value of PLATFORM_RESPONDER_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyPLATFORM_RESPONDER_ID(value uint64) (err error) { 
    return instance.SetProperty("PLATFORM_RESPONDER_ID", (value))
}

// GetPLATFORM_RESPONDER_ID gets the value of PLATFORM_RESPONDER_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyPLATFORM_RESPONDER_ID()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PLATFORM_RESPONDER_ID")
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

// SetPLATFORM_TARGET_ID sets the value of PLATFORM_TARGET_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyPLATFORM_TARGET_ID(value uint64) (err error) { 
    return instance.SetProperty("PLATFORM_TARGET_ID", (value))
}

// GetPLATFORM_TARGET_ID gets the value of PLATFORM_TARGET_ID for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyPLATFORM_TARGET_ID()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PLATFORM_TARGET_ID")
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

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyRawRecord(value []uint8) (err error) { 
    return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyRawRecord()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("RawRecord")
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

// SetRecordId sets the value of RecordId for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyRecordId(value uint64) (err error) { 
    return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyRecordId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("RecordId")
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

// SetSize sets the value of Size for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertySize(value uint32) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertySize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Size")
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

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyType(value PlatformSpecificError_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyType()(value PlatformSpecificError_Type, err error) { 
    retValue, err := instance.GetProperty("Type")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PlatformSpecificError_Type(valuetmp)

    return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PlatformSpecificError) SetPropertyVALIDATION_BITS(value uint64) (err error) { 
    return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_PlatformSpecificError) GetPropertyVALIDATION_BITS()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VALIDATION_BITS")
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

