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

// PerformanceStates struct
type PerformanceStates struct { 
	*cim.WmiInstance

	// 
	Count uint32

	// 
	Current uint32

	// 
	State []PerformanceState

	// 
	TransitionFunction uint32

	// 
	TransitionLatency uint32
}

	func NewPerformanceStatesEx1(instance *cim.WmiInstance) (newInstance *PerformanceStates, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PerformanceStates {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPerformanceStatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PerformanceStates, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PerformanceStates {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCount sets the value of Count for the instance
func (instance *PerformanceStates) SetPropertyCount(value uint32) (err error) { 
    return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *PerformanceStates) GetPropertyCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Count")
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

// SetCurrent sets the value of Current for the instance
func (instance *PerformanceStates) SetPropertyCurrent(value uint32) (err error) { 
    return instance.SetProperty("Current", (value))
}

// GetCurrent gets the value of Current for the instance
func (instance *PerformanceStates) GetPropertyCurrent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Current")
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

// SetState sets the value of State for the instance
func (instance *PerformanceStates) SetPropertyState(value []PerformanceState) (err error) { 
    return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *PerformanceStates) GetPropertyState()(value []PerformanceState, err error) { 
    retValue, err := instance.GetProperty("State")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(PerformanceState); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " PerformanceState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, PerformanceState(valuetmp))
    }

    return
}

// SetTransitionFunction sets the value of TransitionFunction for the instance
func (instance *PerformanceStates) SetPropertyTransitionFunction(value uint32) (err error) { 
    return instance.SetProperty("TransitionFunction", (value))
}

// GetTransitionFunction gets the value of TransitionFunction for the instance
func (instance *PerformanceStates) GetPropertyTransitionFunction()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TransitionFunction")
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

// SetTransitionLatency sets the value of TransitionLatency for the instance
func (instance *PerformanceStates) SetPropertyTransitionLatency(value uint32) (err error) { 
    return instance.SetProperty("TransitionLatency", (value))
}

// GetTransitionLatency gets the value of TransitionLatency for the instance
func (instance *PerformanceStates) GetPropertyTransitionLatency()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TransitionLatency")
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

