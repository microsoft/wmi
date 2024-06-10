// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_EnabledLogicalElement struct
type CIM_EnabledLogicalElement struct {
	*CIM_LogicalElement

	// 364
	AvailableRequestedStates []EnabledLogicalElement_AvailableRequestedStates

	// 361
	EnabledDefault EnabledLogicalElement_EnabledDefault

	// 346
	EnabledState EnabledLogicalElement_EnabledState

	// 353
	OtherEnabledState string

	// 354
	RequestedState EnabledLogicalElement_RequestedState

	// 363
	TimeOfLastStateChange string

	// 366
	TransitioningToState EnabledLogicalElement_TransitioningToState
}

func NewCIM_EnabledLogicalElementEx1(instance *cim.WmiInstance) (newInstance *CIM_EnabledLogicalElement, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_EnabledLogicalElement{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_EnabledLogicalElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_EnabledLogicalElement, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_EnabledLogicalElement{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetAvailableRequestedStates sets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyAvailableRequestedStates(value []EnabledLogicalElement_AvailableRequestedStates) (err error) {
	return instance.SetProperty("AvailableRequestedStates", (value))
}

// GetAvailableRequestedStates gets the value of AvailableRequestedStates for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyAvailableRequestedStates() (value []EnabledLogicalElement_AvailableRequestedStates, err error) {
	retValue, err := instance.GetProperty("AvailableRequestedStates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledDefault() (value EnabledLogicalElement_EnabledDefault, err error) {
	retValue, err := instance.GetProperty("EnabledDefault")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledState() (value EnabledLogicalElement_EnabledState, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyOtherEnabledState() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEnabledState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyRequestedState() (value EnabledLogicalElement_RequestedState, err error) {
	retValue, err := instance.GetProperty("RequestedState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyTimeOfLastStateChange() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastStateChange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
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
func (instance *CIM_EnabledLogicalElement) GetPropertyTransitioningToState() (value EnabledLogicalElement_TransitioningToState, err error) {
	retValue, err := instance.GetProperty("TransitioningToState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = EnabledLogicalElement_TransitioningToState(valuetmp)

	return
}

// 367

// <param name="RequestedState" type="EnabledLogicalElement_RequestedState ">371</param>
// <param name="TimeoutPeriod" type="string ">373</param>

// <param name="Job" type="CIM_ConcreteJob ">372</param>
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
