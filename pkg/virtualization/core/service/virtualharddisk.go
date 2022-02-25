// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"
	"strings"

	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

func (vmms *VirtualSystemManagementService) AddSCSIController(vm *virtualsystem.VirtualMachine) (err error) {
	tmp, err := vm.NewSCSIController()
	if err != nil {
		return err
	}
	defer tmp.Close()

	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer vmsetting.Close()

	// apply the settings
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, tmp.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "AddVirtualSystemResource")
		return
	}

	return

}

// AttachVirtualHardDisk -
// * Create a Synthetic Disk Drive
// *    Add a drive to available first controller at available location
// * Connects the Disk to the Drive
// Returns Disk and Drive
func (vmms *VirtualSystemManagementService) AttachVirtualHardDisk(vm *virtualsystem.VirtualMachine, path string) (
	vhd *disk.VirtualHardDisk,
	vhddrive *drive.SyntheticDiskDrive,
	err error) {

	
	vhddrive, err = vmms.AddSyntheticDiskDrive(vm, -1, -1)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			log.Printf("[%+v]\n", err)
			// Remove the drive
			err1 := vmms.RemoveSyntheticDiskDrive(vhddrive)
			log.Printf("RemoveSyntheticDiskDrive [%+v]\n", err1)
			vhddrive.Close()
			vhddrive = nil
		}
	}()
	

	// Add a disk
	vhdtmp, err := vm.NewVirtualHardDisk(path)
	if err != nil {
		return
	}
	defer vhdtmp.Close()

	// Connect disk to drive
	err = vhdtmp.SetPropertyParent(vhddrive.InstancePath())
	if err != nil {
		return
	}

	if !strings.Contains(vhdtmp.InstancePath(), "Definition") {
		err = vmms.ModifyVirtualSystemResourceEx(vhdtmp.WmiInstance, -1)
		if err != nil {
			return
		}
	}

	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer vmsetting.Close()

	// apply the settings
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, vhdtmp.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		// Sometimes this could hapen - Find out why
		vhd, err = vm.GetVirtualHardDiskByPath(path)
		return
	}

	vhdInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	vhd, err = disk.NewVirtualHardDisk(vhdInstance)
	if err != nil {
		vhdInstance.Close()
	}
	return
}




func (vmms *VirtualSystemManagementService) DetachVirtualHardDisk(vhd *disk.VirtualHardDisk) (err error) {
	drive, err := vhd.GetDrive()
	if err != nil {
		return
	}
	defer drive.Close()

	// Remove Disk
	err1 := vmms.RemoveVirtualSystemResource(vhd.CIM_ResourceAllocationSettingData, -1)
	// Remove Drive
	err = vmms.RemoveVirtualSystemResource(drive.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	if err1 != nil {
		err = err1
		return
	}
	return
}

func (vmms *VirtualSystemManagementService) AddSyntheticDiskDrive(vm *virtualsystem.VirtualMachine,
	controllernumber,
	controllerlocation int32) (vhddrive *drive.SyntheticDiskDrive, err error) {
	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer vmsetting.Close()

	vhddrivetmp, err := vm.NewSyntheticDiskDrive(controllernumber, controllerlocation)
	if err != nil {
		return
	}
	defer vhddrivetmp.Close()
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, vhddrivetmp.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	defer resultcol.Close()
	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "AddVirtualSystemResource")
		return
	}
	driveInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	vhddrive, err = drive.NewSyntheticDiskDrive(driveInstance)

	return
}

func (vmms *VirtualSystemManagementService) RemoveSyntheticDiskDrive(
	vhddrive *drive.SyntheticDiskDrive) (err error) {
	err = vmms.RemoveVirtualSystemResource(vhddrive.CIM_ResourceAllocationSettingData, -1)
	return
}




