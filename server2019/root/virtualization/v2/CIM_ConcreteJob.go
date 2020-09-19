// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ConcreteJob struct
type CIM_ConcreteJob struct {
	*CIM_Job

	// JobState is an integer enumeration that indicates the operational state of a Job. It can also indicate transitions between these states, for example, 'Shutting Down' and 'Starting'. Following is a brief description of the states:
	///New (2) indicates that the job has never been started.
	///Starting (3) indicates that the job is moving from the 'New', 'Suspended', or 'Service' states into the 'Running' state.
	///Running (4) indicates that the Job is running.
	///Suspended (5) indicates that the Job is stopped, but can be restarted in a seamless manner.
	///Shutting Down (6) indicates that the job is moving to a 'Completed', 'Terminated', or 'Killed' state.
	///Completed (7) indicates that the job has completed normally.
	///Terminated (8) indicates that the job has been stopped by a 'Terminate' state change request. The job and all its underlying processes are ended and can be restarted (this is job-specific) only as a new job.
	///Killed (9) indicates that the job has been stopped by a 'Kill' state change request. Underlying processes might have been left running, and cleanup might be required to free up resources.
	///Exception (10) indicates that the Job is in an abnormal state that might be indicative of an error condition. Actual status might be displayed though job-specific objects.
	///Service (11) indicates that the Job is in a vendor-specific state that supports problem discovery, or resolution, or both.
	///Query pending (12) waiting for a client to resolve a query
	JobState ConcreteJob_JobState

	// The amount of time that the Job is retained after it has finished executing, either succeeding or failing in that execution. The job must remain in existence for some period of time regardless of the value of the DeleteOnCompletion property.
	///The default is five minutes.
	TimeBeforeRemoval string

	// The date or time when the state of the Job last changed. If the state of the Job has not changed and this property is populated, then it must be set to a 0 interval value. If a state change was requested, but rejected or not yet processed, the property must not be updated.
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
func (instance *CIM_ConcreteJob) SetPropertyJobState(value ConcreteJob_JobState) (err error) {
	return instance.SetProperty("JobState", (value))
}

// GetJobState gets the value of JobState for the instance
func (instance *CIM_ConcreteJob) GetPropertyJobState() (value ConcreteJob_JobState, err error) {
	retValue, err := instance.GetProperty("JobState")
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

	value = ConcreteJob_JobState(valuetmp)

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

// Requests that the state of the job be changed to the value specified in the RequestedState parameter. Invoking the RequestStateChange method multiple times could result in earlier requests being overwritten or lost.
///If 0 is returned, then the task completed successfully. Any other return code indicates an error condition.

// <param name="RequestedState" type="ConcreteJob_RequestedState ">RequestStateChange changes the state of a job. The possible values are as follows:
///Start (2) changes the state to 'Running'.
///Suspend (3) stops the job temporarily. The intention is to subsequently restart the job with 'Start'. It might be possible to enter the 'Service' state while suspended. (This is job-specific.)
///Terminate (4) stops the job cleanly, saving data, preserving the state, and shutting down all underlying processes in an orderly manner.
///Kill (5) terminates the job immediately with no requirement to save data or preserve the state.
///Service (6) puts the job into a vendor-specific service state. It might be possible to restart the job.</param>
// <param name="TimeoutPeriod" type="string ">A timeout period that specifies the maximum amount of time that the client expects the transition to the new state to take. The interval format must be used to specify the TimeoutPeriod. A value of 0 or a null parameter indicates that the client has no time requirements for the transition.
///If this property does not contain 0 or null and the implementation does not support this parameter, a return code of 'Use Of Timeout Parameter Not Supported' must be returned.</param>

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

// When the job is executing or has terminated without error, then this method returns no CIM_Error instance. However, if the job has failed because of some internal problem or because the job has been terminated by a client, then a CIM_Error instance is returned.

// <param name="Error" type="string ">If the OperationalStatus on the Job is not "OK", then this method will return a CIM Error instance. Otherwise, when the Job is "OK", null is returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ConcreteJob) GetError( /* OUT */ Error string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetError")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
