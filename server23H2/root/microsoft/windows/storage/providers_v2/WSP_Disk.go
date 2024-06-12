// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_Disk struct
type WSP_Disk struct {
	*MSFT_Disk
}

func NewWSP_DiskEx1(instance *cim.WmiInstance) (newInstance *WSP_Disk, err error) {
	tmp, err := NewMSFT_DiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_Disk{
		MSFT_Disk: tmp,
	}
	return
}

func NewWSP_DiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_Disk, err error) {
	tmp, err := NewMSFT_DiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_Disk{
		MSFT_Disk: tmp,
	}
	return
}
