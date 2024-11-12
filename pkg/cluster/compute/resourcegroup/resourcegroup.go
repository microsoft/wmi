// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

import (
	"reflect"

	"github.com/go-ole/go-ole"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type ResourceGroup struct {
	*fc.MSCluster_ResourceGroup
}

// NewResourceGroup
func NewResourceGroup(instance *wmi.WmiInstance) (*ResourceGroup, error) {
	wmigpset, err := fc.NewMSCluster_ResourceGroupEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ResourceGroup{wmigpset}, nil
}

// GetResourceGroup gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetResourceGroups(whost *host.WmiHost) (crgSetcollection ResourceGroupCollection, err error) {
	query := query.NewWmiQuery("MSCluster_ResourceGroup")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	crgSetcollection, err = NewResourceGroupCollection(instances)
	return
}

// GetResourceGroup gets an existing resource group
// Make sure to call Close once done using this instance
func GetResourceGroup(whost *host.WmiHost, grpSetName string) (crgSet *ResourceGroup, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_ResourceGroup", "Name", grpSetName)
	wmigpset, err := fc.NewMSCluster_ResourceGroupEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crgSet = &ResourceGroup{wmigpset}
	return
}

// State gets the value of for the instance
func (c *ResourceGroup) State() (value int32, err error) {
	retValue, err := c.GetProperty("State")
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

// IsPendingState returns if the resource group is in partial or pending state
func (c *ResourceGroup) IsPendingState() (pendingState bool, err error) {
	retState, err := c.State()
	if err != nil {
		return
	}

	pendingState = (retState >= CLUSTER_RESOURCE_GROUP_STATE_PARTIAL_ONLINE)
	return
}

// GetPreferredOwnersEx1 gets the possible owners of the resource group
func (c *ResourceGroup) GetPreferredOwnersEx1() (preferredOwners []string, err error) {
	nodeListVariant := ole.NewVariant(ole.VT_ARRAY|ole.VT_BSTR, 0)
	_, err = c.InvokeMethod("GetPreferredOwners", &nodeListVariant)
	if err != nil {
		return
	}

	nodeNameValues, err := wmi.GetVariantValues(&nodeListVariant)
	if err != nil {
		return
	}

	for _, ownerNode := range nodeNameValues {
		preferredOwners = append(preferredOwners, ownerNode.(string))
	}

	return
}
