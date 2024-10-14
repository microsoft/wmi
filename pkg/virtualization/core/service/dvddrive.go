// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"

	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

func (vmms *VirtualSystemManagementService) AddISODisk(vm *virtualsystem.VirtualMachine, isoPath string) (
	ld *disk.LogicalDisk,
	dvddrive *drive.DvdDrive,
	err error) {

	dvddrive, err = vmms.AddDvdDrive(vm)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			log.Printf("[%+v]\n", err)
			err1 := vmms.RemoveDvdDrive(dvddrive)
			log.Printf("RemoveDvdDrive [%+v]\n", err1)
			dvddrive.Close()
			dvddrive = nil
		}
	}()

	// create the logical disk
	tmpld, err := vm.NewLogicalDisk()
	if err != nil {
		return
	}
	defer tmpld.Close()

	// set the host resource to isoPath, parent to the DVD drive and resourcesubtype to cd/dvd disk
	err = tmpld.SetPropertyHostResource([]string{isoPath})
	if err != nil {
		return
	}
	err = tmpld.SetPropertyParent(dvddrive.InstancePath())
	if err != nil {
		return
	}
	err = tmpld.SetProperty("ResourceSubType", "Microsoft:Hyper-V:Virtual CD/DVD Disk")
	if err != nil {
		return
	}

	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer vmsetting.Close()

	// apply the settings
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, tmpld.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "AddVirtualSystemResource")
		return
	}

	ldInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	ld, err = disk.NewLogicalDisk(ldInstance)
	if err != nil {
		ldInstance.Close()
	}
	return

}

func (vmms *VirtualSystemManagementService) GetISODisk(vm *virtualsystem.VirtualMachine, isoPath string) (ld *disk.LogicalDisk, err error) {
	// col, err := vm.GetDvdDrives()
	// if err != nil {
	// 	return
	// }
	// defer col.Close()

	// TODO: probably put this find functionality somewhere else
	// for _, inst := range col {
	// 	dvddrive, err := drive.NewDvdDrive(inst)
	// 	if err != nil {
	// 		return nil, err
	// }
	// Get the logical disk
	// ld, err := disk.GetLogicalDiskByHostResource(vm, isoPath)
	// if err != nil {
	// 	return nil
	// }
	return nil, errors.NotImplemented
}

func (vmms *VirtualSystemManagementService) RemoveISODisk(ld *disk.LogicalDisk) (err error) {
	dvddrive, err := ld.GetDrive()
	if err != nil {
		return
	}
	defer dvddrive.Close()

	// Remove Disk
	err1 := vmms.RemoveVirtualSystemResource(ld.CIM_ResourceAllocationSettingData, -1)
	// Remove Drive
	err = vmms.RemoveVirtualSystemResource(dvddrive.CIM_ResourceAllocationSettingData, -1)
	if err != nil {
		return
	}
	if err1 != nil {
		err = err1
		return
	}
	return
}

func (vmms *VirtualSystemManagementService) AddDvdDrive(vm *virtualsystem.VirtualMachine) (dvd *drive.DvdDrive, err error) {
	// Add the dvd drive
	tmp, err := vm.NewDvdDrive()
	if err != nil {
		return
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

	dvdInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	dvd, err = drive.NewDvdDrive(dvdInstance)
	return
}

func (vmms *VirtualSystemManagementService) RemoveDvdDrive(dvd *drive.DvdDrive) (err error) {
	err = vmms.RemoveVirtualSystemResource(dvd.CIM_ResourceAllocationSettingData, -1)
	return
}
