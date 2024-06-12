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

// SamGetDomPwdInfo_End struct
type SamGetDomPwdInfo_End struct {
	*SamGetDomPwdInfo
}

func NewSamGetDomPwdInfo_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetDomPwdInfo_End, err error) {
	tmp, err := NewSamGetDomPwdInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetDomPwdInfo_End{
		SamGetDomPwdInfo: tmp,
	}
	return
}

func NewSamGetDomPwdInfo_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetDomPwdInfo_End, err error) {
	tmp, err := NewSamGetDomPwdInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetDomPwdInfo_End{
		SamGetDomPwdInfo: tmp,
	}
	return
}
