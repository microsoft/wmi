// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

import (
	"fmt"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	vswitch "github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
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
	if err != nil {
		return
	}

	mux.Lock()
	defer mux.Unlock()
	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getService(whost *host.WmiHost) (mgmt *VirtualSystemManagementService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualSystemManagementService", "Caption", "Hyper-V Virtual System Management Service")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := v2.NewMsvm_VirtualSystemManagementServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &VirtualSystemManagementService{vmmswmi}
	return
}

// GetVirtualMachines would get all virtual machines
func (vmms *VirtualSystemManagementService) GetVirtualMachines() (virtualsystem.VirtualMachineCollection, error) {
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
	return vmc, nil
}

// GetVirtualMachineByName
func (vmms *VirtualSystemManagementService) GetVirtualMachineByName(vmName string) (*virtualsystem.VirtualMachine, error) {
	vms, err := vmms.GetVirtualMachines()
	if err != nil {
		return nil, err
	}
	defer vms.Close()

	for _, vm := range vms {
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
	result, err := vmms.InvokeMethodWithReturn("DestroySystem", vm.InstancePath(), nil)
	//method, err := vmms.GetWmiMethod("DestroySystem")
	if err != nil {
		return err
	}

	//method.AddInParam("AffectedSystem", (vmms.GetIDispatch()))
	//result, err := method.Execute()
	//if err != nil {
	//	return err
	//}
	return vmms.WaitForJobCompletion(result, v2.ConcreteJob_JobType_Destroy_Virtual_Machine)
}

// ModifyVirtualSystemResource
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResource(data *wmi.WmiInstance) error {
	embeddedInstance, err := data.EmbeddedXMLInstance()
	if err != nil {
		return err
	}
	result, err := vmms.InvokeMethodWithReturn("ModifyResourceSettings", []string{embeddedInstance})
	if err != nil {
		return err
	}
	return vmms.WaitForJobCompletion(result, v2.ConcreteJob_JobType_Modify_Virtual_Machine_Resources)
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
		defer vmjob.Close()

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Virtual Machine Result[%d] JobType[%d]", result, jobType)
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
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return job.NewVirtualSystemJob(jobInstance)
}

func (vmms *VirtualSystemManagementService) RenameVirtualNetworkAdapter(vm *virtualsystem.VirtualMachine, adapterName, newName string) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()
	err = adapter.SetPropertyElementName(newName)
	if err != nil {
		return
	}

	return vmms.ModifyVirtualSystemResource(adapter.WmiInstance)
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterMACAddress(vm *virtualsystem.VirtualMachine, adapterName, macAddress string) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()
	err = adapter.SetPropertyAddress(macAddress)
	if err != nil {
		return
	}
	return vmms.ModifyVirtualSystemResource(adapter.WmiInstance)
}

func (vmms *VirtualSystemManagementService) ConnectAdapterToVirtualSwitch(vm *virtualsystem.VirtualMachine, adapterName string, virtSwitch *vswitch.VirtualSwitch) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	pasd, err := adapter.GetRelated("Msvm_EthernetPortAllocationSettingData")
	if err != nil {
		return
	}
	defer pasd.Close()
	err = pasd.SetProperty("EnabledState", 2)
	if err != nil {
		return
	}
	err = pasd.SetProperty("HostResource", []string{virtSwitch.InstancePath()})
	if err != nil {
		return
	}
	return vmms.ModifyVirtualSystemResource(pasd)
}

func (vmms *VirtualSystemManagementService) DisconnectAdapterFromVirtualSwitch(vm *virtualsystem.VirtualMachine, adapterName string) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	pasd, err := adapter.GetRelated("Msvm_EthernetPortAllocationSettingData")
	if err != nil {
		return
	}
	defer pasd.Close()
	err = pasd.SetProperty("EnabledState", 3)
	if err != nil {
		return
	}
	return vmms.ModifyVirtualSystemResource(pasd)
}
