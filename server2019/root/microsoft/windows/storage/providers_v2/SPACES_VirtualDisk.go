// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_VirtualDisk struct
type SPACES_VirtualDisk struct {
	*MSFT_VirtualDisk
}

func NewSPACES_VirtualDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_VirtualDisk, err error) {
	tmp, err := NewMSFT_VirtualDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDisk{
		MSFT_VirtualDisk: tmp,
	}
	return
}

func NewSPACES_VirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_VirtualDisk, err error) {
	tmp, err := NewMSFT_VirtualDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDisk{
		MSFT_VirtualDisk: tmp,
	}
	return
}
