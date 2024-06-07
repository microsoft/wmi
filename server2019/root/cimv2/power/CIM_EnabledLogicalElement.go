// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
standardcimv2 "github.com/microsoft/wmi/server2019/root/standardcimv2"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_EnabledLogicalElement struct
type CIM_EnabledLogicalElement struct { 
	*CIM_LogicalElement

	// 
	AvailableRequestedStates []uint16

	// 
	EnabledDefault uint16

	// 
	EnabledState uint16

	// 
	OtherEnabledState string

	// 
	RequestedState uint16

	// 
	TimeOfLastStateChange string

	// 
	TransitioningToState uint16
}

	func NewCIM_EnabledLogicalElementEx1(instance *cim.WmiInstance) (newInstance *CIM_EnabledLogicalElement, err error) {tmp, err := NewCIM_LogicalElementEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_EnabledLogicalElement {
	CIM_LogicalElement: tmp,
	}
	return
	}
	

	func NewCIM_EnabledLogicalElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_EnabledLogicalElement, err error) {tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_EnabledLogicalElement {
	CIM_LogicalElement: tmp,
	}
	return
	}
	

// SetAvailableRequestedStates sets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyAvailableRequestedStates(value []uint16) (err error) { 
    return instance.SetProperty("AvailableRequestedStates", (value))
}

// GetAvailableRequestedStates gets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyAvailableRequestedStates()(value []uint16, err error) { 
    retValue, err := instance.GetProperty("AvailableRequestedStates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint16); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint16(valuetmp))
    }

    return
}

// SetEnabledDefault sets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledDefault(value uint16) (err error) { 
    return instance.SetProperty("EnabledDefault", (value))
}

// GetEnabledDefault gets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledDefault()(value uint16, err error) { 
    retValue, err := instance.GetProperty("EnabledDefault")
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

// SetEnabledState sets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledState(value uint16) (err error) { 
    return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledState()(value uint16, err error) { 
    retValue, err := instance.GetProperty("EnabledState")
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

// SetOtherEnabledState sets the value of OtherEnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyOtherEnabledState(value string) (err error) { 
    return instance.SetProperty("OtherEnabledState", (value))
}

// GetOtherEnabledState gets the value of OtherEnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyOtherEnabledState()(value string, err error) { 
    retValue, err := instance.GetProperty("OtherEnabledState")
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

// SetRequestedState sets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyRequestedState(value uint16) (err error) { 
    return instance.SetProperty("RequestedState", (value))
}

// GetRequestedState gets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyRequestedState()(value uint16, err error) { 
    retValue, err := instance.GetProperty("RequestedState")
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

// SetTimeOfLastStateChange sets the value of TimeOfLastStateChange for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyTimeOfLastStateChange(value string) (err error) { 
    return instance.SetProperty("TimeOfLastStateChange", (value))
}

// GetTimeOfLastStateChange gets the value of TimeOfLastStateChange for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyTimeOfLastStateChange()(value string, err error) { 
    retValue, err := instance.GetProperty("TimeOfLastStateChange")
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

// SetTransitioningToState sets the value of TransitioningToState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyTransitioningToState(value uint16) (err error) { 
    return instance.SetProperty("TransitioningToState", (value))
}

// GetTransitioningToState gets the value of TransitioningToState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyTransitioningToState()(value uint16, err error) { 
    retValue, err := instance.GetProperty("TransitioningToState")
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

// 

// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_EnabledLogicalElement) RequestStateChange( /* IN */ RequestedState uint16,
 /* OUT */ Job standardcimv2.CIM_ConcreteJob,
 /* OPTIONAL IN */ TimeoutPeriod string,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RequestStateChange", Action, PercentComplete, Timeout , RequestedState, TimeoutPeriod)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


