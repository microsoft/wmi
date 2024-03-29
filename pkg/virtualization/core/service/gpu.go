// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"

	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

func (vmms *VirtualSystemManagementService) AttachGPU(vm *virtualsystem.VirtualMachine, hostResource string) (err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		log.Printf("[WMI] Error getting VirtualSystemSettingData for virtual machine [%s] - Error details [%+v]\n", vm, err)
		return
	}
	defer vmsetting.Close()

	gpu, err := vm.NewPcieDevice(hostResource)
	if err != nil {
		log.Printf("[WMI] Error getting new GPU for host resource [%s] - Error details [%+v]\n", hostResource, err)
		return err
	}
	defer gpu.Close()

	// Attach the GPU
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		log.Printf("[WMI] Error attaching GPU with resource allocation setting data [%s] - Error details [%+v]\n", gpu.CIM_ResourceAllocationSettingData, err)
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "Unable to attach GPU with host resource [%s] to VM [%s]", hostResource, vm.Name())
		log.Printf("[WMI] Virtual system resource not found - Error details [%+v]\n", err)
		return
	}

	log.Printf("[WMI] Successfully attached GPU with host resource [%s] to the virtual machine [%s]", hostResource, vm.Name())
	return
}

func (vmms *VirtualSystemManagementService) DetachGPU(vm *virtualsystem.VirtualMachine, hostResource string) (err error) {
	gpu, err := vm.GetPcieDevice(hostResource)
	if err != nil {
		log.Printf("[WMI] Error getting GPU device for host resource [%s] - Error details [%+v]\n", hostResource, err)
		if errors.IsNotFound(err) {
			err = errors.NotFound
		}
		return err
	}
	defer gpu.Close()

	// Detach the GPU
	err = vmms.RemoveVirtualSystemResource(gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		log.Printf("[WMI] Error detaching GPU with resource allocation setting data [%s] - Error details [%+v]\n", gpu.CIM_ResourceAllocationSettingData, err)
		return
	}

	log.Printf("[WMI] Successfully detached GPU with host resource [%s] from the virtual machine [%s]", hostResource, vm.Name())
	return
}
