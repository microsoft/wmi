// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package network

import (
	"github.com/microsoft/wmi/pkg/base/host"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

var (
	switchStoreMap map[string]*wmi.WmiInstance
)

func GetVirtualEthernetSwitchManagementService(whost *host.WmiHost) {
}
