// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// QueryForestTrustInfo_Start struct
type QueryForestTrustInfo_Start struct {
	*QueryForestTrustInfo
}

func NewQueryForestTrustInfo_StartEx1(instance *cim.WmiInstance) (newInstance *QueryForestTrustInfo_Start, err error) {
	tmp, err := NewQueryForestTrustInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryForestTrustInfo_Start{
		QueryForestTrustInfo: tmp,
	}
	return
}

func NewQueryForestTrustInfo_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryForestTrustInfo_Start, err error) {
	tmp, err := NewQueryForestTrustInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryForestTrustInfo_Start{
		QueryForestTrustInfo: tmp,
	}
	return
}
