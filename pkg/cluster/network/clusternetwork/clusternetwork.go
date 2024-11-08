// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternetwork

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"

	"github.com/microsoft/wmi/pkg/errors"
	"reflect"
)

type ClusterNetwork struct {
	*fc.MSCluster_Network
}

// NewClusterNetwork
func NewClusterNetwork(instance *wmi.WmiInstance) (*ClusterNetwork, error) {
	wminetwork, err := fc.NewMSCluster_NetworkEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ClusterNetwork{wminetwork}, nil
}

// GetClusterNetwork gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterNetworks(whost *host.WmiHost) (cnodecollection ClusterNetworkCollection, err error) {
	query := query.NewWmiQuery("MSCluster_Network")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	cnodecollection, err = NewClusterNetworkCollection(instances)
	return
}

// GetClusterNetwork gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterNetwork(whost *host.WmiHost, nodeName string) (cnode *ClusterNetwork, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Network", "Name", nodeName)
	wminetwork, err := fc.NewMSCluster_NetworkEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cnode = &ClusterNetwork{wminetwork}
	return
}

// State gets the value of FaultState for the instance
func (c *ClusterNetwork) State() (value int32, err error) {
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

// IsPartitioned get the cluster network health status
func (c *ClusterNetwork) IsPartitioned() (status bool, err error) {
	state, err := c.State()
	if err != nil {
		return
	}

	status = (state == fcconstant.CLUSTER_NETWORK_STATE_PARTITIONED)
	return
}
