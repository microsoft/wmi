// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_EnabledLogicalElement struct
type CIM_EnabledLogicalElement struct {
	CIM_LogicalElement

	//
	AvailableRequestedStates []EnabledLogicalElement_AvailableRequestedStates

	//
	EnabledDefault EnabledLogicalElement_EnabledDefault

	//
	EnabledState EnabledLogicalElement_EnabledState

	//
	OtherEnabledState string

	//
	RequestedState EnabledLogicalElement_RequestedState

	//
	TimeOfLastStateChange string

	//
	TransitioningToState EnabledLogicalElement_TransitioningToState
}

// SetAvailableRequestedStates sets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyAvailableRequestedStates(value []EnabledLogicalElement_AvailableRequestedStates) (err error) {
	return instance.SetProperty("AvailableRequestedStates", value)
}

// GetAvailableRequestedStates gets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyAvailableRequestedStates() (value []EnabledLogicalElement_AvailableRequestedStates, err error) {
	retValue, err := instance.GetProperty("AvailableRequestedStates")
	if err != nil {
		return
	}
	value, ok := retValue.([]EnabledLogicalElement_AvailableRequestedStates)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledDefault sets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledDefault(value EnabledLogicalElement_EnabledDefault) (err error) {
	return instance.SetProperty("EnabledDefault", value)
}

// GetEnabledDefault gets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledDefault() (value EnabledLogicalElement_EnabledDefault, err error) {
	retValue, err := instance.GetProperty("EnabledDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(EnabledLogicalElement_EnabledDefault)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledState(value EnabledLogicalElement_EnabledState) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledState() (value EnabledLogicalElement_EnabledState, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(EnabledLogicalElement_EnabledState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherEnabledState sets the value of OtherEnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyOtherEnabledState(value string) (err error) {
	return instance.SetProperty("OtherEnabledState", value)
}

// GetOtherEnabledState gets the value of OtherEnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyOtherEnabledState() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestedState sets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyRequestedState(value EnabledLogicalElement_RequestedState) (err error) {
	return instance.SetProperty("RequestedState", value)
}

// GetRequestedState gets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyRequestedState() (value EnabledLogicalElement_RequestedState, err error) {
	retValue, err := instance.GetProperty("RequestedState")
	if err != nil {
		return
	}
	value, ok := retValue.(EnabledLogicalElement_RequestedState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastStateChange sets the value of TimeOfLastStateChange for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyTimeOfLastStateChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastStateChange", value)
}

// GetTimeOfLastStateChange gets the value of TimeOfLastStateChange for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyTimeOfLastStateChange() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastStateChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransitioningToState sets the value of TransitioningToState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyTransitioningToState(value EnabledLogicalElement_TransitioningToState) (err error) {
	return instance.SetProperty("TransitioningToState", value)
}

// GetTransitioningToState gets the value of TransitioningToState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyTransitioningToState() (value EnabledLogicalElement_TransitioningToState, err error) {
	retValue, err := instance.GetProperty("TransitioningToState")
	if err != nil {
		return
	}
	value, ok := retValue.(EnabledLogicalElement_TransitioningToState)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="RequestedState" type="EnabledLogicalElement_RequestedState "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_EnabledLogicalElement) RequestStateChange( /* IN */ RequestedState EnabledLogicalElement_RequestedState,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestStateChange", Action, PercentComplete, Timeout, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
