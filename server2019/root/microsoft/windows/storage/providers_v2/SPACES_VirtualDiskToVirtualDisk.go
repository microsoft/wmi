// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_VirtualDiskToVirtualDisk struct
type SPACES_VirtualDiskToVirtualDisk struct {
	*MSFT_VirtualDiskToVirtualDisk
}

func NewSPACES_VirtualDiskToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_VirtualDiskToVirtualDisk, err error) {
	tmp, err := NewMSFT_VirtualDiskToVirtualDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDiskToVirtualDisk{
		MSFT_VirtualDiskToVirtualDisk: tmp,
	}
	return
}

func NewSPACES_VirtualDiskToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_VirtualDiskToVirtualDisk, err error) {
	tmp, err := NewMSFT_VirtualDiskToVirtualDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDiskToVirtualDisk{
		MSFT_VirtualDiskToVirtualDisk: tmp,
	}
	return
}
