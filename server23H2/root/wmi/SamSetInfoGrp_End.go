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

// SamSetInfoGrp_End struct
type SamSetInfoGrp_End struct {
	*SamSetInfoGrp
}

func NewSamSetInfoGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetInfoGrp_End, err error) {
	tmp, err := NewSamSetInfoGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoGrp_End{
		SamSetInfoGrp: tmp,
	}
	return
}

func NewSamSetInfoGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetInfoGrp_End, err error) {
	tmp, err := NewSamSetInfoGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoGrp_End{
		SamSetInfoGrp: tmp,
	}
	return
}
