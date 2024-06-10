// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Storage.Providers_v2
//
// ////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_DiskToPartition struct
type WSP_DiskToPartition struct {
	*MSFT_DiskToPartition
}

func NewWSP_DiskToPartitionEx1(instance *cim.WmiInstance) (newInstance *WSP_DiskToPartition, err error) {
	tmp, err := NewMSFT_DiskToPartitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_DiskToPartition{
		MSFT_DiskToPartition: tmp,
	}
	return
}

func NewWSP_DiskToPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_DiskToPartition, err error) {
	tmp, err := NewMSFT_DiskToPartitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_DiskToPartition{
		MSFT_DiskToPartition: tmp,
	}
	return
}
