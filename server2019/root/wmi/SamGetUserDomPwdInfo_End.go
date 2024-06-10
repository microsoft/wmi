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

// SamGetUserDomPwdInfo_End struct
type SamGetUserDomPwdInfo_End struct {
	*SamGetUserDomPwdInfo
}

func NewSamGetUserDomPwdInfo_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetUserDomPwdInfo_End, err error) {
	tmp, err := NewSamGetUserDomPwdInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetUserDomPwdInfo_End{
		SamGetUserDomPwdInfo: tmp,
	}
	return
}

func NewSamGetUserDomPwdInfo_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetUserDomPwdInfo_End, err error) {
	tmp, err := NewSamGetUserDomPwdInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetUserDomPwdInfo_End{
		SamGetUserDomPwdInfo: tmp,
	}
	return
}
