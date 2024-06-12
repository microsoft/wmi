// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.HardwareManagement
//////////////////////////////////////////////
package hardwaremanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Job struct
type CIM_Job struct {
	*CIM_LogicalElement

	//
	DeleteOnCompletion bool

	//
	ElapsedTime string

	//
	ErrorCode uint16

	//
	ErrorDescription string

	//
	JobRunTimes uint32

	//
	JobStatus string

	//
	LocalOrUtcTime uint16

	//
	Notify string

	//
	OtherRecoveryAction string

	//
	Owner string

	//
	PercentComplete uint16

	//
	Priority uint32

	//
	RecoveryAction uint16

	//
	RunDay int8

	//
	RunDayOfWeek int8

	//
	RunMonth uint8

	//
	RunStartInterval string

	//
	ScheduledStartTime string

	//
	StartTime string

	//
	TimeSubmitted string

	//
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
	return instance.SetProperty("DeleteOnCompletion", (value))
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *CIM_Job) GetPropertyDeleteOnCompletion() (value bool, err error) {
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

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *CIM_Job) SetPropertyElapsedTime(value string) (err error) {
	return instance.SetProperty("ElapsedTime", (value))
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *CIM_Job) GetPropertyElapsedTime() (value string, err error) {
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
func (instance *CIM_Job) SetPropertyErrorCode(value uint16) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_Job) GetPropertyErrorCode() (value uint16, err error) {
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
func (instance *CIM_Job) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", (value))
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *CIM_Job) GetPropertyErrorDescription() (value string, err error) {
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

// SetJobRunTimes sets the value of JobRunTimes for the instance
func (instance *CIM_Job) SetPropertyJobRunTimes(value uint32) (err error) {
	return instance.SetProperty("JobRunTimes", (value))
}

// GetJobRunTimes gets the value of JobRunTimes for the instance
func (instance *CIM_Job) GetPropertyJobRunTimes() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobRunTimes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetJobStatus sets the value of JobStatus for the instance
func (instance *CIM_Job) SetPropertyJobStatus(value string) (err error) {
	return instance.SetProperty("JobStatus", (value))
}

// GetJobStatus gets the value of JobStatus for the instance
func (instance *CIM_Job) GetPropertyJobStatus() (value string, err error) {
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
func (instance *CIM_Job) SetPropertyLocalOrUtcTime(value uint16) (err error) {
	return instance.SetProperty("LocalOrUtcTime", (value))
}

// GetLocalOrUtcTime gets the value of LocalOrUtcTime for the instance
func (instance *CIM_Job) GetPropertyLocalOrUtcTime() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalOrUtcTime")
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

// SetNotify sets the value of Notify for the instance
func (instance *CIM_Job) SetPropertyNotify(value string) (err error) {
	return instance.SetProperty("Notify", (value))
}

// GetNotify gets the value of Notify for the instance
func (instance *CIM_Job) GetPropertyNotify() (value string, err error) {
	retValue, err := instance.GetProperty("Notify")
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

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", (value))
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) GetPropertyOtherRecoveryAction() (value string, err error) {
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

// SetOwner sets the value of Owner for the instance
func (instance *CIM_Job) SetPropertyOwner(value string) (err error) {
	return instance.SetProperty("Owner", (value))
}

// GetOwner gets the value of Owner for the instance
func (instance *CIM_Job) GetPropertyOwner() (value string, err error) {
	retValue, err := instance.GetProperty("Owner")
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
func (instance *CIM_Job) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *CIM_Job) GetPropertyPercentComplete() (value uint16, err error) {
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

// SetPriority sets the value of Priority for the instance
func (instance *CIM_Job) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_Job) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *CIM_Job) SetPropertyRecoveryAction(value uint16) (err error) {
	return instance.SetProperty("RecoveryAction", (value))
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *CIM_Job) GetPropertyRecoveryAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
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

// SetRunDay sets the value of RunDay for the instance
func (instance *CIM_Job) SetPropertyRunDay(value int8) (err error) {
	return instance.SetProperty("RunDay", (value))
}

// GetRunDay gets the value of RunDay for the instance
func (instance *CIM_Job) GetPropertyRunDay() (value int8, err error) {
	retValue, err := instance.GetProperty("RunDay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetRunDayOfWeek sets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) SetPropertyRunDayOfWeek(value int8) (err error) {
	return instance.SetProperty("RunDayOfWeek", (value))
}

// GetRunDayOfWeek gets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) GetPropertyRunDayOfWeek() (value int8, err error) {
	retValue, err := instance.GetProperty("RunDayOfWeek")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetRunMonth sets the value of RunMonth for the instance
func (instance *CIM_Job) SetPropertyRunMonth(value uint8) (err error) {
	return instance.SetProperty("RunMonth", (value))
}

// GetRunMonth gets the value of RunMonth for the instance
func (instance *CIM_Job) GetPropertyRunMonth() (value uint8, err error) {
	retValue, err := instance.GetProperty("RunMonth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetRunStartInterval sets the value of RunStartInterval for the instance
func (instance *CIM_Job) SetPropertyRunStartInterval(value string) (err error) {
	return instance.SetProperty("RunStartInterval", (value))
}

// GetRunStartInterval gets the value of RunStartInterval for the instance
func (instance *CIM_Job) GetPropertyRunStartInterval() (value string, err error) {
	retValue, err := instance.GetProperty("RunStartInterval")
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

// SetScheduledStartTime sets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) SetPropertyScheduledStartTime(value string) (err error) {
	return instance.SetProperty("ScheduledStartTime", (value))
}

// GetScheduledStartTime gets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) GetPropertyScheduledStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("ScheduledStartTime")
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

// SetStartTime sets the value of StartTime for the instance
func (instance *CIM_Job) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *CIM_Job) GetPropertyStartTime() (value string, err error) {
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

// SetTimeSubmitted sets the value of TimeSubmitted for the instance
func (instance *CIM_Job) SetPropertyTimeSubmitted(value string) (err error) {
	return instance.SetProperty("TimeSubmitted", (value))
}

// GetTimeSubmitted gets the value of TimeSubmitted for the instance
func (instance *CIM_Job) GetPropertyTimeSubmitted() (value string, err error) {
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

// SetUntilTime sets the value of UntilTime for the instance
func (instance *CIM_Job) SetPropertyUntilTime(value string) (err error) {
	return instance.SetProperty("UntilTime", (value))
}

// GetUntilTime gets the value of UntilTime for the instance
func (instance *CIM_Job) GetPropertyUntilTime() (value string, err error) {
	retValue, err := instance.GetProperty("UntilTime")
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

// <param name="DeleteOnKill" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Job) KillJob( /* IN */ DeleteOnKill bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("KillJob", DeleteOnKill)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
