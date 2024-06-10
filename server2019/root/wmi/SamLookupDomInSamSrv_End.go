// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamLookupDomInSamSrv_End struct
type SamLookupDomInSamSrv_End struct {
	*SamLookupDomInSamSrv
}

func NewSamLookupDomInSamSrv_EndEx1(instance *cim.WmiInstance) (newInstance *SamLookupDomInSamSrv_End, err error) {
	tmp, err := NewSamLookupDomInSamSrvEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamLookupDomInSamSrv_End{
		SamLookupDomInSamSrv: tmp,
	}
	return
}

func NewSamLookupDomInSamSrv_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamLookupDomInSamSrv_End, err error) {
	tmp, err := NewSamLookupDomInSamSrvEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamLookupDomInSamSrv_End{
		SamLookupDomInSamSrv: tmp,
	}
	return
}
