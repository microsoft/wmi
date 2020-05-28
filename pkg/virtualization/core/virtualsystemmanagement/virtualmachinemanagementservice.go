// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

import (
	"log"
	"strings"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
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
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt.Close()
		mgmt = val
		return
	}

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

func (vmms *VirtualSystemManagementService) CreateVirtualMachine(settings *virtualsystem.VirtualSystemSettingData) (
	vm *virtualsystem.VirtualMachine,
	err error) {

	method, err := vmms.GetWmiMethod("DefineSystem")
	if err != nil {
		return
	}

	embeddedInstance, err := settings.EmbeddedXMLInstance()
	if err != nil {
		return
	}

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("SystemSettings", embeddedInstance))
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", nil))
	inparams = append(inparams, wmi.NewWmiMethodParam("ReferenceConfiguration", nil))
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingSystem", nil))

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	if result.ReturnValue != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}
	val, ok := result.OutMethodParams["ResultingSystem"]
	if ok {
		vminstance, err := instance.GetWmiInstanceFromPath(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err == nil {
			vm, err = virtualsystem.NewVirtualMachine(vminstance)
			//
		}
	}

	val, ok = result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue)
	return
}
func (vmms *VirtualSystemManagementService) DeleteVirtualMachine(vm *virtualsystem.VirtualMachine) (err error) {
	method, err := vmms.GetWmiMethod("DestroySystem")
	if err != nil {
		return
	}

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedSystem", vm.InstancePath()))
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	if result.ReturnValue != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	val, ok := result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	return job.WaitForJobCompletion(result.ReturnValue)
}

//
func (vmms *VirtualSystemManagementService) AddVirtualSystemResource(
	vmsettings *virtualsystem.VirtualSystemSettingData,
	data *v2.CIM_ResourceAllocationSettingData) (resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstance, err := data.EmbeddedXMLInstance()
	if err != nil {
		return
	}

	method, err := vmms.GetWmiMethod("AddResourceSettings")
	if err != nil {
		return
	}

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedConfiguration", vmsettings.InstancePath()))
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", []string{embeddedInstance}))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingResourceSettings", nil))

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	returnVal := result.ReturnValue
	if returnVal != 0 && returnVal != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	// Try to get the Out Params
	val := result.OutMethodParams["ResultingResourceSettings"]
	if val.Value != nil {
		for _, resultingVal := range val.Value.([]interface{}) {
			inst, err1 := instance.GetWmiInstanceFromPath(vmms.GetWmiHost(), string(constant.Virtualization), resultingVal.(string))
			if err1 != nil {
				err = err1
				return
			}
			resultingResources = append(resultingResources, inst)
		}
	}
	if result.ReturnValue == 0 {
		return
	}

	val = result.OutMethodParams["Job"]
	job, err := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue)
	return
}

func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResourceEx(data *wmi.WmiInstance) (err error) {
	result, err := vmms.ModifyVirtualSystemResource(wmi.WmiInstanceCollection{data})
	if err != nil {
		return
	}
	result.Close()
	return
}
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResourceEx2(data *wmi.WmiInstance) (instance *wmi.WmiInstance, err error) {
	result, err := vmms.ModifyVirtualSystemResource(wmi.WmiInstanceCollection{data})
	if err != nil {
		return
	}
	result.Close()
	if len(result) == 0 {
		err = errors.Wrapf(errors.NotFound, "ModifyVirtualSystemResource didnt not return any resulting system")
		return
	}
	instance, err = result[0].Clone()
	return
}

// ModifyVirtualSystemResource
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResource(data wmi.WmiInstanceCollection) (
	resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstance, err := data.EmbeddedXMLInstances()
	if err != nil {
		return
	}
	method, err := vmms.GetWmiMethod("ModifyResourceSettings")
	if err != nil {
		return
	}

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", embeddedInstance))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingResourceSettings", nil))

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	returnVal := result.ReturnValue
	if returnVal != 0 && returnVal != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	// Try to get the Out Params
	val := result.OutMethodParams["ResultingResourceSettings"]
	if val.Value != nil {
		for _, resultingVal := range val.Value.([]interface{}) {
			inst, err1 := instance.GetWmiInstanceFromPath(vmms.GetWmiHost(), string(constant.Virtualization), resultingVal.(string))
			if err1 != nil {
				err = err1
				return
			}
			resultingResources = append(resultingResources, inst)
		}
	}
	if result.ReturnValue == 0 {
		return
	}

	val = result.OutMethodParams["Job"]
	job, err := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue)
	return

}

// RemoveVirtualSystemResource - Will be removed, when auto gen code is regenerated
func (vmms *VirtualSystemManagementService) RemoveVirtualSystemResource(
	data *v2.CIM_ResourceAllocationSettingData) (err error) {

	//embeddedInstance, err := data.EmbeddedXMLInstance()
	//if err != nil {
	//	return
	//}

	method, err := vmms.GetWmiMethod("RemoveResourceSettings")
	if err != nil {
		return
	}

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", []string{data.InstancePath()}))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	returnVal := result.ReturnValue
	if returnVal != 0 && returnVal != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	// Try to get the Out Params
	if result.ReturnValue == 0 {
		return
	}

	val := result.OutMethodParams["Job"]
	job, err := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()

	err = job.WaitForJobCompletion(result.ReturnValue)
	return
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

	return vmms.ModifyVirtualSystemResourceEx(adapter.WmiInstance)
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
	return vmms.ModifyVirtualSystemResourceEx(adapter.WmiInstance)
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
	return vmms.ModifyVirtualSystemResourceEx(pasd)
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
	return vmms.ModifyVirtualSystemResourceEx(pasd)
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

	// Add a drive
	vhddrive, err = vmms.AddSyntheticDiskDrive(vm, -1, -1)
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			return
		}
		log.Printf("[%+v]\n", err)
		// Remove the drive
		err1 := vmms.RemoveSyntheticDiskDrive(vhddrive)
		log.Printf("RemoveSyntheticDiskDrive [%+v]\n", err1)
		vhddrive.Close()
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
		err = vmms.ModifyVirtualSystemResourceEx(vhdtmp.WmiInstance)
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
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, vhdtmp.CIM_ResourceAllocationSettingData)
	if err != nil {
		return
	}
	defer resultcol.Close()

	if len(resultcol) == 0 {
		err = errors.Wrapf(errors.NotFound, "AddVirtualSystemResource")
		return
	}

	vhdInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	vhd, err = disk.NewVirtualHardDisk(vhdInstance)
	return
}

func (vmms *VirtualSystemManagementService) DetachVirtualHardDisk(vhd *disk.VirtualHardDisk) (err error) {
	// Remove Disk
	err = vmms.RemoveVirtualSystemResource(vhd.CIM_ResourceAllocationSettingData)
	if err != nil {
		return
	}
	// Remove Drive
	drive, err := vhd.GetDrive()
	err = vmms.RemoveVirtualSystemResource(drive.CIM_ResourceAllocationSettingData)
	if err != nil {
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
	resultcol, err := vmms.AddVirtualSystemResource(vmsetting, vhddrivetmp.CIM_ResourceAllocationSettingData)
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
	err = vmms.RemoveVirtualSystemResource(vhddrive.CIM_ResourceAllocationSettingData)
	return
}
