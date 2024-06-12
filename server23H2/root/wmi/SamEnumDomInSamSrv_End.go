// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamEnumDomInSamSrv_End struct
type SamEnumDomInSamSrv_End struct {
	*SamEnumDomInSamSrv
}

func NewSamEnumDomInSamSrv_EndEx1(instance *cim.WmiInstance) (newInstance *SamEnumDomInSamSrv_End, err error) {
	tmp, err := NewSamEnumDomInSamSrvEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamEnumDomInSamSrv_End{
		SamEnumDomInSamSrv: tmp,
	}
	return
}

func NewSamEnumDomInSamSrv_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamEnumDomInSamSrv_End, err error) {
	tmp, err := NewSamEnumDomInSamSrvEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamEnumDomInSamSrv_End{
		SamEnumDomInSamSrv: tmp,
	}
	return
}
