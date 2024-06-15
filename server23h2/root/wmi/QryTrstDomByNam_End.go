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

// QryTrstDomByNam_End struct
type QryTrstDomByNam_End struct {
	*QryTrstDomByNam
}

func NewQryTrstDomByNam_EndEx1(instance *cim.WmiInstance) (newInstance *QryTrstDomByNam_End, err error) {
	tmp, err := NewQryTrstDomByNamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QryTrstDomByNam_End{
		QryTrstDomByNam: tmp,
	}
	return
}

func NewQryTrstDomByNam_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QryTrstDomByNam_End, err error) {
	tmp, err := NewQryTrstDomByNamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QryTrstDomByNam_End{
		QryTrstDomByNam: tmp,
	}
	return
}
