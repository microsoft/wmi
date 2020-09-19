// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageJob struct
type MSFT_StorageJob struct {
	*MSFT_StorageObject

	// Indicates the number of bytes processed by this job so far.
	BytesProcessed uint64

	// Indicates the total number of bytes being processed by this job.
	BytesTotal uint64

	// If TRUE, the storage job will be automatically deleted after a short time interval.
	DeleteOnCompletion bool

	// The Description property provides a textual description of the storage job operation.
	Description string

	// The time interval that the job has been executing or the total execution time if the storage job is complete.
	ElapsedTime string

	// If the operation that this storage job was tracking has failed, the provider will set this with an error code defined by the method that invoked the operation. If this job tracked a background task, the error code can be set to any valid Storage Management error code as defined in the value map below. If there was no error, this property must be set to 0 - 'Success'. This property should be NULL until the operation has completed.
	ErrorCode uint16

	// A free-form string that contains the vendor error description.
	ErrorDescription string

	// If TRUE, this storage job represents an automated background task initiated by the storage subsystem. For all user / management initiated operations, this value should be set to FALSE.
	IsBackgroundTask bool

	// The current execution state of the storage job.
	JobState StorageJob_JobState

	// A free-form string that represents the status of the job. The primary status is reflected in the inherited OperationalStatus property. JobStatus provides additional, implementation-specific details.
	JobStatus string

	// This property indicates whether the times represented in the StartTime, TimeOfLastStateChange, and TimeSubmitted properties represent local times or UTC times. Time values are synchronized worldwide by using the enumeration value 2 - 'UTC Time'.
	LocalOrUtcTime StorageJob_LocalOrUtcTime

	// A system defined name for this storage job.
	Name string

	// Indicates the current statuses of the element.
	OperationalStatus []StorageJob_OperationalStatus

	// Denotes a vendor-specific recovery action to be taken for an unsuccessfully run job. This value should only be set if RecoveryAction is set to 1 - 'Other'.
	OtherRecoveryAction string

	// The percentage of the job that has completed at the time that this value is requested.
	PercentComplete uint16

	// Describes the recovery action to be taken for an unsuccessfully run job. The possible values are:
	///0 - 'Unknown' meaning it is unknown as to what recovery action to take
	///1 - 'Other' indicating that the recovery action will be specified in the OtherRecoveryAction property
	///2 - 'Do Not Continue' meaning stop the execution of the job and appropriately update its status
	///3 - 'Continue With Next Job' meaning continue with the next job in the queue
	///4 - 'Re-run Job' indicating that the job should be re-run
	///
	RecoveryAction StorageJob_RecoveryAction

	// The time that the job was actually started.
	StartTime string

	// Strings describing the various OperationalStatus array values. For example, if "Stopping" is the value assigned to OperationalStatus, this property may contain an explanation as to why an object is being stopped. Note that entries in this array are correlated with those at the same array index in OperationalStatus.
	StatusDescriptions []string

	// The amount of time that the Job is retained after it has finished executing, regardless of whether it failed during execution. The job must remain in existence for some period of time regardless of the value of the DeleteOnCompletion property.
	///
	TimeBeforeRemoval string

	// The date or time when the state of the job last changed. If the state of the job has not changed and this property is populated, it must be set to a 0 interval value. If a state change was requested, but was rejected or not yet processed, the property must not be updated.
	TimeOfLastStateChange string

	// The time that the job was submitted to execute. A value of all zeroes indicates that the owning element is not capable of reporting a date and time.
	TimeSubmitted string
}

func NewMSFT_StorageJobEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageJob, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageJob{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_StorageJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageJob, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageJob{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetBytesProcessed sets the value of BytesProcessed for the instance
func (instance *MSFT_StorageJob) SetPropertyBytesProcessed(value uint64) (err error) {
	return instance.SetProperty("BytesProcessed", (value))
}

// GetBytesProcessed gets the value of BytesProcessed for the instance
func (instance *MSFT_StorageJob) GetPropertyBytesProcessed() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesProcessed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetBytesTotal sets the value of BytesTotal for the instance
func (instance *MSFT_StorageJob) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", (value))
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *MSFT_StorageJob) GetPropertyBytesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDeleteOnCompletion sets the value of DeleteOnCompletion for the instance
func (instance *MSFT_StorageJob) SetPropertyDeleteOnCompletion(value bool) (err error) {
	return instance.SetProperty("DeleteOnCompletion", (value))
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *MSFT_StorageJob) GetPropertyDeleteOnCompletion() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteOnCompletion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageJob) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageJob) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *MSFT_StorageJob) SetPropertyElapsedTime(value string) (err error) {
	return instance.SetProperty("ElapsedTime", (value))
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *MSFT_StorageJob) GetPropertyElapsedTime() (value string, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
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

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *MSFT_StorageJob) SetPropertyErrorCode(value uint16) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *MSFT_StorageJob) GetPropertyErrorCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
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

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *MSFT_StorageJob) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", (value))
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *MSFT_StorageJob) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
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

// SetIsBackgroundTask sets the value of IsBackgroundTask for the instance
func (instance *MSFT_StorageJob) SetPropertyIsBackgroundTask(value bool) (err error) {
	return instance.SetProperty("IsBackgroundTask", (value))
}

// GetIsBackgroundTask gets the value of IsBackgroundTask for the instance
func (instance *MSFT_StorageJob) GetPropertyIsBackgroundTask() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBackgroundTask")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetJobState sets the value of JobState for the instance
func (instance *MSFT_StorageJob) SetPropertyJobState(value StorageJob_JobState) (err error) {
	return instance.SetProperty("JobState", (value))
}

// GetJobState gets the value of JobState for the instance
func (instance *MSFT_StorageJob) GetPropertyJobState() (value StorageJob_JobState, err error) {
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

	value = StorageJob_JobState(valuetmp)

	return
}

// SetJobStatus sets the value of JobStatus for the instance
func (instance *MSFT_StorageJob) SetPropertyJobStatus(value string) (err error) {
	return instance.SetProperty("JobStatus", (value))
}

// GetJobStatus gets the value of JobStatus for the instance
func (instance *MSFT_StorageJob) GetPropertyJobStatus() (value string, err error) {
	retValue, err := instance.GetProperty("JobStatus")
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

// SetLocalOrUtcTime sets the value of LocalOrUtcTime for the instance
func (instance *MSFT_StorageJob) SetPropertyLocalOrUtcTime(value StorageJob_LocalOrUtcTime) (err error) {
	return instance.SetProperty("LocalOrUtcTime", (value))
}

// GetLocalOrUtcTime gets the value of LocalOrUtcTime for the instance
func (instance *MSFT_StorageJob) GetPropertyLocalOrUtcTime() (value StorageJob_LocalOrUtcTime, err error) {
	retValue, err := instance.GetProperty("LocalOrUtcTime")
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

	value = StorageJob_LocalOrUtcTime(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageJob) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageJob) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageJob) SetPropertyOperationalStatus(value []StorageJob_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageJob) GetPropertyOperationalStatus() (value []StorageJob_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
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
		value = append(value, StorageJob_OperationalStatus(valuetmp))
	}

	return
}

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *MSFT_StorageJob) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", (value))
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *MSFT_StorageJob) GetPropertyOtherRecoveryAction() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecoveryAction")
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

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSFT_StorageJob) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSFT_StorageJob) GetPropertyPercentComplete() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
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

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *MSFT_StorageJob) SetPropertyRecoveryAction(value StorageJob_RecoveryAction) (err error) {
	return instance.SetProperty("RecoveryAction", (value))
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *MSFT_StorageJob) GetPropertyRecoveryAction() (value StorageJob_RecoveryAction, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
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

	value = StorageJob_RecoveryAction(valuetmp)

	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *MSFT_StorageJob) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MSFT_StorageJob) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
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

// SetStatusDescriptions sets the value of StatusDescriptions for the instance
func (instance *MSFT_StorageJob) SetPropertyStatusDescriptions(value []string) (err error) {
	return instance.SetProperty("StatusDescriptions", (value))
}

// GetStatusDescriptions gets the value of StatusDescriptions for the instance
func (instance *MSFT_StorageJob) GetPropertyStatusDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("StatusDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetTimeBeforeRemoval sets the value of TimeBeforeRemoval for the instance
func (instance *MSFT_StorageJob) SetPropertyTimeBeforeRemoval(value string) (err error) {
	return instance.SetProperty("TimeBeforeRemoval", (value))
}

// GetTimeBeforeRemoval gets the value of TimeBeforeRemoval for the instance
func (instance *MSFT_StorageJob) GetPropertyTimeBeforeRemoval() (value string, err error) {
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
func (instance *MSFT_StorageJob) SetPropertyTimeOfLastStateChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastStateChange", (value))
}

// GetTimeOfLastStateChange gets the value of TimeOfLastStateChange for the instance
func (instance *MSFT_StorageJob) GetPropertyTimeOfLastStateChange() (value string, err error) {
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

// SetTimeSubmitted sets the value of TimeSubmitted for the instance
func (instance *MSFT_StorageJob) SetPropertyTimeSubmitted(value string) (err error) {
	return instance.SetProperty("TimeSubmitted", (value))
}

// GetTimeSubmitted gets the value of TimeSubmitted for the instance
func (instance *MSFT_StorageJob) GetPropertyTimeSubmitted() (value string, err error) {
	retValue, err := instance.GetProperty("TimeSubmitted")
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

// <param name="RequestedState" type="StorageJob_RequestedState ">RequestStateChange changes the state of a job. The possible values are as follows:
///2 - 'Start' changes the state to 'Running'.
///3 - 'Suspend' stops the job temporarily. The intention is to subsequently restart the job with a second call to RequestStateChange requesting 1 - 'Start'. It might be possible to enter the 'Service' state while suspended. (This is job-specific.)
///4 - 'Terminate' stops the job cleanly, saving data, preserving the state, and shutting down all underlying processes in an orderly manner.
///5 - 'Kill' terminates the job immediately with no requirement to save data or preserve the state.
///6 - 'Service' puts the job into a vendor-specific service state. It might be possible to restart the job.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageJob) RequestStateChange( /* IN */ RequestedState StorageJob_RequestedState,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RequestStateChange", RequestedState)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageJob) GetExtendedStatus( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetExtendedStatus")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Channels" type="uint16 []"></param>
// <param name="Messages" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageJob) GetMessages( /* OUT */ Channels []uint16,
	/* OUT */ Messages []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMessages")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OutParameters" type="MSFT_StorageJobOutParams "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageJob) GetOutParameters( /* OUT */ OutParameters MSFT_StorageJobOutParams) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetOutParameters")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
