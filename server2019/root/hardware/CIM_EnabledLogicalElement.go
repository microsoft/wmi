// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
	standardcimv2 "github.com/microsoft/wmi/server2019/root/standardcimv2"
)

// CIM_EnabledLogicalElement struct
type CIM_EnabledLogicalElement struct {
	CIM_LogicalElement

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
}

// SetEnabledDefault sets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledDefault(value uint16) (err error) {
	return instance.SetProperty("EnabledDefault", value)
}

// GetEnabledDefault gets the value of EnabledDefault for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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
func (instance *CIM_EnabledLogicalElement) SetPropertyRequestedState(value uint16) (err error) {
	return instance.SetProperty("RequestedState", value)
}

// GetRequestedState gets the value of RequestedState for the instance
func (instance *CIM_EnabledLogicalElement) GetPropertyRequestedState() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequestedState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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

//

// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_EnabledLogicalElement) RequestStateChange( /* IN */ RequestedState uint16,
	/* OUT */ Job standardcimv2.CIM_ConcreteJob,
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
