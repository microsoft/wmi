// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternode

import (
	"os"

	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
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

// GetClusterNode gets an existing virtual machine
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
	} ()

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


// IsUp get the cluster health status
func (c *ClusterNode) IsUp()  (status bool) {
	state, err := c.GetPropertyState()
	if err != nil {
		return
	}

	return (int32(state) == fcconstant.CLUSTER_NODE_STATE_UP)
}

// IsPaused get the cluster health status
func (c *ClusterNode) IsPaused()  (status bool) {
	state, err := c.GetPropertyState()
	if err != nil {
		return
	}

	return (int32(state) == fcconstant.CLUSTER_NODE_STATE_PAUSED)
}