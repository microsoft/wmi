// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

import (
	"reflect"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	//fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/go-ole/go-ole"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type Resource struct {
	*fc.MSCluster_Resource
}

// NewResource
func NewResource(instance *wmi.WmiInstance) (*Resource, error) {
	wmigpset, err := fc.NewMSCluster_ResourceEx1(instance)
	if err != nil {
		return nil, err
	}
	return &Resource{wmigpset}, nil
}

// GetResource gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetResources(whost *host.WmiHost) (crSetcollection ResourceCollection, err error) {
	query := query.NewWmiQuery("MSCluster_Resource")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	crSetcollection, err = NewResourceCollection(instances)
	return
}

// GetResource gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetResource(whost *host.WmiHost, grpSetName string) (crSet *Resource, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Resource", "Name", grpSetName)
	wmigpset, err := fc.NewMSCluster_ResourceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	crSet = &Resource{wmigpset}
	return
}

// State gets the value of for the instance
func (c *Resource) State() (value int32, err error) {
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

func (c *Resource) GetPossibleOwnersEx1() (nodeNames []string, err error) {
	nodeListVariant := ole.NewVariant(ole.VT_ARRAY|ole.VT_BSTR, 0)
	_, err = c.InvokeMethod("GetPossibleOwners", &nodeListVariant)
	if err != nil {
		return
	}

	nodeNameValues, err := wmi.GetVariantValues(&nodeListVariant)
	if err != nil {
		return
	}

	for _, ownerNode := range nodeNameValues {
		nodeNames = append(nodeNames, ownerNode.(string))
	}

	return
}
