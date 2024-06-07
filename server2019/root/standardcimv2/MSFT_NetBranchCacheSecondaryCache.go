// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheSecondaryCache struct
type MSFT_NetBranchCacheSecondaryCache struct { 
	*MSFT_NetBranchCacheCache
}

	func NewMSFT_NetBranchCacheSecondaryCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheSecondaryCache, err error) {tmp, err := NewMSFT_NetBranchCacheCacheEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetBranchCacheSecondaryCache {
	MSFT_NetBranchCacheCache: tmp,
	}
	return
	}
	

	func NewMSFT_NetBranchCacheSecondaryCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetBranchCacheSecondaryCache, err error) {tmp, err := NewMSFT_NetBranchCacheCacheEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetBranchCacheSecondaryCache {
	MSFT_NetBranchCacheCache: tmp,
	}
	return
	}
	

