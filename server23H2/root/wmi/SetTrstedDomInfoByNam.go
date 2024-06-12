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

// SetTrstedDomInfoByNam struct
type SetTrstedDomInfoByNam struct {
	*MSLSATrace
}

func NewSetTrstedDomInfoByNamEx1(instance *cim.WmiInstance) (newInstance *SetTrstedDomInfoByNam, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam{
		MSLSATrace: tmp,
	}
	return
}

func NewSetTrstedDomInfoByNamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetTrstedDomInfoByNam, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetTrstedDomInfoByNam{
		MSLSATrace: tmp,
	}
	return
}
