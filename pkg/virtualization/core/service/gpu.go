// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

func (vmms *VirtualSystemManagementService) AttachGpuDda(vm *virtualsystem.VirtualMachine, hostResource string) (err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return errors.Wrapf(err, "Failed to get VirtualSystemSettingData for virtual machine [%+v]", vm)
	}
	defer vmsetting.Close()

	gpu, err := vm.NewPcieDevice(hostResource)
	if err != nil {
		return errors.Wrapf(err, "Failed to get new GPU for host resource [%s]", hostResource)
	}
	defer gpu.Close()

	// Attach the GPU
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to attach GPU with resource allocation setting data [%+v] to VM [%+v]", gpu.CIM_ResourceAllocationSettingData, vm)
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		return errors.Wrapf(errors.NotFound, "Unable to find virtual system resource after GPU-DDA attach operation on VM [%+v] with host resource [%s]", vm, hostResource)
	}

	return
}

func (vmms *VirtualSystemManagementService) DetachGpuDda(vm *virtualsystem.VirtualMachine, hostResource string) (err error) {
	gpu, err := vm.GetPcieDevice(hostResource)
	if err != nil {
		return errors.Wrapf(err, "Failed to get GPU device for host resource [%s] on VM [%+v]", hostResource, vm)
	}
	defer gpu.Close()

	// Detach the GPU
	err = vmms.RemoveVirtualSystemResource(gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to detach GPU with resource allocation setting data [%+v] from VM [%+v]", gpu.CIM_ResourceAllocationSettingData, vm)
	}

	return
}

func (vmms *VirtualSystemManagementService) AttachGpuP(vm *virtualsystem.VirtualMachine, partitionSizeBytes uint64) (err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return errors.Wrapf(err, "Failed to get VirtualSystemSettingData for virtual machine [%+v]", vm)
	}
	defer vmsetting.Close()

	gpu, err := vm.NewGpuPartition(partitionSizeBytes)
	if err != nil {
		return errors.Wrapf(err, "Failed to get new GPU partition of size [%d bytes] for VM [%+v]", partitionSizeBytes, vm)
	}
	defer gpu.Close()

	// Attach the GPU Partition
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, gpu.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to attach GPU partition with resource allocation setting data [%+v] to VM [%+v]", gpu.CIM_ResourceAllocationSettingData, vm)
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		return errors.Wrapf(errors.NotFound, "Unable to find virtual system resource after GPU-P attach operation on VM [%+v] with partition size [%d bytes]", vm, partitionSizeBytes)
	}

	return
}

func (vmms *VirtualSystemManagementService) DetachGpuP(vm *virtualsystem.VirtualMachine, partitionSizeBytes uint64) (err error) {
	partitionSettingData, err := vm.GetGpuPartitionSettingData(partitionSizeBytes)
	if err != nil {
		return errors.Wrapf(err, "Failed to get GPU partition of size [%d bytes] for VM [%+v]", partitionSizeBytes, vm)
	}
	defer partitionSettingData.Close()

	// Detach the GPU Partition
	err = vmms.RemoveVirtualSystemResource(partitionSettingData.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to detach GPU partition with resource allocation setting data [%+v] from VM [%+v]", partitionSettingData.CIM_ResourceAllocationSettingData, vm)
	}

	return
}
