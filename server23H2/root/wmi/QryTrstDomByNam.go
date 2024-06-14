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

// QryTrstDomByNam struct
type QryTrstDomByNam struct {
	*MSLSATrace
}

func NewQryTrstDomByNamEx1(instance *cim.WmiInstance) (newInstance *QryTrstDomByNam, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QryTrstDomByNam{
		MSLSATrace: tmp,
	}
	return
}

func NewQryTrstDomByNamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QryTrstDomByNam, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QryTrstDomByNam{
		MSLSATrace: tmp,
	}
	return
}
