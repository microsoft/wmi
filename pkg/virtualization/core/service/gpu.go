// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"

	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

func (vmms *VirtualSystemManagementService) AttachGPU(vm *virtualsystem.VirtualMachine, locationPath string) (err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		log.Printf("[WMI] Error getting VirtualSystemSettingData for virtual machine [%s]", vm)
		return
	}
	defer vmsetting.Close()

	gpu, err := vm.GetPcieSettingByLocationPath(locationPath)
	if err != nil {
		log.Printf("[WMI] Error getting new GPU for location path [%s]", locationPath)
		return err
	}
	defer gpu.Close()

	// Attach the GPU
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		log.Printf("[WMI] Error attaching GPU with resource allocation setting data [%s]", gpu.CIM_ResourceAllocationSettingData)
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "Unable to attach GPU at location [%s] to VM [%s]", locationPath, vm.Name())
		log.Printf("[WMI] Virtual system resource not found")
		return
	}

	log.Printf("[WMI] Successfully attached GPU with location path [%s] to the virtual machine", locationPath)
	return
}

func (vmms *VirtualSystemManagementService) DetachGPU(vm *virtualsystem.VirtualMachine, locationPath string) (err error) {
	gpu, err := vm.GetPcieDeviceByLocationPath(locationPath)
	if err != nil {
		log.Printf("[WMI] Error getting GPU device for location path [%s]", locationPath)
		if errors.IsNotFound(err) {
			err = errors.NotFound
		}
		return err
	}
	defer gpu.Close()

	// Detach the GPU
	err = vmms.RemoveVirtualSystemResource(gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		log.Printf("[WMI] Error detaching GPU with resource allocation setting data [%s]", gpu.CIM_ResourceAllocationSettingData)
		return
	}

	log.Printf("[WMI] Successfully detached GPU with location path [%s] from the virtual machine", locationPath)
	return
}
