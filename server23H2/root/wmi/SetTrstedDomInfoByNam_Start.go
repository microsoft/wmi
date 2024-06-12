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

// SetTrstedDomInfoByNam_Start struct
type SetTrstedDomInfoByNam_Start struct {
	*SetTrstedDomInfoByNam
}

func NewSetTrstedDomInfoByNam_StartEx1(instance *cim.WmiInstance) (newInstance *SetTrstedDomInfoByNam_Start, err error) {
	tmp, err := NewSetTrstedDomInfoByNamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam_Start{
		SetTrstedDomInfoByNam: tmp,
	}
	return
}

func NewSetTrstedDomInfoByNam_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetTrstedDomInfoByNam_Start, err error) {
	tmp, err := NewSetTrstedDomInfoByNamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam_Start{
		SetTrstedDomInfoByNam: tmp,
	}
	return
}
