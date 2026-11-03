// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PortCls_ServiceGroup struct
type PortCls_ServiceGroup struct {
	*PortClsEvent
}

func NewPortCls_ServiceGroupEx1(instance *cim.WmiInstance) (newInstance *PortCls_ServiceGroup, err error) {
	tmp, err := NewPortClsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PortCls_ServiceGroup{
		PortClsEvent: tmp,
	}
	return
}

func NewPortCls_ServiceGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PortCls_ServiceGroup, err error) {
	tmp, err := NewPortClsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PortCls_ServiceGroup{
		PortClsEvent: tmp,
	}
	return
}
