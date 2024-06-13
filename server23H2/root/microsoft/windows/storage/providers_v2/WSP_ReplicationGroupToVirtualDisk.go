// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_ReplicationGroupToVirtualDisk struct
type WSP_ReplicationGroupToVirtualDisk struct {
	*MSFT_ReplicationGroupToVirtualDisk
}

func NewWSP_ReplicationGroupToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroupToVirtualDisk, err error) {
	tmp, err := NewMSFT_ReplicationGroupToVirtualDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToVirtualDisk{
		MSFT_ReplicationGroupToVirtualDisk: tmp,
	}
	return
}

func NewWSP_ReplicationGroupToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationGroupToVirtualDisk, err error) {
	tmp, err := NewMSFT_ReplicationGroupToVirtualDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToVirtualDisk{
		MSFT_ReplicationGroupToVirtualDisk: tmp,
	}
	return
}
