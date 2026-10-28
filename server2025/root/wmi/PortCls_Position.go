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

// PortCls_Position struct
type PortCls_Position struct {
	*PortClsEvent
}

func NewPortCls_PositionEx1(instance *cim.WmiInstance) (newInstance *PortCls_Position, err error) {
	tmp, err := NewPortClsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PortCls_Position{
		PortClsEvent: tmp,
	}
	return
}

func NewPortCls_PositionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PortCls_Position, err error) {
	tmp, err := NewPortClsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PortCls_Position{
		PortClsEvent: tmp,
	}
	return
}
