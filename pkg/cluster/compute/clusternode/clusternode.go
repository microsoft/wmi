// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternode

import (
	"os"

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

type ClusterNode struct {
	*fc.MSCluster_Node
}

// NewClusterNode
func NewClusterNode(instance *wmi.WmiInstance) (*ClusterNode, error) {
	wmivm, err := fc.NewMSCluster_NodeEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ClusterNode{wmivm}, nil
}

// GetClusterNodes gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterNodes(whost *host.WmiHost) (cnodecollection ClusterNodeCollection, err error) {
	query := query.NewWmiQuery("MSCluster_Node")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	cnodecollection, err = NewClusterNodeCollection(instances)
	return
}

// GetClusterNode gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterNode(whost *host.WmiHost, nodeName string) (cnode *ClusterNode, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Node", "Name", nodeName)
	wmivm, err := fc.NewMSCluster_NodeEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cnode = &ClusterNode{wmivm}
	return
}

// GetLocalClusterNode gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetLocalClusterNode(whost *host.WmiHost) (cnode *ClusterNode, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	return GetClusterNode(whost, hostname)
}

// State gets the value of FaultState for the instance
func (c *ClusterNode) State() (value int32, err error) {
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

// IsUp get the cluster health status
func (c *ClusterNode) IsUp() (status bool, err error) {
	state, err := c.State()
	if err != nil {
		return
	}

	status = (state == fcconstant.CLUSTER_NODE_STATE_UP)
	return
}

// IsPaused get the cluster health status
func (c *ClusterNode) IsPaused() (status bool, err error) {
	state, err := c.State()
	if err != nil {
		return
	}

	status = (state == fcconstant.CLUSTER_NODE_STATE_PAUSED)
	return
}
