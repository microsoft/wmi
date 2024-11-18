// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

import (
	"fmt"
	"reflect"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
	"math"

	"github.com/microsoft/wmi/pkg/cluster/compute/internal"
	"github.com/microsoft/wmi/pkg/errors"
)

var (
	groupTypeString = fmt.Sprintf("%d", CLUSTER_RESOURCE_GROUP_TYPE_VIRTUAL_MACHINE)
)

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

// SetVirtualMachineAutoFailback sets the virtual machine auto failback policy
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

// DisableVirtualMachineDefaultOwner
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
func GetVirtualMachineID(whost *host.WmiHost, virtualMachineName string) (vmID string, err error) {
	query := fmt.Sprintf("ASSOCIATORS OF {MSCluster_ResourceGroup.Name='%s'} WHERE AssocClass = MSCluster_ResourceGroupToResource", virtualMachineName)
	instances, err := instance.GetWmiInstancesFromHostRawQuery(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}
	defer instances.Close()

	if len(instances) == 0 {
		err = errors.Wrapf(errors.NotFound, "No resource group found for virtual machine %s", virtualMachineName)
		return
	}

	xmlString, err := instances[0].EmbeddedXMLInstance()
	if err != nil {
		return
	}
	vmID, err = internal.GetVMID(xmlString)
	if err != nil {
		return
	}

	if vmID == "" {
		err = errors.Wrapf(errors.NotFound, "Missing VMID for VM '%s' in the cluster", virtualMachineName)
	}
	return
}

// SetVirtualMachinePriority sets the priority of the virtual machine
func (c *ResourceGroup) SetVirtualMachinePriority(value int32) (err error) {
	err = c.SetProperty("Priority", value)
	if err != nil {
		return err
	}
	err = c.Commit()

	return
}

// GetVirtualMachinePriority sets the priority of the virtual machine
func (c *ResourceGroup) GetVirtualMachinePriority() (value int32, err error) {
	// The documentation says uint32, but it is not working
	retValue, err := c.GetProperty("Priority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	value, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	return
}
