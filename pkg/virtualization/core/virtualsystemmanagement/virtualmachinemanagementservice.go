// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*VirtualSystemManagementService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*VirtualSystemManagementService{}

	whost := host.NewWmiLocalHost()
	serviceStoreMap[whost.HostName], _ = getService(whost)
}

type VirtualSystemManagementService struct {
	*v2.Msvm_VirtualSystemManagementService
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetVirtualSystemManagementService(whost *host.WmiHost) (mgmt *VirtualSystemManagementService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	mgmt, err = getService(whost)

	mux.Lock()
	defer mux.Unlock()
	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getService(whost *host.WmiHost) (mgmt *VirtualSystemManagementService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualSystemManagementService", "Caption", "Hyper-V Virtual System Management Service")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := v2.NewMsvm_VirtualSystemManagementServiceEx6(whost.HostName, "root/virtualization/v2", creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &VirtualSystemManagementService{vmmswmi}
	return
}

// GetVirtualMachines would get all virtual machines
func (vmms *VirtualSystemManagementService) GetVirtualMachines() (*virtualsystem.VirtualMachineCollection, error) {
	query := query.NewWmiQuery("Msvm_ComputerSystem", "Caption", "Virtual Machine")
	instances, err := vmms.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	vmc := virtualsystem.VirtualMachineCollection{}
	for _, ins := range instances {
		vm, err := virtualsystem.NewVirtualMachine(ins)
		if err != nil {
			return nil, err
		}

		vmc = append(vmc, vm)
	}
	return &vmc, nil
}

// FindVirtualMachineByName
func (vmms *VirtualSystemManagementService) FindVirtualMachineByName(vmName string) (*virtualsystem.VirtualMachine, error) {
	vms, err := vmms.GetVirtualMachines()
	if err != nil {
		return nil, err
	}
	defer vms.Close()

	for _, vm := range *vms {
		curVmName, err := vm.GetPropertyElementName()
		if err != nil {
			return nil, err
		}
		if curVmName != vmName {
			continue
		}

		vminst, err := vm.Clone()
		if err != nil {
			return nil, err
		}
		return virtualsystem.NewVirtualMachine(vminst)
	}

	return nil, errors.Wrapf(errors.NotFound, "Unable to find a virtual system with name [%s]", vmName)
}

func (vmms *VirtualSystemManagementService) DeleteVirtualMachine(vm *virtualsystem.VirtualMachine) error {
	rawInst := vm.GetRawInstance()
	if rawInst == nil {
		return errors.Wrapf(errors.InvalidInput, "nil iDispatch [%+v], VM[%+v]", vm.GetRawInstance(), vm)
	}
	fmt.Printf("inst[%+v], Type[%+v]\n", vm.GetRawInstance(), reflect.TypeOf(rawInst.ToIDispatch()))
	_, err := vm.EmbeddedInstance()
	if err != nil {
		return err
	}
	result, err := vmms.InvokeMethodWithReturn("DestroySystem", vm.InstancePath(), nil)
	//method, err := vmms.GetWmiMethod("DestroySystem")
	if err != nil {
		return err
	}

	//method.AddInParam("AffectedSystem", (vm.GetIDispatch()))
	//result, err := method.Execute()
	//if err != nil {
	//	return err
	//}
	return vmms.WaitForJobCompletion(result, v2.ConcreteJob_JobType_Destroy_Virtual_Machine)
}

func (vmms *VirtualSystemManagementService) WaitForJobCompletion(result int32, jobType v2.ConcreteJob_JobType) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := vmms.getJob(jobType)
		if err != nil {
			// Job is scheduled, but we were not able to find the job
			return err
		}

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Virtual Machine [%d][%d]", result, jobType)
	}
}

func (vmms *VirtualSystemManagementService) getJob(jobType v2.ConcreteJob_JobType) (*job.VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := vmms.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	// FIXME: Find the correct Job, when multiple jobs are returned
	return job.NewVirtualSystemJob(jobs[0])
}
