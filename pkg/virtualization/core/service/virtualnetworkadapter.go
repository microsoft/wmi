// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"
	"time"

	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	vswitch "github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

func (vmms *VirtualSystemManagementService) AddVirtualEthernetConnection(vm *virtualsystem.VirtualMachine, vna *na.VirtualNetworkAdapter) (
	epas *switchport.EthernetPortAllocationSettingData, err error) {

	tmp, err := vm.NewEthernetPortAllocationSettingData(vna)
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

	resultInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}
	return switchport.NewEthernetPortAllocationSettingData(resultInstance)
}

func (vmms *VirtualSystemManagementService) AddVirtualNetworkAdapter(vm *virtualsystem.VirtualMachine, name string) (vna *na.VirtualNetworkAdapter, err error) {
	// Add a adapter
	tmp, err := vm.NewSyntheticNetworkAdapter(name)
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

	resultInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	vna, err = na.NewVirtualNetworkAdapter(resultInstance)
	if err != nil {
		resultInstance.Close()
	}
	return

}

func (vmms *VirtualSystemManagementService) AddVirtualNetworkAdapterWithMac(vm *virtualsystem.VirtualMachine, name, macAddress string) (vna *na.VirtualNetworkAdapter, err error) {
	// Add a adapter
	tmp, err := vm.NewSyntheticNetworkAdapter(name)
	if err != nil {
		return
	}
	defer tmp.Close()

	vmsetting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer vmsetting.Close()

	err = tmp.SetProperty("StaticMacAddress", true)
	if err != nil {
		return
	}
	err = tmp.SetPropertyAddress(macAddress)
	if err != nil {
		return
	}

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

	resultInstance, err := resultcol[0].Clone()
	if err != nil {
		return
	}

	vna, err = na.NewVirtualNetworkAdapter(resultInstance)
	if err != nil {
		resultInstance.Close()
	}
	return

}

func (vmms *VirtualSystemManagementService) RemoveVirtualNetworkAdapter(vna *na.VirtualNetworkAdapter) (err error) {
	// Remove adapter
	err = vmms.RemoveVirtualSystemResource(vna.CIM_ResourceAllocationSettingData, -1)
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

	return vmms.ModifyVirtualSystemResourceEx(adapter.WmiInstance, -1)
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterMACAddress(vm *virtualsystem.VirtualMachine, adapterName, macAddress string) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	err = adapter.SetProperty("StaticMacAddress", true)
	if err != nil {
		return
	}
	err = adapter.SetPropertyAddress(macAddress)
	if err != nil {
		return
	}
	err = vmms.ModifyVirtualSystemResourceEx(adapter.WmiInstance, -1)
	return err
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterIOVOffloadWeight(vm *virtualsystem.VirtualMachine, adapterName string, iovOffloadWeight uint32) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	netPortAllocSettingData, err := adapter.GetEthernetPortAllocationSettingData()
	if err != nil {
		return
	}
	defer netPortAllocSettingData.Close()

	portOffloadSettingData, err := netPortAllocSettingData.GetEthernetSwitchPortOffloadSettingData()
	if err != nil {
		return
	}
	defer portOffloadSettingData.Close()

	err = portOffloadSettingData.SetPropertyIOVOffloadWeight(iovOffloadWeight)
	if err != nil {
		return
	}

	err = vmms.ModifyVirtualSystemFeatureEx(portOffloadSettingData.WmiInstance, -1)
	return err
}

func (vmms *VirtualSystemManagementService) ConnectAdapterToVirtualSwitch(vm *virtualsystem.VirtualMachine, adapterName string, virtSwitch *vswitch.VirtualSwitch) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	// If the adapter is already connected, then Msvm_EthernetPortAllocationSettingData would exists, if not it would be null
	pasd, err := adapter.GetEthernetPortAllocationSettingData()
	if err != nil {
		if !errors.IsNotFound(err) {
			return
		}
		pasd, err = vmms.AddVirtualEthernetConnection(vm, adapter)
		if err != nil {
			return
		}
	} else {
		// Already existing case
		err = pasd.SetProperty("EnabledState", 2)
		if err != nil {
			pasd.Close()
			return
		}
	}
	defer pasd.Close()
	err = pasd.SetProperty("HostResource", []string{virtSwitch.InstancePath()})
	if err != nil {
		return
	}
	return vmms.ModifyVirtualSystemResourceEx(pasd.WmiInstance, -1)
}

func (vmms *VirtualSystemManagementService) DisconnectAdapterFromVirtualSwitch(vm *virtualsystem.VirtualMachine, adapterName string) (err error) {
	adapter, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		return
	}
	defer adapter.Close()

	pasd, err := adapter.GetEthernetPortAllocationSettingData()
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return
	}
	defer pasd.Close()
	err = pasd.SetProperty("EnabledState", 3)
	if err != nil {
		return
	}
	return vmms.ModifyVirtualSystemResourceEx(pasd.WmiInstance, -1)
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterPortProfile(
	vna *na.VirtualNetworkAdapter, profileName, vendorGuid, profileGuid string, profileData uint32) (err error) {

	hc, err := vmms.GetHostComputerSystem()
	if err != nil {
		return
	}
	defer hc.Close()

	profileSettings, err := hc.GetDefaultEthernetSwitchPortProfileSettingData(profileName, vendorGuid, profileGuid, profileData)
	if err != nil {
		return
	}
	defer profileSettings.Close()
	pasd, err := vna.GetEthernetPortAllocationSettingData()
	if err != nil {
		return
	}
	defer pasd.Close()

	return vmms.AddEthernetFeatureEx1(pasd.Msvm_EthernetPortAllocationSettingData, profileSettings.WmiInstance, -1)
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterPortIsolation(
	vna *na.VirtualNetworkAdapter, vlanId uint16) (err error) {
	return errors.NotImplemented
}

func (vmms *VirtualSystemManagementService) SetVirtualNetworkAdapterAccessVLAN(vna *na.VirtualNetworkAdapter, vlanId uint16) (err error) {
	hc, err := vmms.GetHostComputerSystem()
	if err != nil {
		return
	}
	defer hc.Close()

	vlanSettings, err := hc.GetDefaultEthernetSwitchPortVLANSettingData(vlanId)
	if err != nil {
		return
	}
	defer vlanSettings.Close()

	pasd, err := vna.GetEthernetPortAllocationSettingData()
	if err != nil {
		return
	}
	defer pasd.Close()

	return vmms.AddEthernetFeatureEx1(pasd.Msvm_EthernetPortAllocationSettingData, vlanSettings.WmiInstance, -1)
}

func (vmms *VirtualSystemManagementService) AddEthernetFeatureEx1(
	settings *v2.Msvm_EthernetPortAllocationSettingData,
	featureSetting *wmi.WmiInstance,
	timeoutSeconds int16) (err error) {

	result, err := vmms.AddEthernetFeature(settings, wmi.WmiInstanceCollection{featureSetting}, timeoutSeconds)
	if err != nil {
		return
	}
	defer result.Close()

	return
}

func (vmms *VirtualSystemManagementService) AddEthernetFeature(
	settings *v2.Msvm_EthernetPortAllocationSettingData,
	col wmi.WmiInstanceCollection,
	timeoutSeconds int16) (resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstances, err := col.EmbeddedXMLInstances()
	if err != nil {
		return
	}

	method, err := vmms.GetWmiMethod("AddFeatureSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedConfiguration", settings.InstancePath()))
	inparams = append(inparams, wmi.NewWmiMethodParam("FeatureSettings", embeddedInstances))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingFeatureSettings", nil))

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

		// Try to get the Out Params
		val := result.OutMethodParams["ResultingFeatureSettings"]
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
		job, err1 := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			err = err1
			return
		}
		defer job.Close()
		err = job.WaitForJobCompletion(result.ReturnValue, -1)
		break
	}
	return
}

func (vmms *VirtualSystemManagementService) ModifyEthernetFeature(
	col wmi.WmiInstanceCollection,
	timeoutSeconds int16) (resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstances, err := col.EmbeddedXMLInstances()
	if err != nil {
		return
	}

	method, err := vmms.GetWmiMethod("ModifyFeatureSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("FeatureSettings", embeddedInstances))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingFeatureSettings", nil))

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

		// Try to get the Out Params
		val := result.OutMethodParams["ResultingFeatureSettings"]
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
		job, err1 := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			err = err1
			return
		}
		defer job.Close()
		err = job.WaitForJobCompletion(result.ReturnValue, -1)
		break
	}
	return
}

func (vmms *VirtualSystemManagementService) RemoveEthernetFeature(
	col wmi.WmiInstanceCollection,
	timeoutSeconds int16) (err error) {

	embeddedInstances, err := col.EmbeddedXMLInstances()
	if err != nil {
		return
	}

	method, err := vmms.GetWmiMethod("RemoveFeatureSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("FeatureSettings", embeddedInstances))

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

		val := result.OutMethodParams["Job"]
		job, err1 := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			err = err1
			return
		}
		defer job.Close()
		err = job.WaitForJobCompletion(result.ReturnValue, -1)
		break
	}
	return
}
