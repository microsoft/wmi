// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

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
	defer method.Close()

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
	defer method.Close()

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
