// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheHashCache struct
type MSFT_NetBranchCacheHashCache struct {
	*MSFT_NetBranchCachePrimaryCache
}

func NewMSFT_NetBranchCacheHashCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheHashCache, err error) {
	tmp, err := NewMSFT_NetBranchCachePrimaryCacheEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheHashCache{
		MSFT_NetBranchCachePrimaryCache: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheHashCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheHashCache, err error) {
	tmp, err := NewMSFT_NetBranchCachePrimaryCacheEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheHashCache{
		MSFT_NetBranchCachePrimaryCache: tmp,
	}
	return
}
