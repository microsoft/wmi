// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

import (
	"fmt"
	"math"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"

	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/cluster/compute/internal"
)

type FailoverClusterVirtualMachine struct {
	*ResourceGroup
	VmId 		string
}

var (
	groupTypeString = fmt.Sprintf("%d", CLUSTER_RESOURCE_GROUP_TYPE_VIRTUAL_MACHINE)
)

func NewFailoverClusterVirtualMachine(instance *wmi.WmiInstance) (newInstance *FailoverClusterVirtualMachine, err error) {
	tmp, err := NewResourceGroup(instance)

	if err != nil {
		return
	}
	newInstance = &FailoverClusterVirtualMachine{
		ResourceGroup: tmp,
	}
	return
}

// GetVirtualMachineResourceGroups gets list of resource groups that are of type VirtualMachine
// Make sure to call Close once done using this instance
func GetVirtualMachineResourceGroups(whost *host.WmiHost) (crgSetcollection ResourceGroupCollection, err error) {
	query := query.NewWmiQuery("MSCluster_ResourceGroup", "GroupType", groupTypeString)
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	crgSetcollection, err = NewResourceGroupCollection(instances)
	return
}

// GetVirtualMachineResourceGroup gets an existing virtual machine resource group
// Make sure to call Close once done using this instance
func GetVirtualMachineResourceGroup(whost *host.WmiHost, virtualMachineName string) (crgSet *ResourceGroup, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_ResourceGroup", "Name", virtualMachineName, "GroupType", groupTypeString)
	wmigpset, err := fc.NewMSCluster_ResourceGroupEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crgSet = &ResourceGroup{wmigpset}
	return
}

// GetVirtualMachineResourceGroupByVmID gets an existing virtual machine resource group
// Make sure to call Close once done using this instance
func GetVirtualMachineResourceGroupByVmID(whost *host.WmiHost, vmID string) (crgSet *ResourceGroup, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_ResourceGroup", "PrivateProperties.VmID", vmID, "GroupType", groupTypeString)
	wmigpset, err := fc.NewMSCluster_ResourceGroupEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crgSet = &ResourceGroup{wmigpset}
	return
}


// SetVirtualMachineAutoFailback returns if the resource group is in partial or pending state
func (c *ResourceGroup) SetVirtualMachineAutoFailback(enable bool) (err error) {
	value := 0
	if enable {
		value = 1
	}

	err = c.SetPropertyAutoFailbackType(uint32(value))
	if err != nil {
		return err
	}
	err = c.Commit()

	return
}

// DisableVirtualMachineDefaultOwner returns if the resource group is in partial or pending state
// DefaultOwner is set to node on which cluster group is created.
// With Autofail back enabled, currently VM will be moved to default owner node when default owner comes online.
// Disabling default owner will ensure that VM fails back based on preferred owners only.
// Disabled value is equal to -1 / MaxUint32
func (c *ResourceGroup) DisableVirtualMachineDefaultOwner() (err error) {
	err = c.SetPropertyDefaultOwner(math.MaxUint32)
	if err != nil {
		return err
	}
	err = c.Commit()

	return
}

// GetVirtualMachineResourceGroupViaAssociators gets list of resource groups that are of type VirtualMachine
// Make sure to call Close once done using this instance
func GetVirtualMachineResourceGroupViaAssociators(whost *host.WmiHost, virtualMachineName string) (fcVm *FailoverClusterVirtualMachine, err error) {
	query := fmt.Sprintf("ASSOCIATORS OF {MSCluster_ResourceGroup.Name='%s'} WHERE AssocClass = MSCluster_ResourceGroupToResource", virtualMachineName)
	instances, err := instance.GetWmiInstancesFromHostRawQuery(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	if len(instances) == 0 {
		err = errors.Wrapf(errors.NotFound, "No resource group found for virtual machine %s", virtualMachineName)
		return
	}

	fcVm, err = NewFailoverClusterVirtualMachine(instances[0])
	if err != nil {
		return
	}

	// Get the VMId
	xmlString, err := fcVm.EmbeddedXMLInstance()
	if err != nil {
		return nil, err
	}
	vmID, err := internal.GetVMID(xmlString)
	if err != nil {
		return nil, errors.Wrapf(errors.NotFound, "Missing VMID for VM '%s' in the cluster", virtualMachineName)
	}
	fcVm.VmId = vmID;

	return
}