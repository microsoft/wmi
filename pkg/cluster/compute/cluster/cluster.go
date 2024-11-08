// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cluster

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type Cluster struct {
	*fc.MSCluster_Cluster
}

// NewCluster
func NewCluster(instance *wmi.WmiInstance) (*Cluster, error) {
	wmicluster, err := fc.NewMSCluster_ClusterEx1(instance)
	if err != nil {
		return nil, err
	}
	return &Cluster{wmicluster}, nil
}

// GetCluster gets an existing cluster
// Make sure to call Close once done using this instance
func GetCluster(whost *host.WmiHost) (ccluster *Cluster, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Cluster")
	wmicluster, err := fc.NewMSCluster_ClusterEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	return &Cluster{wmicluster}, nil
}

// IsNodeClusterStateRunning get the cluster health status
func (c *Cluster) IsNodeClusterStateRunning() (status bool) {
	state, err := c.InvokeMethodWithReturn("GetNodeClusterState")
	if err != nil {
		return false
	}

	return (state == fcconstant.ClusterStateRunning)
}
