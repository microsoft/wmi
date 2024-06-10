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

// SamQueryDisplayInfo_End struct
type SamQueryDisplayInfo_End struct {
	*SamQueryDisplayInfo
}

func NewSamQueryDisplayInfo_EndEx1(instance *cim.WmiInstance) (newInstance *SamQueryDisplayInfo_End, err error) {
	tmp, err := NewSamQueryDisplayInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQueryDisplayInfo_End{
		SamQueryDisplayInfo: tmp,
	}
	return
}

func NewSamQueryDisplayInfo_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQueryDisplayInfo_End, err error) {
	tmp, err := NewSamQueryDisplayInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQueryDisplayInfo_End{
		SamQueryDisplayInfo: tmp,
	}
	return
}
