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

// SamGetMemInGrp_End struct
type SamGetMemInGrp_End struct {
	*SamGetMemInGrp
}

func NewSamGetMemInGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetMemInGrp_End, err error) {
	tmp, err := NewSamGetMemInGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInGrp_End{
		SamGetMemInGrp: tmp,
	}
	return
}

func NewSamGetMemInGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetMemInGrp_End, err error) {
	tmp, err := NewSamGetMemInGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInGrp_End{
		SamGetMemInGrp: tmp,
	}
	return
}
