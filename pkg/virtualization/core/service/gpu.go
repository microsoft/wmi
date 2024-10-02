// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/gpu"
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
	gpu, err := vm.NewGpuPartition(partitionSizeBytes)
	if err != nil {
		return errors.Wrapf(err, "Failed to get new GPU partition of size [%d bytes] for VM [%+v]", partitionSizeBytes, vm)
	}
	defer gpu.Close()

	err = vmms.AttachGpuPartitionToVM(vm, gpu)
	return
}

func (vmms *VirtualSystemManagementService) AttachDefaultGpuP(vm *virtualsystem.VirtualMachine) (err error) {
	gpu, err := vm.NewDefaultGpuPartition()
	if err != nil {
		return errors.Wrapf(err, "Failed to get new default GPU partition of size VM [%+v]", vm)
	}
	defer gpu.Close()

	err = vmms.AttachGpuPartitionToVM(vm, gpu)
	return
}

func (vmms *VirtualSystemManagementService) AttachGpuPartitionToVM(vm *virtualsystem.VirtualMachine, gpuPartition *gpu.GpuPartitionSettingData) (err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return errors.Wrapf(err, "Failed to get VirtualSystemSettingData for virtual machine [%+v]", vm)
	}
	defer vmsetting.Close()

	// Attach the GPU Partition
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, gpuPartition.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to attach GPU partition with resource allocation setting data [%+v] to VM [%+v]", gpuPartition.CIM_ResourceAllocationSettingData, vm)
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		return errors.Wrapf(errors.NotFound, "Unable to find virtual system resource after GPU-P attach operation on VM [%+v]", vm)
	}

	return
}

func (vmms *VirtualSystemManagementService) DetachGpuP(vm *virtualsystem.VirtualMachine, partitionSizeBytes uint64) (err error) {
	if partitionSizeBytes == 0 {
		return errors.Wrapf(errors.InvalidInput, "Detach GpuP with 0 partition size is not supported.")
	}

	partitionSettingData, err := vm.GetGpuPartitionSettingData(partitionSizeBytes)
	if err != nil {
		return errors.Wrapf(err, "Failed to get GPU partition of size [%d bytes] for VM [%+v]", partitionSizeBytes, vm)
	}
	defer partitionSettingData.Close()

	// Detach the GPU Partition
	err = vmms.DetachGpuPartitionFromVM(vm, partitionSettingData)
	return
}

func (vmms *VirtualSystemManagementService) DetachDefaultGpuP(vm *virtualsystem.VirtualMachine) (err error) {
	partitionSettingData, err := vm.GetDefaultGpuPartitionSettingData()
	if err != nil {
		return errors.Wrapf(err, "Failed to get GPU partition for VM [%+v]", vm)
	}
	defer partitionSettingData.Close()

	err = vmms.DetachGpuPartitionFromVM(vm, partitionSettingData)
	return
}

func (vmms *VirtualSystemManagementService) DetachGpuPartitionFromVM(vm *virtualsystem.VirtualMachine, gpuPartition *gpu.GpuPartitionSettingData) (err error) {
	err = vmms.RemoveVirtualSystemResource(gpuPartition.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return errors.Wrapf(err, "Failed to detach GPU partition with resource allocation setting data [%+v] from VM [%+v]", gpuPartition.CIM_ResourceAllocationSettingData, vm)
	}

	return
}
