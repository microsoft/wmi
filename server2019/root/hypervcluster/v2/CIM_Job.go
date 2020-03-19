// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Job struct
type CIM_Job struct {
	*CIM_LogicalElement

	// Indicates whether or not the job should be automatically deleted upon completion. Note that the 'completion' of a recurring job is defined by its JobRunTimes or UntilTime properties, or when the Job is terminated by manual intervention. If this property is set to false and the job completes, then the extrinsic method DeleteInstance must be used to delete the job instead of updating this property.
	DeleteOnCompletion bool

	// The time interval that the Job has been executing or the total execution time if the Job is complete. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the processing information for recurring Jobs, because only the 'last' run time can be stored in this single-valued property.
	ElapsedTime string

	// A vendor-specific error code. The value must be set to zero if the Job completed without error. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the processing information for recurring Jobs, because only the 'last' run error can be stored in this single-valued property.
	ErrorCode uint16

	// A free-form string that contains the vendor error description. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the processing information for recurring Jobs, because only the 'last' run error can be stored in this single-valued property.
	ErrorDescription string

	// The number of times that the Job should be run. A value of 1 indicates that the Job is not recurring, while any non-zero value indicates a limit to the number of times that the Job will recur. Zero indicates that there is no limit to the number of times that the Job can be processed, but that it is terminated either after the UntilTime or by manual intervention. By default, a Job is processed once.
	JobRunTimes uint32

	// A free-form string that represents the status of the job. The primary status is reflected in the inherited OperationalStatus property. JobStatus provides additional, implementation-specific details.
	JobStatus string

	// This property indicates whether the times represented in the RunStartInterval and UntilTime properties represent local times or UTC times. Time values are synchronized worldwide by using the enumeration value 2, "UTC Time".
	LocalOrUtcTime Job_LocalOrUtcTime

	// The User who is to be notified upon the Job completion or failure.
	Notify string

	// A string describing the recovery action when the RecoveryAction property of the instance is 1 ("Other").
	OtherRecoveryAction string

	// The User that submitted the Job, or the Service or method name that caused the job to be created.
	Owner string

	// The percentage of the job that has completed at the time that this value is requested. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the processing information for recurring Jobs, because only the 'last' run data can be stored in this single-valued property.
	///Note that the value 101 is undefined and will be not be allowed in the next major revision of the specification.
	PercentComplete uint16

	// Indicates the urgency or importance of execution of the Job. The lower the number, the higher the priority. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the setting information that would influence the results of a job.
	Priority uint32

	// Describes the recovery action to be taken for an unsuccessfully run Job. The possible values are:
	///0 = "Unknown", meaning it is unknown as to what recovery action to take
	///1 = "Other", indicating that the recovery action will be specified in the OtherRecoveryAction property
	///2 = "Do Not Continue", meaning stop the execution of the job and appropriately update its status
	///3 = "Continue With Next Job", meaning continue with the next job in the queue
	///4 = "Re-run Job", indicating that the job should be re-run
	///5 = "Run Recovery Job", meaning run the Job associated using the RecoveryJob relationship. Note that the recovery Job must already be in the queue from which it will run.
	RecoveryAction Job_RecoveryAction

	// The day in the month on which the Job should be processed. There are two different interpretations for this property, depending on the value of DayOfWeek. In one case, RunDay defines the day-in-month on which the Job is processed. This interpretation is used when the DayOfWeek is 0. A positive or negative integer indicates whether the RunDay should be calculated from the beginning or end of the month. For example, 5 indicates the fifth day in the RunMonth and -1 indicates the last day in the RunMonth.
	///
	///When RunDayOfWeek is not 0, RunDay is the day-in-month on which the Job is processed, defined in conjunction with RunDayOfWeek. For example, if RunDay is 15 and RunDayOfWeek is Saturday, then the Job is processed on the first Saturday on or after the 15th day in the RunMonth (for example, the third Saturday in the month). If RunDay is 20 and RunDayOfWeek is -Saturday, then this indicates the first Saturday on or before the 20th day in the RunMonth. If RunDay is -1 and RunDayOfWeek is -Sunday, then this indicates the last Sunday in the RunMonth.
	RunDay int8

	// A positive or negative integer used in conjunction with RunDay to indicate the day of the week on which the Job is processed. RunDayOfWeek is set to 0 to indicate an exact day of the month, such as March 1. A positive integer (representing Sunday, Monday, ..., Saturday) means that the day of week is found on or after the specified RunDay. A negative integer (representing -Sunday, -Monday, ..., -Saturday) means that the day of week is found on or BEFORE the RunDay.
	RunDayOfWeek Job_RunDayOfWeek

	// The month during which the Job should be processed. Specify 0 for January, 1 for February, and so on.
	RunMonth Job_RunMonth

	// The time interval after midnight when the Job should be processed. For example,
	///00000000020000.000000:000
	///indicates that the Job should be run on or after two o'clock, local time or UTC time (distinguished using the LocalOrUtcTime property.
	RunStartInterval string

	// The time that the current Job is scheduled to start. This time can be represented by the actual date and time, or an interval relative to the time that this property is requested. A value of all zeroes indicates that the Job is already executing. The property is deprecated in lieu of the more expressive scheduling properties, RunMonth, RunDay, RunDayOfWeek, and RunStartInterval.
	ScheduledStartTime string

	// The time that the Job was actually started. This time can be represented by an actual date and time, or by an interval relative to the time that this property is requested. Note that this property is also present in the JobProcessingStatistics class. This class is necessary to capture the processing information for recurring Jobs, because only the 'last' run time can be stored in this single-valued property.
	StartTime string

	// The time that the Job was submitted to execute. A value of all zeroes indicates that the owning element is not capable of reporting a date and time. Therefore, the ScheduledStartTime and StartTime are reported as intervals relative to the time their values are requested.
	TimeSubmitted string

	// The time after which the Job is invalid or should be stopped. This time can be represented by an actual date and time, or by an interval relative to the time that this property is requested. A value of all nines indicates that the Job can run indefinitely.
	UntilTime string
}

func NewCIM_JobEx1(instance *cim.WmiInstance) (newInstance *CIM_Job, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Job{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_JobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Job, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Job{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetDeleteOnCompletion sets the value of DeleteOnCompletion for the instance
func (instance *CIM_Job) SetPropertyDeleteOnCompletion(value bool) (err error) {
	return instance.SetProperty("DeleteOnCompletion", value)
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *CIM_Job) GetPropertyDeleteOnCompletion() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteOnCompletion")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *CIM_Job) SetPropertyElapsedTime(value string) (err error) {
	return instance.SetProperty("ElapsedTime", value)
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *CIM_Job) GetPropertyElapsedTime() (value string, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *CIM_Job) SetPropertyErrorCode(value uint16) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_Job) GetPropertyErrorCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *CIM_Job) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *CIM_Job) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobRunTimes sets the value of JobRunTimes for the instance
func (instance *CIM_Job) SetPropertyJobRunTimes(value uint32) (err error) {
	return instance.SetProperty("JobRunTimes", value)
}

// GetJobRunTimes gets the value of JobRunTimes for the instance
func (instance *CIM_Job) GetPropertyJobRunTimes() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobRunTimes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobStatus sets the value of JobStatus for the instance
func (instance *CIM_Job) SetPropertyJobStatus(value string) (err error) {
	return instance.SetProperty("JobStatus", value)
}

// GetJobStatus gets the value of JobStatus for the instance
func (instance *CIM_Job) GetPropertyJobStatus() (value string, err error) {
	retValue, err := instance.GetProperty("JobStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalOrUtcTime sets the value of LocalOrUtcTime for the instance
func (instance *CIM_Job) SetPropertyLocalOrUtcTime(value Job_LocalOrUtcTime) (err error) {
	return instance.SetProperty("LocalOrUtcTime", value)
}

// GetLocalOrUtcTime gets the value of LocalOrUtcTime for the instance
func (instance *CIM_Job) GetPropertyLocalOrUtcTime() (value Job_LocalOrUtcTime, err error) {
	retValue, err := instance.GetProperty("LocalOrUtcTime")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_LocalOrUtcTime)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotify sets the value of Notify for the instance
func (instance *CIM_Job) SetPropertyNotify(value string) (err error) {
	return instance.SetProperty("Notify", value)
}

// GetNotify gets the value of Notify for the instance
func (instance *CIM_Job) GetPropertyNotify() (value string, err error) {
	retValue, err := instance.GetProperty("Notify")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", value)
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) GetPropertyOtherRecoveryAction() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwner sets the value of Owner for the instance
func (instance *CIM_Job) SetPropertyOwner(value string) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *CIM_Job) GetPropertyOwner() (value string, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *CIM_Job) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *CIM_Job) GetPropertyPercentComplete() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *CIM_Job) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_Job) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *CIM_Job) SetPropertyRecoveryAction(value Job_RecoveryAction) (err error) {
	return instance.SetProperty("RecoveryAction", value)
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *CIM_Job) GetPropertyRecoveryAction() (value Job_RecoveryAction, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RecoveryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunDay sets the value of RunDay for the instance
func (instance *CIM_Job) SetPropertyRunDay(value int8) (err error) {
	return instance.SetProperty("RunDay", value)
}

// GetRunDay gets the value of RunDay for the instance
func (instance *CIM_Job) GetPropertyRunDay() (value int8, err error) {
	retValue, err := instance.GetProperty("RunDay")
	if err != nil {
		return
	}
	value, ok := retValue.(int8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunDayOfWeek sets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) SetPropertyRunDayOfWeek(value Job_RunDayOfWeek) (err error) {
	return instance.SetProperty("RunDayOfWeek", value)
}

// GetRunDayOfWeek gets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) GetPropertyRunDayOfWeek() (value Job_RunDayOfWeek, err error) {
	retValue, err := instance.GetProperty("RunDayOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RunDayOfWeek)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunMonth sets the value of RunMonth for the instance
func (instance *CIM_Job) SetPropertyRunMonth(value Job_RunMonth) (err error) {
	return instance.SetProperty("RunMonth", value)
}

// GetRunMonth gets the value of RunMonth for the instance
func (instance *CIM_Job) GetPropertyRunMonth() (value Job_RunMonth, err error) {
	retValue, err := instance.GetProperty("RunMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RunMonth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunStartInterval sets the value of RunStartInterval for the instance
func (instance *CIM_Job) SetPropertyRunStartInterval(value string) (err error) {
	return instance.SetProperty("RunStartInterval", value)
}

// GetRunStartInterval gets the value of RunStartInterval for the instance
func (instance *CIM_Job) GetPropertyRunStartInterval() (value string, err error) {
	retValue, err := instance.GetProperty("RunStartInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScheduledStartTime sets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) SetPropertyScheduledStartTime(value string) (err error) {
	return instance.SetProperty("ScheduledStartTime", value)
}

// GetScheduledStartTime gets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) GetPropertyScheduledStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("ScheduledStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *CIM_Job) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *CIM_Job) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeSubmitted sets the value of TimeSubmitted for the instance
func (instance *CIM_Job) SetPropertyTimeSubmitted(value string) (err error) {
	return instance.SetProperty("TimeSubmitted", value)
}

// GetTimeSubmitted gets the value of TimeSubmitted for the instance
func (instance *CIM_Job) GetPropertyTimeSubmitted() (value string, err error) {
	retValue, err := instance.GetProperty("TimeSubmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUntilTime sets the value of UntilTime for the instance
func (instance *CIM_Job) SetPropertyUntilTime(value string) (err error) {
	return instance.SetProperty("UntilTime", value)
}

// GetUntilTime gets the value of UntilTime for the instance
func (instance *CIM_Job) GetPropertyUntilTime() (value string, err error) {
	retValue, err := instance.GetProperty("UntilTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// KillJob is being deprecated because there is no distinction made between an orderly shutdown and an immediate kill. CIM_ConcreteJob.RequestStateChange() provides 'Terminate' and 'Kill' options to allow this distinction.
///A method to kill this job and any underlying processes, and to remove any 'dangling' associations.

// <param name="DeleteOnKill" type="bool ">Indicates whether or not the Job should be automatically deleted upon termination. This parameter takes precedence over the property, DeleteOnCompletion.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Job) KillJob( /* IN */ DeleteOnKill bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("KillJob", DeleteOnKill)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
