// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamDeleteGrp_End struct
type SamDeleteGrp_End struct {
	*SamDeleteGrp
}

func NewSamDeleteGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamDeleteGrp_End, err error) {
	tmp, err := NewSamDeleteGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamDeleteGrp_End{
		SamDeleteGrp: tmp,
	}
	return
}

func NewSamDeleteGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamDeleteGrp_End, err error) {
	tmp, err := NewSamDeleteGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamDeleteGrp_End{
		SamDeleteGrp: tmp,
	}
	return
}
