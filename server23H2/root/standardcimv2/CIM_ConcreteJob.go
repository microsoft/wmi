// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ConcreteJob struct
type CIM_ConcreteJob struct {
	*CIM_Job

	//
	JobState uint16

	//
	TimeBeforeRemoval string

	//
	TimeOfLastStateChange string
}

func NewCIM_ConcreteJobEx1(instance *cim.WmiInstance) (newInstance *CIM_ConcreteJob, err error) {
	tmp, err := NewCIM_JobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteJob{
		CIM_Job: tmp,
	}
	return
}

func NewCIM_ConcreteJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ConcreteJob, err error) {
	tmp, err := NewCIM_JobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteJob{
		CIM_Job: tmp,
	}
	return
}

// SetJobState sets the value of JobState for the instance
func (instance *CIM_ConcreteJob) SetPropertyJobState(value uint16) (err error) {
	return instance.SetProperty("JobState", (value))
}

// GetJobState gets the value of JobState for the instance
func (instance *CIM_ConcreteJob) GetPropertyJobState() (value uint16, err error) {
	retValue, err := instance.GetProperty("JobState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetTimeBeforeRemoval sets the value of TimeBeforeRemoval for the instance
func (instance *CIM_ConcreteJob) SetPropertyTimeBeforeRemoval(value string) (err error) {
	return instance.SetProperty("TimeBeforeRemoval", (value))
}

// GetTimeBeforeRemoval gets the value of TimeBeforeRemoval for the instance
func (instance *CIM_ConcreteJob) GetPropertyTimeBeforeRemoval() (value string, err error) {
	retValue, err := instance.GetProperty("TimeBeforeRemoval")
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

// SetTimeOfLastStateChange sets the value of TimeOfLastStateChange for the instance
func (instance *CIM_ConcreteJob) SetPropertyTimeOfLastStateChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastStateChange", (value))
}

// GetTimeOfLastStateChange gets the value of TimeOfLastStateChange for the instance
func (instance *CIM_ConcreteJob) GetPropertyTimeOfLastStateChange() (value string, err error) {
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

//

// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ConcreteJob) RequestStateChange( /* IN */ RequestedState uint16,
	/* IN */ TimeoutPeriod string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RequestStateChange", RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Error" type="CIM_Error "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ConcreteJob) GetError( /* OUT */ Error CIM_Error) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetError")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
