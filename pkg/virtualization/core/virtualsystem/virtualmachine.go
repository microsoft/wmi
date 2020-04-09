// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualMachine struct {
	*v2.Msvm_ComputerSystem
}

type VirtualMachineState int32

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/virtual/msvm-computersystem?redirectedfrom=MSDN
const (
	Unknown            VirtualMachineState = 0
	Other              VirtualMachineState = 1
	Running            VirtualMachineState = 2
	Off                VirtualMachineState = 3
	Stopping           VirtualMachineState = 4
	Saved              VirtualMachineState = 6
	Paused             VirtualMachineState = 9
	Starting           VirtualMachineState = 10
	Reset              VirtualMachineState = 11
	Saving             VirtualMachineState = 32773
	Pausing            VirtualMachineState = 32776
	Resuming           VirtualMachineState = 32777
	FastSaved          VirtualMachineState = 32779
	FastSaving         VirtualMachineState = 32780
	ForceShutdown      VirtualMachineState = 32781
	ForceReboot        VirtualMachineState = 32782
	Hibernated         VirtualMachineState = 32783
	ComponentServicing VirtualMachineState = 32784
	RunningCritical    VirtualMachineState = 32785
	OffCritical        VirtualMachineState = 32786
	StoppingCritial    VirtualMachineState = 32787
	SavedCritical      VirtualMachineState = 32788
	PausedCritical     VirtualMachineState = 32789
	StartingCritical   VirtualMachineState = 32790
	ResetCritical      VirtualMachineState = 32791
	SavingCritical     VirtualMachineState = 32792
	PausingCritical    VirtualMachineState = 32793
	ResumingCritical   VirtualMachineState = 32794
	FastSaveCritical   VirtualMachineState = 32795
	FastSavingCritical VirtualMachineState = 32796
)

// NewVirtualMachine
func NewVirtualMachine(instance *wmi.WmiInstance) (*VirtualMachine, error) {
	wmivm, err := v2.NewMsvm_ComputerSystemEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualMachine{wmivm}, nil
}

// GetVirtualMachine gets an existing virtual machine
func GetVirtualMachine(whost *host.WmiHost, vmName string) (vm *VirtualMachine, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ComputerSystem", "ElementName", vmName)
	wmivm, err := v2.NewMsvm_ComputerSystemEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	vm = &VirtualMachine{wmivm}
	return
}

// GetVirtualMachineState get the virtual machine state
func (vm *VirtualMachine) State() (VirtualMachineState, error) {
	state, err := vm.GetPropertyEnabledState()
	if err != nil {
		return Unknown, err
	}
	return VirtualMachineState(state), nil
}

// Stop Virtual Machine
func (vm *VirtualMachine) Stop(force bool) error {
	jobState := v2.ConcreteJob_JobType_Shut_Down_Virtual_Machine
	state := Stopping
	//job := v2.CIM_ConcreteJob{}
	if force {
		state = Off
		jobState = v2.ConcreteJob_JobType_Power_Off_Virtual_Machine
	}
	return vm.ChangeState(state, jobState)
}

// Start Virtual Machine
func (vm *VirtualMachine) Start() error {
	return vm.ChangeState(Running, v2.ConcreteJob_JobType_Start_Virtual_Machine)
}

// ChangeState changes the state of the Virtual Machine
func (vm *VirtualMachine) ChangeState(state VirtualMachineState, jobType v2.ConcreteJob_JobType) error {
	result, err := vm.InvokeMethodWithReturn("RequestStateChange", int32(state))
	if err != nil {
		return err
	}
	return vm.WaitForJobCompletion(result, jobType)
}

func (vm *VirtualMachine) WaitForJobCompletion(result int32, jobType v2.ConcreteJob_JobType) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := vm.getJob(jobType)
		if err != nil {
			return err
			// Job is scheduled, but we were not able to find the job
		}

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Change the state of the Virtual Machine Result[%d]", result)
	}
}

func (vm *VirtualMachine) getJob(jobType v2.ConcreteJob_JobType) (*job.VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := vm.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	return job.NewVirtualSystemJob(jobs[0])
}
