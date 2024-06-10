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

// SamQueryInfoUser_End struct
type SamQueryInfoUser_End struct {
	*SamQueryInfoUser
}

func NewSamQueryInfoUser_EndEx1(instance *cim.WmiInstance) (newInstance *SamQueryInfoUser_End, err error) {
	tmp, err := NewSamQueryInfoUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoUser_End{
		SamQueryInfoUser: tmp,
	}
	return
}

func NewSamQueryInfoUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQueryInfoUser_End, err error) {
	tmp, err := NewSamQueryInfoUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoUser_End{
		SamQueryInfoUser: tmp,
	}
	return
}
