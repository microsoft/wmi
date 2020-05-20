// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourcepool"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*ResourcePoolConfigurationService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*ResourcePoolConfigurationService{}

	whost := host.NewWmiLocalHost()
	serviceStoreMap[whost.HostName], _ = getService(whost)
}

type ResourcePoolConfigurationService struct {
	*v2.Msvm_ResourcePoolConfigurationService
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetResourcePoolConfigurationService(whost *host.WmiHost) (mgmt *ResourcePoolConfigurationService, err error) {
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
	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getService(whost *host.WmiHost) (mgmt *ResourcePoolConfigurationService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ResourcePoolConfigurationService", "Caption", "Hyper-V Resource Pool Configuration Service")
	// TODO: Regenerate wrappers that would take WmiHost directly
	rpcswmi, err := v2.NewMsvm_ResourcePoolConfigurationServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &ResourcePoolConfigurationService{rpcswmi}
	return
}

// ModifyPoolResources
func (rpcs *ResourcePoolConfigurationService) ModifyPoolResourcesEx(child resourcepool.ResourcePool,
	parentPools resourcepool.ResourcePoolCollection,
	rasds resourceallocation.ResourceAllocationSettingDataCollection) error {
	embeddedInstance, err := rasds.EmbeddedXMLInstances()
	if err != nil {
		return err
	}
	result, err := rpcs.InvokeMethodWithReturn("ModifyPoolResources", child, parentPools, embeddedInstance)
	if err != nil {
		return err
	}
	return rpcs.WaitForJobCompletion(result, v2.ConcreteJob_JobType_Modify_Virtual_Machine_Resources)
}

func (rpcs *ResourcePoolConfigurationService) WaitForJobCompletion(result int32, jobType v2.ConcreteJob_JobType) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := rpcs.getJob(jobType)
		if err != nil {
			// Job is scheduled, but we were not able to find the job
			return err
		}
		defer vmjob.Close()

		return vmjob.WaitForAction(wmi.Wait, 100, 10)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Resource Pool Result[%d] JobType[%d]", result, jobType)
	}
}

func (rpcs *ResourcePoolConfigurationService) getJob(jobType v2.ConcreteJob_JobType) (*job.VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := rpcs.GetAllRelatedWithQuery(query)
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
