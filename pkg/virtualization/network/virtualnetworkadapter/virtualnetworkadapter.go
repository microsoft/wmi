// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	"fmt"
	"log"

	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualNetworkAdapter struct {
	*EthernetPortAllocationSettingData
}

// NewVirtualNetworkAdapter
func NewVirtualNetworkAdapter(instance *wmi.WmiInstance) (*VirtualNetworkAdapter, error) {
	wmivm, err := NewEthernetPortAllocationSettingData(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkAdapter{wmivm}, nil
}

func (vm *VirtualNetworkAdapter) SetMACAddress(mac string) (err error) {
	return vm.SetPropertyAddress(mac)
}

func (vm *VirtualNetworkAdapter) Rename(newName string) (err error) {
	return vm.SetPropertyElementName(newName)
}

func (vm *VirtualNetworkAdapter) GetIPv4Address() (ip string, err error) {
	defer func() {
		log.Printf("[WMI][VirtualNetworkAdapter] GetIPv4Address [%s]\n", ip)
	}()

	gnac, err := vm.GetGuestNetworkAdapterConfiguration()
	if err != nil {
		return
	}
	defer gnac.Close()
	v4s, _, err := gnac.GetIPAddresses()
	if err != nil {
		return
	}
	if len(v4s) == 0 {
		// Empty IPAddress - Valid
		return
	}

	ip = v4s[0]
	return
}

func (vm *VirtualNetworkAdapter) GetGuestNetworkAdapterInstanceID() (id string, err error) {
	gnac, err := vm.GetGuestNetworkAdapterConfiguration()
	if err != nil {
		return
	}
	defer gnac.Close()
	id, err = gnac.GetPropertyInstanceID()
	return
}
func (vm *VirtualNetworkAdapter) GetGuestNetworkAdapterConfiguration() (*GuestNetworkAdapterConfiguration, error) {
	inst, err := vm.GetRelated("Msvm_GuestNetworkAdapterConfiguration")
	if err != nil {
		return nil, err
	}
	return NewGuestNetworkAdapterConfiguration(inst)
}

func (vm *VirtualNetworkAdapter) WaitForJobCompletion(result int32, jobType v2.ConcreteJob_JobType) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := vm.getJob(jobType)
		if err != nil {
			return err
			// Job is scheduled, but we were not able to find the job
		}
		defer vmjob.Close()

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed,
			"Unable to Change the state of the Virtual Machine Result[%d] [%d]", result, jobType)
	}
}

func (vm *VirtualNetworkAdapter) getJob(jobType v2.ConcreteJob_JobType) (*job.VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := vm.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	defer jobs.Close()
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return job.NewVirtualSystemJob(jobInstance)
}

// GetVirtualMachine gets the VM that the adapter is attached to
func (vm *VirtualNetworkAdapter) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := vm.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
