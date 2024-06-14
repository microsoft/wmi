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

// SamQueryInfoGrp_End struct
type SamQueryInfoGrp_End struct {
	*SamQueryInfoGrp
}

func NewSamQueryInfoGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamQueryInfoGrp_End, err error) {
	tmp, err := NewSamQueryInfoGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoGrp_End{
		SamQueryInfoGrp: tmp,
	}
	return
}

func NewSamQueryInfoGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQueryInfoGrp_End, err error) {
	tmp, err := NewSamQueryInfoGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoGrp_End{
		SamQueryInfoGrp: tmp,
	}
	return
}
