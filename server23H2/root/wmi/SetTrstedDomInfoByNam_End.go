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

// SetTrstedDomInfoByNam_End struct
type SetTrstedDomInfoByNam_End struct {
	*SetTrstedDomInfoByNam
}

func NewSetTrstedDomInfoByNam_EndEx1(instance *cim.WmiInstance) (newInstance *SetTrstedDomInfoByNam_End, err error) {
	tmp, err := NewSetTrstedDomInfoByNamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam_End{
		SetTrstedDomInfoByNam: tmp,
	}
	return
}

func NewSetTrstedDomInfoByNam_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetTrstedDomInfoByNam_End, err error) {
	tmp, err := NewSetTrstedDomInfoByNamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam_End{
		SetTrstedDomInfoByNam: tmp,
	}
	return
}
