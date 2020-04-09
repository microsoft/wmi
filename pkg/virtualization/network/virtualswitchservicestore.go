// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package network

import (
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	switchStoreMap map[string]*v2.Msvm_VirtualEthernetSwitchManagementService
	mux            sync.Mutex
)

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetVirtualEthernetSwitchManagementService(whost *host.WmiHost) (mgmt *v2.Msvm_VirtualEthernetSwitchManagementService, err error) {
	mux.Lock()
	defer mux.Unlock()

	if val, ok := switchStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualEthernetSwitchManagementService", "Caption", "Hyper-V Virtual System Management Service")
	// TODO: Regenerate wrappers that would take WmiHost directly
	mgmt, err = v2.NewMsvm_VirtualEthernetSwitchManagementServiceEx6(whost.HostName, "root/virtualization/v2", creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	switchStoreMap[whost.HostName] = mgmt
	return
}
