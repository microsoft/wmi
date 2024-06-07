// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_EnabledLogicalElement struct
type CIM_EnabledLogicalElement struct { 
	*CIM_LogicalElement

	// AvailableRequestedStates indicates the possible values for the RequestedState parameter of the method RequestStateChange, used to initiate a state change. The values listed shall be a subset of the values contained in the RequestedStatesSupported property of the associated instance of CIM_EnabledLogicalElementCapabilities where the values selected are a function of the current state of the CIM_EnabledLogicalElement. This property may be non-null if an implementation is able to advertise the set of possible values as a function of the current state. This property shall be null if an implementation is unable to determine the set of possible values as a function of the current state.
	AvailableRequestedStates []EnabledLogicalElement_AvailableRequestedStates

	// An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
	EnabledDefault EnabledLogicalElement_EnabledDefault

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled. The following text briefly summarizes the various enabled and disabled states: 
	///Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests. 
	///Disabled (3) indicates that the element will not execute commands and will drop any new requests. 
	///Shutting Down (4) indicates that the element is in the process of going to a Disabled state. 
	///Not Applicable (5) indicates the element does not support being enabled or disabled. 
	///Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests. 
	///Test (7) indicates that the element is in a test state. 
	///Deferred (8) indicates that the element might be completing commands, but will queue any new requests. 
	///Quiesce (9) indicates that the element is enabled but in a restricted mode.
	///Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	EnabledState EnabledLogicalElement_EnabledState

	// A string that describes the enabled or disabled state of the element when the EnabledState property is set to 1 ("Other"). This property must be set to null when EnabledState is any value other than 1.
	OtherEnabledState string

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration. 
	///"Unknown" (0) indicates the last requested state for the element is unknown.
	///Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState. 
	///It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests. 
	///
	///This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code. 
	///
	///If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	RequestedState EnabledLogicalElement_RequestedState

	// The date or time when the EnabledState of the element last changed. If the state of the element has not changed and this property is populated, then it must be set to a 0 interval value. If a state change was requested, but rejected or not yet processed, the property must not be updated.
	TimeOfLastStateChange string

	// TransitioningToState indicates the target state to which the instance is transitioning. 
	///A value of 5 "No Change" shall indicate that no transition is in progress.A value of 12 "Not Applicable" shall indicate the implementation does not support representing ongoing transitions. 
	///A value other than 5 or 12 shall identify the state to which the element is in the process of transitioning.
	TransitioningToState EnabledLogicalElement_TransitioningToState
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
func (instance *CIM_EnabledLogicalElement) SetPropertyAvailableRequestedStates(value []EnabledLogicalElement_AvailableRequestedStates) (err error) { 
    return instance.SetProperty("AvailableRequestedStates", (value))
}

// GetAvailableRequestedStates gets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyAvailableRequestedStates()(value []EnabledLogicalElement_AvailableRequestedStates, err error) { 
    retValue, err := instance.GetProperty("AvailableRequestedStates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, EnabledLogicalElement_AvailableRequestedStates(valuetmp))
    }

    return
}

// SetEnabledDefault sets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledDefault(value EnabledLogicalElement_EnabledDefault) (err error) { 
    return instance.SetProperty("EnabledDefault", (value))
}

// GetEnabledDefault gets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledDefault()(value EnabledLogicalElement_EnabledDefault, err error) { 
    retValue, err := instance.GetProperty("EnabledDefault")
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

    value = EnabledLogicalElement_EnabledDefault(valuetmp)

    return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledState(value EnabledLogicalElement_EnabledState) (err error) { 
    return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledState()(value EnabledLogicalElement_EnabledState, err error) { 
    retValue, err := instance.GetProperty("EnabledState")
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

    value = EnabledLogicalElement_EnabledState(valuetmp)

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
func (instance *CIM_EnabledLogicalElement) SetPropertyRequestedState(value EnabledLogicalElement_RequestedState) (err error) { 
    return instance.SetProperty("RequestedState", (value))
}

// GetRequestedState gets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyRequestedState()(value EnabledLogicalElement_RequestedState, err error) { 
    retValue, err := instance.GetProperty("RequestedState")
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

    value = EnabledLogicalElement_RequestedState(valuetmp)

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
func (instance *CIM_EnabledLogicalElement) SetPropertyTransitioningToState(value EnabledLogicalElement_TransitioningToState) (err error) { 
    return instance.SetProperty("TransitioningToState", (value))
}

// GetTransitioningToState gets the value of TransitioningToState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyTransitioningToState()(value EnabledLogicalElement_TransitioningToState, err error) { 
    retValue, err := instance.GetProperty("TransitioningToState")
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

    value = EnabledLogicalElement_TransitioningToState(valuetmp)

    return
}

// Requests that the state of the element be changed to the value specified in the RequestedState parameter. When the requested state change takes place, the EnabledState and RequestedState of the element will be the same. Invoking the RequestStateChange method multiple times could result in earlier requests being overwritten or lost. 
///A return code of 0 shall indicate the state change was successfully initiated. 
///A return code of 3 shall indicate that the state transition cannot complete within the interval specified by the TimeoutPeriod parameter. 
///A return code of 4096 (0x1000) shall indicate the state change was successfully initiated, a ConcreteJob has been created, and its reference returned in the output parameter Job. Any other return code indicates an error condition.

// <param name="RequestedState" type="EnabledLogicalElement_RequestedState ">The state requested for the element. This information will be placed into the RequestedState property of the instance if the return code of the RequestStateChange method is 0 ('Completed with No Error'), or 4096 (0x1000) ('Job Started'). Refer to the description of the EnabledState and RequestedState properties for the detailed explanations of the RequestedState values.</param>
// <param name="TimeoutPeriod" type="string ">A timeout period that specifies the maximum amount of time that the client expects the transition to the new state to take. The interval format must be used to specify the TimeoutPeriod. A value of 0 or a null parameter indicates that the client has no time requirements for the transition. 
///If this property does not contain 0 or null and the implementation does not support this parameter, a return code of 'Use Of Timeout Parameter Not Supported' shall be returned.</param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_EnabledLogicalElement) RequestStateChange( /* IN */ RequestedState EnabledLogicalElement_RequestedState,
 /* OUT */ Job CIM_ConcreteJob,
 /* OPTIONAL IN */ TimeoutPeriod string,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RequestStateChange", Action, PercentComplete, Timeout , RequestedState, TimeoutPeriod)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


