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

// MSMCAEvent_InvalidError struct
type MSMCAEvent_InvalidError struct { 
	*WMIEvent

	// 
	Active bool

	// 
	AdditionalErrors uint32

	// 
	Cpu uint32

	// 
	ErrorSeverity InvalidError_ErrorSeverity

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
	Type InvalidError_Type
}

	func NewMSMCAEvent_InvalidErrorEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_InvalidError, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_InvalidError {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewMSMCAEvent_InvalidErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSMCAEvent_InvalidError, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSMCAEvent_InvalidError {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_InvalidError) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyAdditionalErrors(value uint32) (err error) { 
    return instance.SetProperty("AdditionalErrors", (value))
}

// GetAdditionalErrors gets the value of AdditionalErrors for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyAdditionalErrors()(value uint32, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyCpu(value uint32) (err error) { 
    return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyCpu()(value uint32, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyErrorSeverity(value InvalidError_ErrorSeverity) (err error) { 
    return instance.SetProperty("ErrorSeverity", (value))
}

// GetErrorSeverity gets the value of ErrorSeverity for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyErrorSeverity()(value InvalidError_ErrorSeverity, err error) { 
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

    value = InvalidError_ErrorSeverity(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_InvalidError) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyInstanceName()(value string, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyLogToEventlog(value uint32) (err error) { 
    return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyLogToEventlog()(value uint32, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyRawRecord(value []uint8) (err error) { 
    return instance.SetProperty("RawRecord", (value))
}

// GetRawRecord gets the value of RawRecord for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyRawRecord()(value []uint8, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyRecordId(value uint64) (err error) { 
    return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyRecordId()(value uint64, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertySize(value uint32) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertySize()(value uint32, err error) { 
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
func (instance *MSMCAEvent_InvalidError) SetPropertyType(value InvalidError_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSMCAEvent_InvalidError) GetPropertyType()(value InvalidError_Type, err error) { 
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

    value = InvalidError_Type(valuetmp)

    return
}

