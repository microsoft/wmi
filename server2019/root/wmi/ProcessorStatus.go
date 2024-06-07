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

// ProcessorStatus struct
type ProcessorStatus struct { 
	*MSProcessorClass

	// 
	Active bool

	// 
	CurrentPerfState uint32

	// 
	InstanceName string

	// 
	LastRequestedThrottle uint32

	// 
	LastTransitionResult uint32

	// 
	LowestPerfState uint32

	// 
	PerfStates PerformanceStates

	// 
	ThrottleValue uint32

	// 
	UsingLegacyInterface uint32
}

	func NewProcessorStatusEx1(instance *cim.WmiInstance) (newInstance *ProcessorStatus, err error) {tmp, err := NewMSProcessorClassEx1(instance)
		
	if err != nil { return }
	newInstance = &ProcessorStatus {
	MSProcessorClass: tmp,
	}
	return
	}
	

	func NewProcessorStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProcessorStatus, err error) {tmp, err := NewMSProcessorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProcessorStatus {
	MSProcessorClass: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *ProcessorStatus) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorStatus) GetPropertyActive()(value bool, err error) { 
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

// SetCurrentPerfState sets the value of CurrentPerfState for the instance
func (instance *ProcessorStatus) SetPropertyCurrentPerfState(value uint32) (err error) { 
    return instance.SetProperty("CurrentPerfState", (value))
}

// GetCurrentPerfState gets the value of CurrentPerfState for the instance
func (instance *ProcessorStatus) GetPropertyCurrentPerfState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentPerfState")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorStatus) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorStatus) GetPropertyInstanceName()(value string, err error) { 
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

// SetLastRequestedThrottle sets the value of LastRequestedThrottle for the instance
func (instance *ProcessorStatus) SetPropertyLastRequestedThrottle(value uint32) (err error) { 
    return instance.SetProperty("LastRequestedThrottle", (value))
}

// GetLastRequestedThrottle gets the value of LastRequestedThrottle for the instance
func (instance *ProcessorStatus) GetPropertyLastRequestedThrottle()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LastRequestedThrottle")
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

// SetLastTransitionResult sets the value of LastTransitionResult for the instance
func (instance *ProcessorStatus) SetPropertyLastTransitionResult(value uint32) (err error) { 
    return instance.SetProperty("LastTransitionResult", (value))
}

// GetLastTransitionResult gets the value of LastTransitionResult for the instance
func (instance *ProcessorStatus) GetPropertyLastTransitionResult()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LastTransitionResult")
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

// SetLowestPerfState sets the value of LowestPerfState for the instance
func (instance *ProcessorStatus) SetPropertyLowestPerfState(value uint32) (err error) { 
    return instance.SetProperty("LowestPerfState", (value))
}

// GetLowestPerfState gets the value of LowestPerfState for the instance
func (instance *ProcessorStatus) GetPropertyLowestPerfState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LowestPerfState")
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

// SetPerfStates sets the value of PerfStates for the instance
func (instance *ProcessorStatus) SetPropertyPerfStates(value PerformanceStates) (err error) { 
    return instance.SetProperty("PerfStates", (value))
}

// GetPerfStates gets the value of PerfStates for the instance
func (instance *ProcessorStatus) GetPropertyPerfStates()(value PerformanceStates, err error) { 
    retValue, err := instance.GetProperty("PerfStates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(PerformanceStates); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " PerformanceStates is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PerformanceStates(valuetmp)

    return
}

// SetThrottleValue sets the value of ThrottleValue for the instance
func (instance *ProcessorStatus) SetPropertyThrottleValue(value uint32) (err error) { 
    return instance.SetProperty("ThrottleValue", (value))
}

// GetThrottleValue gets the value of ThrottleValue for the instance
func (instance *ProcessorStatus) GetPropertyThrottleValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ThrottleValue")
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

// SetUsingLegacyInterface sets the value of UsingLegacyInterface for the instance
func (instance *ProcessorStatus) SetPropertyUsingLegacyInterface(value uint32) (err error) { 
    return instance.SetProperty("UsingLegacyInterface", (value))
}

// GetUsingLegacyInterface gets the value of UsingLegacyInterface for the instance
func (instance *ProcessorStatus) GetPropertyUsingLegacyInterface()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UsingLegacyInterface")
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

