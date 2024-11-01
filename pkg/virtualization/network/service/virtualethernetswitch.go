// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"
	"time"

	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	netservice "github.com/microsoft/wmi/pkg/hardware/network/service"
	vmmshost "github.com/microsoft/wmi/pkg/virtualization/core/host"
	"github.com/microsoft/wmi/pkg/virtualization/network/ethernetport"
	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

func (vsms *VirtualEthernetSwitchManagementService) CreateExternalVirtualSwitch(physicalNicName, externalPortName, internalPortName string,
	settings *virtualswitch.VirtualEthernetSwitchSettingData) (
	vswitch *virtualswitch.VirtualSwitch,
	err error) {

	adapter, err := netservice.GetNetworkAdapterByName(vsms.GetWmiHost(), physicalNicName)
	if err != nil {
		return
	}
	defer adapter.Close()

	interfaceDescription, err := adapter.GetPropertyInterfaceDescription()
	if err != nil {
		return
	}

	internalpasd, err := vmmshost.GetInternalPortAllocationSettingData(vsms.GetWmiHost(), internalPortName)
	if err != nil {
		return
	}
	defer internalpasd.Close()

	externalpasd, err := vmmshost.GetExternalPortAllocationSettingData(vsms.GetWmiHost(), externalPortName, []string{interfaceDescription})
	if err != nil {
		return
	}
	defer internalpasd.Close()

	extPort, err := ethernetport.GetExternalEthernetPort(vsms.GetWmiHost(), interfaceDescription)
	if err != nil {
		return
	}
	defer extPort.Close()

	address, err := extPort.GetPropertyPermanentAddress()
	if err != nil {
		return
	}
	err = internalpasd.SetPropertyAddress(address)
	if err != nil {
		return
	}

	internalpasdXml, err := internalpasd.EmbeddedXMLInstance()
	if err != nil {
		return
	}
	externalpasdXml, err := externalpasd.EmbeddedXMLInstance()
	if err != nil {
		return
	}

	resourceSettings := []string{internalpasdXml, externalpasdXml}
	vswitch, err = vsms.CreateVirtualSwitch(settings, resourceSettings)
	return
}

func (vsms *VirtualEthernetSwitchManagementService) CreatePrivateVirtualSwitch(settings *virtualswitch.VirtualEthernetSwitchSettingData) (
	vswitch *virtualswitch.VirtualSwitch,
	err error) {
	vswitch, err = vsms.CreateVirtualSwitch(settings, []string{})
	return
}

func (vsms *VirtualEthernetSwitchManagementService) CreateInternalVirtualSwitch(internalPortName string, settings *virtualswitch.VirtualEthernetSwitchSettingData) (
	vswitch *virtualswitch.VirtualSwitch,
	err error) {
	internalpasd, err := vmmshost.GetInternalPortAllocationSettingData(vsms.GetWmiHost(), internalPortName)
	if err != nil {
		return
	}
	defer internalpasd.Close()
	internalpasdXml, err := internalpasd.EmbeddedXMLInstance()
	if err != nil {
		return
	}

	vswitch, err = vsms.CreateVirtualSwitch(settings, []string{internalpasdXml})
	return
}

func (vsms *VirtualEthernetSwitchManagementService) CreateVirtualSwitch(settings *virtualswitch.VirtualEthernetSwitchSettingData, resourceSettings []string) (
	vswitch *virtualswitch.VirtualSwitch,
	err error) {

	method, err := vsms.GetWmiMethod("DefineSystem")
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
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", resourceSettings))
	inparams = append(inparams, wmi.NewWmiMethodParam("ReferenceConfiguration", nil))
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingSystem", nil))

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if !(result.ReturnValue == 4096 || result.ReturnValue == 0) {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}
	val, ok := result.OutMethodParams["ResultingSystem"]
	if ok && val.Value != nil {
		vswitchinstance, err := instance.GetWmiInstanceFromPath(vsms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err == nil {
			vswitch, err = virtualswitch.NewVirtualSwitch(vswitchinstance)
		}
	}

	if result.ReturnValue == 0 {
		return
	}

	val, ok = result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(vsms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue, -1)

	if vswitch != nil {
		return
	}

	affectedElement, err := job.GetFirstRelatedEx("Msvm_AffectedJobElement", "", "", "")
	if err != nil {
		// For now, ignore the error
		err = nil
		return
	}
	vswitch, err = virtualswitch.NewVirtualSwitch(affectedElement)
	return
}

func (vsms *VirtualEthernetSwitchManagementService) DeleteVirtualSwitch(vswitch *virtualswitch.VirtualSwitch) (err error) {
	method, err := vsms.GetWmiMethod("DestroySystem")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedSystem", vswitch.InstancePath()))
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	for {
		result, err1 := method.Execute(inparams, outparams)
		if err1 != nil {
			err = err1
			return
		}

		returnVal := result.ReturnValue
		if returnVal != 0 && returnVal != 4096 {
			// Virtual System is in Invalid State, try to retry
			if returnVal == 32775 {
				log.Printf("[WMI] Method [%s] failed with [%d]. Retrying ...", method.Name, returnVal)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
			return
		}

		if result.ReturnValue == 0 {
			return
		}

		val, ok := result.OutMethodParams["Job"]
		if !ok || val.Value == nil {
			err = errors.Wrapf(errors.NotFound, "Job")
			return
		}
		job, err1 := instance.GetWmiJob(vsms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			err = err1
			return
		}
		defer job.Close()
		return job.WaitForJobCompletion(result.ReturnValue, -1)
	}
}
