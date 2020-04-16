// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
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
func (vm *VirtualMachine) State() (state VirtualMachineState, err error) {
	err = vm.Refresh()
	if err != nil {
		return
	}

	retValue, err := vm.GetProperty("EnabledState")
	if err != nil {
		return
	}
	intstate, ok := retValue.(int32)
	if !ok {
		return Unknown, errors.Wrapf(errors.Failed, "Failed to get the state of the VM [%+v]", retValue)
	}
	return VirtualMachineState(intstate), nil
}

// Stop Virtual Machine
func (vm *VirtualMachine) Stop(force bool) error {
	return vm.ChangeState(Off, v2.ConcreteJob_JobType_Power_Off_Virtual_Machine)
}

// Start Virtual Machine
func (vm *VirtualMachine) Start() error {
	err := vm.ChangeState(Running, v2.ConcreteJob_JobType_Start_Virtual_Machine)
	if err != nil {
		return err
	}
	return vm.WaitForState(Running, 30)
}

// ChangeState changes the state of the Virtual Machine
func (vm *VirtualMachine) ChangeState(state VirtualMachineState, jobType v2.ConcreteJob_JobType) (err error) {
	cstate, err := vm.State()
	if err != nil {
		return err
	}
	// If the state is already satisfied, just return
	if cstate == state {
		return nil
	}
	result, err := vm.InvokeMethodWithReturn("RequestStateChange", int32(state))
	if err != nil {
		return err
	}
	return vm.WaitForJobCompletion(result, jobType)
}

// WaitForState
func (vm *VirtualMachine) WaitForState(state VirtualMachineState, timeoutSeconds int32) (err error) {
	start := time.Now()
	// Run the loop, only if the job is actually running
	for {
		curState, err1 := vm.State()
		if err1 != nil {
			return err1
		}

		if curState == state {
			// Break for any valid state
			// TODO: WaitForSomeState
			return nil
		}
		time.Sleep(100 * time.Millisecond)

		// If we have waited enough time, break
		if time.Since(start) > (time.Duration(timeoutSeconds) * time.Second) {
			err = errors.Wrapf(errors.Timedout, "WaitForState timeout")
			break
		}
	}

	return
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
		defer vmjob.Close()

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed,
			"Unable to Change the state of the Virtual Machine Result[%d] [%d]", result, jobType)
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
	defer jobs.Close()
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return job.NewVirtualSystemJob(jobInstance)
}

func (vm *VirtualMachine) GetVirtualMachineSetting() (*VirtualMachineSetting, error) {
	inst, err := vm.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualMachineSetting(inst)
}

func (vm *VirtualMachine) GetVirtualNetworkAdapterByName(name string) (vna *na.VirtualNetworkAdapter, err error) {
	settings, err := vm.GetVirtualMachineSetting()
	if err != nil {
		return
	}
	defer settings.Close()
	vna, err = settings.GetVirtualNetworkAdapterByName(name)
	return
}
