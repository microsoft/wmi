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

// MSMCAEvent_SMBIOSError struct
type MSMCAEvent_SMBIOSError struct { 
	*WMIEvent

	// 
	Active bool

	// 
	AdditionalErrors uint32

	// 
	Cpu uint32

	// 
	ErrorSeverity SMBIOSError_ErrorSeverity

	// 
	InstanceName string

	// 
	LogToEventlog uint32

	// 
	RawRecord []uint8

	// 
	RecordId uint64

	// 
	Size uint32

	// 
	SMBIOS_EVENT_TYPE uint8

	// 
	Type SMBIOSError_Type

	// 
	VALIDATION_BITS uint64
}

	func NewMSMCAEvent_SMBIOSErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_SMBIOSError, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_SMBIOSError {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewMSMCAEvent_SMBIOSErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSMCAEvent_SMBIOSError, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_SMBIOSError {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertyAdditionalErrors(value uint32) (err error) { 
    return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyAdditionalErrors()(value uint32, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertyCpu(value uint32) (err error) { 
    return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyCpu()(value uint32, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertyErrorSeverity(value SMBIOSError_ErrorSeverity) (err error) { 
    return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyErrorSeverity()(value SMBIOSError_ErrorSeverity, err error) { 
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

    value = SMBIOSError_ErrorSeverity(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyInstanceName()(value string, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertyLogToEventlog(value uint32) (err error) { 
    return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyLogToEventlog()(value uint32, err error) { 
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

// SetRawRecord sets the value of RawRecord for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertyRawRecord(value []uint8) (err error) { 
    return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyRawRecord()(value []uint8, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertyRecordId(value uint64) (err error) { 
    return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyRecordId()(value uint64, err error) { 
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
func (instance *MSMCAEvent_SMBIOSError) SetPropertySize(value uint32) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertySize()(value uint32, err error) { 
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

// SetSMBIOS_EVENT_TYPE sets the value of SMBIOS_EVENT_TYPE for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertySMBIOS_EVENT_TYPE(value uint8) (err error) { 
    return instance.SetProperty("SMBIOS_EVENT_TYPE", (value))
}

// GetSMBIOS_EVENT_TYPE gets the value of SMBIOS_EVENT_TYPE for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertySMBIOS_EVENT_TYPE()(value uint8, err error) { 
    retValue, err := instance.GetProperty("SMBIOS_EVENT_TYPE")
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

// SetType sets the value of Type for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertyType(value SMBIOSError_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyType()(value SMBIOSError_Type, err error) { 
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

    value = SMBIOSError_Type(valuetmp)

    return
}

// SetVALIDATION_BITS sets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_SMBIOSError) SetPropertyVALIDATION_BITS(value uint64) (err error) { 
    return instance.SetProperty("VALIDATION_BITS", (value))
}

// GetVALIDATION_BITS gets the value of VALIDATION_BITS for the instance
func (instance *MSMCAEvent_SMBIOSError) GetPropertyVALIDATION_BITS()(value uint64, err error) { 
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

