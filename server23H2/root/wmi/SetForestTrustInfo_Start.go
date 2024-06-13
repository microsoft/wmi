// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SetForestTrustInfo_Start struct
type SetForestTrustInfo_Start struct {
	*SetForestTrustInfo
}

func NewSetForestTrustInfo_StartEx1(instance *cim.WmiInstance) (newInstance *SetForestTrustInfo_Start, err error) {
	tmp, err := NewSetForestTrustInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetForestTrustInfo_Start{
		SetForestTrustInfo: tmp,
	}
	return
}

func NewSetForestTrustInfo_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetForestTrustInfo_Start, err error) {
	tmp, err := NewSetForestTrustInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetForestTrustInfo_Start{
		SetForestTrustInfo: tmp,
	}
	return
}
