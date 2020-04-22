// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package job

import (
	//"fmt"
	//"reflect"
	"time"

	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualSystemJob struct {
	*v2.Msvm_ConcreteJob
}

// NewVirtualSystemJob
func NewVirtualSystemJob(instance *cim.WmiInstance) (*VirtualSystemJob, error) {
	j, err := v2.NewMsvm_ConcreteJobEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSystemJob{j}, nil
}

// WaitForPercentComplete waits for the percentComplete or timeout
func (vmjob *VirtualSystemJob) WaitForPercentComplete(percentComplete, timeoutSeconds uint16) error {
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
		// If we have waited enough time, break
		if time.Since(start) > (time.Duration(timeoutSeconds) * time.Second) {
			break
		}
	}

	return vmjob.GetException()
}

// WaitForAction waits for the task based on the action type, percent complete and timeoutSeconds
func (vmjob *VirtualSystemJob) WaitForAction(action cim.UserAction, percentComplete, timeoutSeconds uint16) error {
	switch action {
	case cim.Wait:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
	case cim.Cancel:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
		// Fixme
		// vm.Cancel()
	case cim.None:
		fallthrough
	case cim.Default:
		fallthrough
	case cim.Async:
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
		return errors.Wrapf(errors.Failed,
			"ErrorCode[%d] ErrorDescription[%s] ErrorSummaryDescription [%s]",
			errorCode, errorDescription, errorSummaryDescription)
	}
	return nil
}
