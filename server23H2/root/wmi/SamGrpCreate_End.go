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

// SamGrpCreate_End struct
type SamGrpCreate_End struct {
	*SamGrpCreate
}

func NewSamGrpCreate_EndEx1(instance *cim.WmiInstance) (newInstance *SamGrpCreate_End, err error) {
	tmp, err := NewSamGrpCreateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGrpCreate_End{
		SamGrpCreate: tmp,
	}
	return
}

func NewSamGrpCreate_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGrpCreate_End, err error) {
	tmp, err := NewSamGrpCreateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGrpCreate_End{
		SamGrpCreate: tmp,
	}
	return
}
