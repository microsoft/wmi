// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternetwork

import (

	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
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
	} ()

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


