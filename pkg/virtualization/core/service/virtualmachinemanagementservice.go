// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"
	"sync"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmihost "github.com/microsoft/wmi/pkg/virtualization/core/host"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
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
	query := query.NewWmiQuery("Msvm_VirtualSystemManagementService")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := v2.NewMsvm_VirtualSystemManagementServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &VirtualSystemManagementService{vmmswmi}
	return
}

func (vmms *VirtualSystemManagementService) GetHostComputerSystem() (*wmihost.HostComputerSystem, error) {
	return wmihost.GetHostComputerSystem(vmms.GetWmiHost())
}

//
func (vmms *VirtualSystemManagementService) AddVirtualSystemResource(
	vmsettings *virtualsystem.VirtualSystemSettingData,
	data *v2.CIM_ResourceAllocationSettingData,
	timeoutSeconds uint16) (resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstance, err := data.EmbeddedXMLInstance()
	if err != nil {
		return
	}

	method, err := vmms.GetWmiMethod("AddResourceSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedConfiguration", vmsettings.InstancePath()))
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", []string{embeddedInstance}))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingResourceSettings", nil))

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
		val := result.OutMethodParams["ResultingResourceSettings"]
		if val.Value != nil {
			for _, resultingVal := range val.Value.([]interface{}) {
				inst, err1 := instance.GetWmiInstanceFromPath(vmms.GetWmiHost(), string(constant.Virtualization), resultingVal.(string))
				if err1 != nil {
					log.Printf("[WMI] ResultingResourceSettings - GetInstanceFromPath [%+v]\n", err1)
					//err = err1
					continue
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
		err = job.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
		return
	}
	return
}

func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResourceEx(data *wmi.WmiInstance, timeoutSeconds uint16) (err error) {
	result, err := vmms.ModifyVirtualSystemResource(wmi.WmiInstanceCollection{data}, timeoutSeconds)
	if err != nil {
		return
	}
	result.Close()
	return
}
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResourceEx2(data *wmi.WmiInstance, timeoutSeconds uint16) (instance *wmi.WmiInstance, err error) {
	result, err := vmms.ModifyVirtualSystemResource(wmi.WmiInstanceCollection{data}, timeoutSeconds)
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
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemResource(data wmi.WmiInstanceCollection, timeoutSeconds uint16) (
	resultingResources wmi.WmiInstanceCollection, err error) {

	embeddedInstance, err := data.EmbeddedXMLInstances()
	if err != nil {
		return
	}
	method, err := vmms.GetWmiMethod("ModifyResourceSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", embeddedInstance))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	outparams = append(outparams, wmi.NewWmiMethodParam("ResultingResourceSettings", nil))

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
		job, err1 := instance.GetWmiJob(vmms.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			err = err1
			return
		}
		defer job.Close()
		err = job.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
		return
	}
	return
}

// RemoveVirtualSystemResource - Will be removed, when auto gen code is regenerated
func (vmms *VirtualSystemManagementService) RemoveVirtualSystemResource(
	data *v2.CIM_ResourceAllocationSettingData, timeoutSeconds uint16) (err error) {
	method, err := vmms.GetWmiMethod("RemoveResourceSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("ResourceSettings", []string{data.InstancePath()}))

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

		// Try to get the Out Params
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

		err = job.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
		return
	}
	return
}

func (vmms *VirtualSystemManagementService) ModifyVirtualSystemFeatureEx(data *wmi.WmiInstance, timeoutSeconds uint16) (err error) {
	result, err := vmms.ModifyVirtualSystemFeature(wmi.WmiInstanceCollection{data}, timeoutSeconds)
	if err != nil {
		return
	}
	result.Close()
	return
}

// ModifyVirtualSystemFeature
func (vmms *VirtualSystemManagementService) ModifyVirtualSystemFeature(data wmi.WmiInstanceCollection, timeoutSeconds uint16) (
	resultingFeatures wmi.WmiInstanceCollection, err error) {

	embeddedInstance, err := data.EmbeddedXMLInstances()
	if err != nil {
		return
	}
	method, err := vmms.GetWmiMethod("ModifyFeatureSettings")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("FeatureSettings", embeddedInstance))

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
				resultingFeatures = append(resultingFeatures, inst)
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
		err = job.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
		return
	}
	return
}
