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

// IdleExitLatency struct
type IdleExitLatency struct { 
	*PowerEvents

	// 
	PlatformState uint32

	// 
	ProcessorState uint32

	// 
	ReturnLatency uint32

	// 
	TotalLatency uint32
}

	func NewIdleExitLatencyEx1(instance *cim.WmiInstance) (newInstance *IdleExitLatency, err error) {tmp, err := NewPowerEventsEx1(instance)
		
	if err != nil { return }
	newInstance = &IdleExitLatency {
	PowerEvents: tmp,
	}
	return
	}
	

	func NewIdleExitLatencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IdleExitLatency, err error) {tmp, err := NewPowerEventsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IdleExitLatency {
	PowerEvents: tmp,
	}
	return
	}
	

// SetPlatformState sets the value of PlatformState for the instance
func (instance *IdleExitLatency) SetPropertyPlatformState(value uint32) (err error) { 
    return instance.SetProperty("PlatformState", (value))
}

// GetPlatformState gets the value of PlatformState for the instance
func (instance *IdleExitLatency) GetPropertyPlatformState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PlatformState")
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

// SetProcessorState sets the value of ProcessorState for the instance
func (instance *IdleExitLatency) SetPropertyProcessorState(value uint32) (err error) { 
    return instance.SetProperty("ProcessorState", (value))
}

// GetProcessorState gets the value of ProcessorState for the instance
func (instance *IdleExitLatency) GetPropertyProcessorState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProcessorState")
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

// SetReturnLatency sets the value of ReturnLatency for the instance
func (instance *IdleExitLatency) SetPropertyReturnLatency(value uint32) (err error) { 
    return instance.SetProperty("ReturnLatency", (value))
}

// GetReturnLatency gets the value of ReturnLatency for the instance
func (instance *IdleExitLatency) GetPropertyReturnLatency()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReturnLatency")
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

// SetTotalLatency sets the value of TotalLatency for the instance
func (instance *IdleExitLatency) SetPropertyTotalLatency(value uint32) (err error) { 
    return instance.SetProperty("TotalLatency", (value))
}

// GetTotalLatency gets the value of TotalLatency for the instance
func (instance *IdleExitLatency) GetPropertyTotalLatency()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalLatency")
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

