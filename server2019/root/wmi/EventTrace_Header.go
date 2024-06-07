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

// EventTrace_Header struct
type EventTrace_Header struct { 
	*EventTraceEvent

	// 
	BootTime uint64

	// 
	BufferSize uint32

	// 
	BuffersLost uint32

	// 
	BuffersWritten uint32

	// 
	CPUSpeed uint32

	// 
	EndTime uint64

	// 
	EventsLost uint32

	// 
	LogFileMode uint32

	// 
	LogFileName uint32

	// 
	LogFileNameString string

	// 
	LoggerName uint32

	// 
	MaxFileSize uint32

	// 
	NumberOfProcessors uint32

	// 
	PerfFreq uint64

	// 
	PointerSize uint32

	// 
	ProviderVersion uint32

	// 
	ReservedFlags uint32

	// 
	SessionNameString string

	// 
	StartBuffers uint32

	// 
	StartTime uint64

	// 
	TimerResolution uint32

	// 
	TimeZoneInformation []uint8

	// 
	Version uint32
}

	func NewEventTrace_HeaderEx1(instance *cim.WmiInstance) (newInstance *EventTrace_Header, err error) {tmp, err := NewEventTraceEventEx1(instance)
		
	if err != nil { return }
	newInstance = &EventTrace_Header {
	EventTraceEvent: tmp,
	}
	return
	}
	

	func NewEventTrace_HeaderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EventTrace_Header, err error) {tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EventTrace_Header {
	EventTraceEvent: tmp,
	}
	return
	}
	

// SetBootTime sets the value of BootTime for the instance
func (instance *EventTrace_Header) SetPropertyBootTime(value uint64) (err error) { 
    return instance.SetProperty("BootTime", (value))
}

// GetBootTime gets the value of BootTime for the instance
func (instance *EventTrace_Header) GetPropertyBootTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BootTime")
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

// SetBufferSize sets the value of BufferSize for the instance
func (instance *EventTrace_Header) SetPropertyBufferSize(value uint32) (err error) { 
    return instance.SetProperty("BufferSize", (value))
}

// GetBufferSize gets the value of BufferSize for the instance
func (instance *EventTrace_Header) GetPropertyBufferSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BufferSize")
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

// SetBuffersLost sets the value of BuffersLost for the instance
func (instance *EventTrace_Header) SetPropertyBuffersLost(value uint32) (err error) { 
    return instance.SetProperty("BuffersLost", (value))
}

// GetBuffersLost gets the value of BuffersLost for the instance
func (instance *EventTrace_Header) GetPropertyBuffersLost()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BuffersLost")
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

// SetBuffersWritten sets the value of BuffersWritten for the instance
func (instance *EventTrace_Header) SetPropertyBuffersWritten(value uint32) (err error) { 
    return instance.SetProperty("BuffersWritten", (value))
}

// GetBuffersWritten gets the value of BuffersWritten for the instance
func (instance *EventTrace_Header) GetPropertyBuffersWritten()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BuffersWritten")
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

// SetCPUSpeed sets the value of CPUSpeed for the instance
func (instance *EventTrace_Header) SetPropertyCPUSpeed(value uint32) (err error) { 
    return instance.SetProperty("CPUSpeed", (value))
}

// GetCPUSpeed gets the value of CPUSpeed for the instance
func (instance *EventTrace_Header) GetPropertyCPUSpeed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CPUSpeed")
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

// SetEndTime sets the value of EndTime for the instance
func (instance *EventTrace_Header) SetPropertyEndTime(value uint64) (err error) { 
    return instance.SetProperty("EndTime", (value))
}

// GetEndTime gets the value of EndTime for the instance
func (instance *EventTrace_Header) GetPropertyEndTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("EndTime")
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

// SetEventsLost sets the value of EventsLost for the instance
func (instance *EventTrace_Header) SetPropertyEventsLost(value uint32) (err error) { 
    return instance.SetProperty("EventsLost", (value))
}

// GetEventsLost gets the value of EventsLost for the instance
func (instance *EventTrace_Header) GetPropertyEventsLost()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EventsLost")
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

// SetLogFileMode sets the value of LogFileMode for the instance
func (instance *EventTrace_Header) SetPropertyLogFileMode(value uint32) (err error) { 
    return instance.SetProperty("LogFileMode", (value))
}

// GetLogFileMode gets the value of LogFileMode for the instance
func (instance *EventTrace_Header) GetPropertyLogFileMode()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogFileMode")
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

// SetLogFileName sets the value of LogFileName for the instance
func (instance *EventTrace_Header) SetPropertyLogFileName(value uint32) (err error) { 
    return instance.SetProperty("LogFileName", (value))
}

// GetLogFileName gets the value of LogFileName for the instance
func (instance *EventTrace_Header) GetPropertyLogFileName()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogFileName")
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

// SetLogFileNameString sets the value of LogFileNameString for the instance
func (instance *EventTrace_Header) SetPropertyLogFileNameString(value string) (err error) { 
    return instance.SetProperty("LogFileNameString", (value))
}

// GetLogFileNameString gets the value of LogFileNameString for the instance
func (instance *EventTrace_Header) GetPropertyLogFileNameString()(value string, err error) { 
    retValue, err := instance.GetProperty("LogFileNameString")
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

// SetLoggerName sets the value of LoggerName for the instance
func (instance *EventTrace_Header) SetPropertyLoggerName(value uint32) (err error) { 
    return instance.SetProperty("LoggerName", (value))
}

// GetLoggerName gets the value of LoggerName for the instance
func (instance *EventTrace_Header) GetPropertyLoggerName()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoggerName")
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

// SetMaxFileSize sets the value of MaxFileSize for the instance
func (instance *EventTrace_Header) SetPropertyMaxFileSize(value uint32) (err error) { 
    return instance.SetProperty("MaxFileSize", (value))
}

// GetMaxFileSize gets the value of MaxFileSize for the instance
func (instance *EventTrace_Header) GetPropertyMaxFileSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxFileSize")
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

// SetNumberOfProcessors sets the value of NumberOfProcessors for the instance
func (instance *EventTrace_Header) SetPropertyNumberOfProcessors(value uint32) (err error) { 
    return instance.SetProperty("NumberOfProcessors", (value))
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *EventTrace_Header) GetPropertyNumberOfProcessors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfProcessors")
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

// SetPerfFreq sets the value of PerfFreq for the instance
func (instance *EventTrace_Header) SetPropertyPerfFreq(value uint64) (err error) { 
    return instance.SetProperty("PerfFreq", (value))
}

// GetPerfFreq gets the value of PerfFreq for the instance
func (instance *EventTrace_Header) GetPropertyPerfFreq()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PerfFreq")
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

// SetPointerSize sets the value of PointerSize for the instance
func (instance *EventTrace_Header) SetPropertyPointerSize(value uint32) (err error) { 
    return instance.SetProperty("PointerSize", (value))
}

// GetPointerSize gets the value of PointerSize for the instance
func (instance *EventTrace_Header) GetPropertyPointerSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PointerSize")
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

// SetProviderVersion sets the value of ProviderVersion for the instance
func (instance *EventTrace_Header) SetPropertyProviderVersion(value uint32) (err error) { 
    return instance.SetProperty("ProviderVersion", (value))
}

// GetProviderVersion gets the value of ProviderVersion for the instance
func (instance *EventTrace_Header) GetPropertyProviderVersion()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProviderVersion")
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

// SetReservedFlags sets the value of ReservedFlags for the instance
func (instance *EventTrace_Header) SetPropertyReservedFlags(value uint32) (err error) { 
    return instance.SetProperty("ReservedFlags", (value))
}

// GetReservedFlags gets the value of ReservedFlags for the instance
func (instance *EventTrace_Header) GetPropertyReservedFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReservedFlags")
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

// SetSessionNameString sets the value of SessionNameString for the instance
func (instance *EventTrace_Header) SetPropertySessionNameString(value string) (err error) { 
    return instance.SetProperty("SessionNameString", (value))
}

// GetSessionNameString gets the value of SessionNameString for the instance
func (instance *EventTrace_Header) GetPropertySessionNameString()(value string, err error) { 
    retValue, err := instance.GetProperty("SessionNameString")
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

// SetStartBuffers sets the value of StartBuffers for the instance
func (instance *EventTrace_Header) SetPropertyStartBuffers(value uint32) (err error) { 
    return instance.SetProperty("StartBuffers", (value))
}

// GetStartBuffers gets the value of StartBuffers for the instance
func (instance *EventTrace_Header) GetPropertyStartBuffers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StartBuffers")
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

// SetStartTime sets the value of StartTime for the instance
func (instance *EventTrace_Header) SetPropertyStartTime(value uint64) (err error) { 
    return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *EventTrace_Header) GetPropertyStartTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("StartTime")
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

// SetTimerResolution sets the value of TimerResolution for the instance
func (instance *EventTrace_Header) SetPropertyTimerResolution(value uint32) (err error) { 
    return instance.SetProperty("TimerResolution", (value))
}

// GetTimerResolution gets the value of TimerResolution for the instance
func (instance *EventTrace_Header) GetPropertyTimerResolution()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TimerResolution")
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

// SetTimeZoneInformation sets the value of TimeZoneInformation for the instance
func (instance *EventTrace_Header) SetPropertyTimeZoneInformation(value []uint8) (err error) { 
    return instance.SetProperty("TimeZoneInformation", (value))
}

// GetTimeZoneInformation gets the value of TimeZoneInformation for the instance
func (instance *EventTrace_Header) GetPropertyTimeZoneInformation()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("TimeZoneInformation")
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

// SetVersion sets the value of Version for the instance
func (instance *EventTrace_Header) SetPropertyVersion(value uint32) (err error) { 
    return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *EventTrace_Header) GetPropertyVersion()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Version")
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

