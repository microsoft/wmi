// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*VirtualEthernetSwitchManagementService
	mux             sync.Mutex
)

type VirtualEthernetSwitchManagementService struct {
	*v2.Msvm_VirtualEthernetSwitchManagementService
	host *host.WmiHost
}

func init() {
	serviceStoreMap = map[string]*VirtualEthernetSwitchManagementService{}
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetVirtualEthernetSwitchManagementService(whost *host.WmiHost) (mgmt *VirtualEthernetSwitchManagementService, err error) {
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

func getService(whost *host.WmiHost) (mgmt *VirtualEthernetSwitchManagementService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualEthernetSwitchManagementService")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vswitchwmi, err := v2.NewMsvm_VirtualEthernetSwitchManagementServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &VirtualEthernetSwitchManagementService{vswitchwmi, whost}
	return
}

// GetVirtualSwitches would get all virtual machines
func (vsms *VirtualEthernetSwitchManagementService) GetVirtualSwitches() (*virtualswitch.VirtualSwitchCollection, error) {
	query := query.NewWmiQuery("Msvm_VirtualEthernetSwitch")
	instances, err := instance.GetWmiInstancesFromHost(vsms.host, string(constant.Virtualization), query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			instances.Close()
		}
	}()
	vmc := virtualswitch.VirtualSwitchCollection{}
	for _, ins := range instances {
		vm, err := virtualswitch.NewVirtualSwitch(ins)
		if err != nil {
			return nil, err
		}

		vmc = append(vmc, vm)
	}
	return &vmc, nil
}

// FindVirtualSwitchByName
func (vsms *VirtualEthernetSwitchManagementService) FindVirtualSwitchByName(vswitchName string) (*virtualswitch.VirtualSwitch, error) {
	vswitchs, err := vsms.GetVirtualSwitches()
	if err != nil {
		return nil, err
	}
	defer vswitchs.Close()

	for _, entity := range *vswitchs {
		entityName, err := entity.GetPropertyElementName()
		if err != nil {
			return nil, err
		}
		if entityName != vswitchName {
			continue
		}

		clonedEntity, err := entity.Clone()
		if err != nil {
			return nil, err
		}
		return virtualswitch.NewVirtualSwitch(clonedEntity)
	}

	return nil, errors.Wrapf(errors.NotFound, "Unable to find a virtual system with name [%s]", vswitchName)
}

func (vsms *VirtualEthernetSwitchManagementService) getJob(jobType v2.ConcreteJob_JobType) (*job.VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := vsms.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return job.NewVirtualSystemJob(jobInstance)
}
