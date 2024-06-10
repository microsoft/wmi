// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheDataCacheExtension struct
type MSFT_NetBranchCacheDataCacheExtension struct {
	*MSFT_NetBranchCacheSecondaryCache
}

func NewMSFT_NetBranchCacheDataCacheExtensionEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheDataCacheExtension, err error) {
	tmp, err := NewMSFT_NetBranchCacheSecondaryCacheEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheDataCacheExtension{
		MSFT_NetBranchCacheSecondaryCache: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheDataCacheExtensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheDataCacheExtension, err error) {
	tmp, err := NewMSFT_NetBranchCacheSecondaryCacheEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheDataCacheExtension{
		MSFT_NetBranchCacheSecondaryCache: tmp,
	}
	return
}
