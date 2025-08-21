// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

import (
	"log"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

// GetVirtualMachineResources gets list of resource groups that are of type VirtualMachine
// Make sure to call Close once done using this instance
func GetVirtualMachineResources(whost *host.WmiHost) (crSetcollection ResourceCollection, err error) {
	query := query.NewWmiQuery("MSCluster_Resource", "Type", CLUSTER_RESOURCE_TYPE_VIRTUAL_MACHINE)
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	crSetcollection, err = NewResourceCollection(instances)
	return
}

// GetVirtualMachineResource gets an existing virtual machine resource group
// Make sure to call Close once done using this instance
func GetVirtualMachineResource(whost *host.WmiHost, virtualMachineName string) (crgSet *Resource, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Resource", "OwnerGroup", virtualMachineName, "Type", CLUSTER_RESOURCE_TYPE_VIRTUAL_MACHINE)
	wmigpset, err := fc.NewMSCluster_ResourceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crgSet = &Resource{wmigpset}
	return
}

// GetVirtualMachineResourceByVmID gets an existing virtual machine resource group
// Make sure to call Close once done using this instance
func GetVirtualMachineResourceByVmID(whost *host.WmiHost, vmID string) (crgSet *Resource, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Resource", "PrivateProperties.VmID", vmID, "Type", CLUSTER_RESOURCE_TYPE_VIRTUAL_MACHINE)

	for attempt := 0; attempt < constant.WmiMethodMaxRetries; attempt++ {
		wmigpset, err := fc.NewMSCluster_ResourceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
		if err != nil {
			// Retry if this is a "Not Found" error and we have attempts left
			if errors.IsNotFound(err) && attempt < constant.WmiMethodMaxRetries-1 {
				backoffDuration := time.Duration(attempt+1) * constant.WmiMethodRetryDelay
				// Add logging for consistency with other functions
				log.Printf("[WMI] GetVirtualMachineResourceByVmID failed with NotFound error. Retrying (%d/%d) after %v...", attempt+1, constant.WmiMethodMaxRetries, backoffDuration)
				time.Sleep(backoffDuration)
				continue
			}
			// For non-retryable errors or last attempt, return immediately
			return nil, err
		}
		crgSet = &Resource{wmigpset}
		return crgSet, nil
	}

	// Return failure if failover cluster VM is not found after all retries
	return nil, errors.Wrapf(errors.Failed, "failed to get FC VM resource by VmID %s after %d attempts", vmID, constant.WmiMethodMaxRetries)
}
