// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package job

import (
	"fmt"
	"strings"
	"time"

	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualSystemJob struct {
	*v2.Msvm_ConcreteJob
}

// NewVirtualSystemJob
func NewVirtualSystemJob(instance *wmi.WmiInstance) (*VirtualSystemJob, error) {
	j, err := v2.NewMsvm_ConcreteJobEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSystemJob{j}, nil
}

func (vmjob *VirtualSystemJob) String() string {
	jtype, err := vmjob.JobType()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Type[%s]", jtype)
}

// GetJobType gets the value of JobType for the instance
func (vmjob *VirtualSystemJob) JobType() (value v2.ConcreteJob_JobType, err error) {
	retValue, err := vmjob.GetProperty("JobType")
	if err != nil {
		return
	}
	valueint, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	value = v2.ConcreteJob_JobType(valueint)
	return
}

// WaitForPercentComplete waits for the percentComplete or timeout
func (vmjob *VirtualSystemJob) WaitForPercentComplete(percentComplete uint16, timeoutSeconds int16) error {
	start := time.Now()
	// Run the loop, only if the job is actually running
	for !vmjob.IsComplete() {
		pComplete, err := vmjob.PercentComplete()
		if err != nil {
			return err
		}
		//fmt.Printf("WaitForPercentComplete [%d]\n ", pComplete)
		// Break if have achieved the target
		if pComplete >= percentComplete {
			break
		}

		time.Sleep(100 * time.Millisecond)
		if timeoutSeconds < 0 {
			continue
		}
		// If we have waited enough time, break
		if time.Since(start) > (time.Duration(timeoutSeconds) * time.Second) {
			state, err2 := vmjob.GetPropertyJobState()
			if err2 != nil {
				state = 0
			}
			exception := vmjob.GetException()
			return errors.Wrapf(errors.Timedout, "WaitForPercentComplete timeout. Current state: [%v], Exception: [%v]", state, exception)
		}
	}

	return vmjob.GetException()
}

// WaitForAction waits for the task based on the action type, percent complete and timeoutSeconds
func (vmjob *VirtualSystemJob) WaitForAction(action wmi.UserAction, percentComplete uint16, timeoutSeconds int16) error {
	switch action {
	case wmi.Wait:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
	case wmi.Cancel:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
		// Fixme
		// vm.Cancel()
	case wmi.None:
		fallthrough
	case wmi.Default:
		fallthrough
	case wmi.Async:
		break
	}
	return nil
}

// PercentComplete
func (vmjob *VirtualSystemJob) PercentComplete() (uint16, error) {
	err := vmjob.Refresh()
	if err != nil {
		return 0, err
	}
	retValue, err := vmjob.GetProperty("PercentComplete")
	if err != nil {
		return 0, err
	}
	return uint16(retValue.(int32)), nil
}

func (vmjob *VirtualSystemJob) IsComplete() bool {
	err := vmjob.Refresh()
	if err != nil {

	}
	state, err := vmjob.GetPropertyJobState()
	if err != nil {
		return false
	}
	switch state {
	case v2.ConcreteJob_JobState_New:
		fallthrough
	case v2.ConcreteJob_JobState_Starting:
		fallthrough
	case v2.ConcreteJob_JobState_Running:
		fallthrough
	case v2.ConcreteJob_JobState_Suspended:
		fallthrough
	case v2.ConcreteJob_JobState_Shutting_Down:
		return false
	case v2.ConcreteJob_JobState_Completed:
		fallthrough
	case v2.ConcreteJob_JobState_Terminated:
		fallthrough
	case v2.ConcreteJob_JobState_Killed:
		fallthrough
	case v2.ConcreteJob_JobState_Exception:
		return true
	}
	return false
}

func (vmjob *VirtualSystemJob) GetException() error {
	vmjob.Refresh()
	state, err := vmjob.GetPropertyJobState()
	if err != nil {
		return err
	}
	switch state {
	case v2.ConcreteJob_JobState_Terminated:
		fallthrough
	case v2.ConcreteJob_JobState_Killed:
		fallthrough
	case v2.ConcreteJob_JobState_Exception:
		errorCode, _ := vmjob.GetPropertyErrorCode()
		errorDescription, _ := vmjob.GetPropertyErrorDescription()
		errorSummaryDescription, _ := vmjob.GetPropertyErrorSummaryDescription()

		if errorCode == 0 {
			if strings.Contains(errorSummaryDescription, errors.OutOfMemoryErrorSummary) {
				errorCode = errors.ERROR_OUTOFMEMORY
			}
		}

		return errors.Wrapf(errors.NewWMIError(errorCode),
			"ErrorCode[%d] ErrorDescription[%s] ErrorSummaryDescription [%s]",
			errorCode, errorDescription, errorSummaryDescription)
	}
	return nil
}

func WaitForJobCompletionEx(result int32, currentJob *VirtualSystemJob, timeoutSeconds int16) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		return currentJob.WaitForAction(wmi.Wait, 100, timeoutSeconds)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Result[%d] ", result)
	}

}

func WaitForJobCompletionEx2(instance *wmi.WmiInstance, result int32, jobName string) error {
	return nil
}

func WaitForJobCompletion(instance *wmi.WmiInstance, result int32, jobType v2.ConcreteJob_JobType, timeoutSeconds int16) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := getJob(instance, jobType)
		if err != nil {
			// Job is scheduled, but we were not able to find the job
			return err
		}
		defer vmjob.Close()
		return vmjob.WaitForAction(wmi.Wait, 100, timeoutSeconds)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Resource Pool Result[%d] JobType[%d]", result, jobType)
	}
}

func getJob(instance *wmi.WmiInstance, jobType v2.ConcreteJob_JobType) (*VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := instance.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return NewVirtualSystemJob(jobInstance)
}

func getCIMJob(instance *wmi.WmiInstance, jobName string) (*VirtualSystemJob, error) {
	query := query.NewWmiQuery("CIM_ConcreteJob", "Name", jobName)
	jobs, err := instance.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with name [%s]", jobName)
	}
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return NewVirtualSystemJob(jobInstance)
}
