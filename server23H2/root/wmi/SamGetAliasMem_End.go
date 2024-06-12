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

// SamGetAliasMem_End struct
type SamGetAliasMem_End struct {
	*SamGetAliasMem
}

func NewSamGetAliasMem_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetAliasMem_End, err error) {
	tmp, err := NewSamGetAliasMemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetAliasMem_End{
		SamGetAliasMem: tmp,
	}
	return
}

func NewSamGetAliasMem_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetAliasMem_End, err error) {
	tmp, err := NewSamGetAliasMemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetAliasMem_End{
		SamGetAliasMem: tmp,
	}
	return
}
