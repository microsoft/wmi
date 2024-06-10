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

// SamSetMemAttrsOfGrp_End struct
type SamSetMemAttrsOfGrp_End struct {
	*SamSetMemAttrsOfGrp
}

func NewSamSetMemAttrsOfGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetMemAttrsOfGrp_End, err error) {
	tmp, err := NewSamSetMemAttrsOfGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetMemAttrsOfGrp_End{
		SamSetMemAttrsOfGrp: tmp,
	}
	return
}

func NewSamSetMemAttrsOfGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetMemAttrsOfGrp_End, err error) {
	tmp, err := NewSamSetMemAttrsOfGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetMemAttrsOfGrp_End{
		SamSetMemAttrsOfGrp: tmp,
	}
	return
}
