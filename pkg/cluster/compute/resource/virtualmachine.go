// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
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
	wmigpset, err := fc.NewMSCluster_ResourceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crgSet = &Resource{wmigpset}
	return
}
