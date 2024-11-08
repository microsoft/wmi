// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

var (
	serviceStoreMap map[string]*ClusterService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*ClusterService{}
}

type ClusterService struct {
	*fc.MSCluster_ClusterService
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetClusterService(whost *host.WmiHost) (mgmt *ClusterService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	mgmt, err = getService(whost)
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

func getService(whost *host.WmiHost) (mgmt *ClusterService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_Service")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := fc.NewMSCluster_ClusterServiceEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &ClusterService{vmmswmi}
	return
}
