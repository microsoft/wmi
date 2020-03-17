// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_ConcreteJob struct
type CIM_ConcreteJob struct {
	CIM_Job

	// 314
	JobState ConcreteJob_JobState

	// 324
	TimeBeforeRemoval string

	// 323
	TimeOfLastStateChange string
}

// SetJobState sets the value of JobState for the instance
func (instance *CIM_ConcreteJob) SetPropertyJobState(value ConcreteJob_JobState) (err error) {
	return instance.SetProperty("JobState", value)
}

// GetJobState gets the value of JobState for the instance
func (instance *CIM_ConcreteJob) GetPropertyJobState() (value ConcreteJob_JobState, err error) {
	retValue, err := instance.GetProperty("JobState")
	if err != nil {
		return
	}
	value, ok := retValue.(ConcreteJob_JobState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeBeforeRemoval sets the value of TimeBeforeRemoval for the instance
func (instance *CIM_ConcreteJob) SetPropertyTimeBeforeRemoval(value string) (err error) {
	return instance.SetProperty("TimeBeforeRemoval", value)
}

// GetTimeBeforeRemoval gets the value of TimeBeforeRemoval for the instance
func (instance *CIM_ConcreteJob) GetPropertyTimeBeforeRemoval() (value string, err error) {
	retValue, err := instance.GetProperty("TimeBeforeRemoval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastStateChange sets the value of TimeOfLastStateChange for the instance
func (instance *CIM_ConcreteJob) SetPropertyTimeOfLastStateChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastStateChange", value)
}

// GetTimeOfLastStateChange gets the value of TimeOfLastStateChange for the instance
func (instance *CIM_ConcreteJob) GetPropertyTimeOfLastStateChange() (value string, err error) {
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

// 325

// <param name="RequestedState" type="ConcreteJob_RequestedState ">336</param>
// <param name="TimeoutPeriod" type="string ">341</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ConcreteJob) RequestStateChange( /* IN */ RequestedState ConcreteJob_RequestedState,
	/* IN */ TimeoutPeriod string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RequestStateChange", RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 342

// <param name="Error" type="CIM_Error ">344</param>
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
