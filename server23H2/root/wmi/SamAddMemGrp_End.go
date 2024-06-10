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

// SamAddMemGrp_End struct
type SamAddMemGrp_End struct {
	*SamAddMemGrp
}

func NewSamAddMemGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamAddMemGrp_End, err error) {
	tmp, err := NewSamAddMemGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamAddMemGrp_End{
		SamAddMemGrp: tmp,
	}
	return
}

func NewSamAddMemGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamAddMemGrp_End, err error) {
	tmp, err := NewSamAddMemGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamAddMemGrp_End{
		SamAddMemGrp: tmp,
	}
	return
}
