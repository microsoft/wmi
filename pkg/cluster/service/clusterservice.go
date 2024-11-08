// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"sync"
	"os"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

var (
	serviceStoreMap map[string]*ClusterService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*ClusterService{}
}

type ClusterService struct {
	*fc.MSCluster_Service
}

// NewClusterService
func NewClusterService(instance *wmi.WmiInstance) (*ClusterService, error) {
	wmics, err := fc.NewMSCluster_ServiceEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ClusterService{wmics}, nil
}


// GetClusterService gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterServices(whost *host.WmiHost) (cservicecollection ClusterServiceCollection, err error) {
	query := query.NewWmiQuery("MSCluster_Service")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	cservicecollection, err = NewClusterServiceCollection(instances)
	return
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetLocalClusterService(whost *host.WmiHost) (mgmt *ClusterService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	mgmt, err = getLocalService(whost)
	if err != nil {
		return
	}

	mux.Lock()
	defer mux.Unlock()
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt.Close()
		mgmt = val
		return
	}

	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getLocalService(whost *host.WmiHost) (mgmt *ClusterService, err error) {
	creds := whost.GetCredential()
	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	query := query.NewWmiQuery("MSCluster_Service", "SystemName", hostname)
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := fc.NewMSCluster_ServiceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &ClusterService{vmmswmi}
	return
}
