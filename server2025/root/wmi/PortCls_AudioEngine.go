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

// PortCls_AudioEngine struct
type PortCls_AudioEngine struct {
	*PortClsEvent
}

func NewPortCls_AudioEngineEx1(instance *cim.WmiInstance) (newInstance *PortCls_AudioEngine, err error) {
	tmp, err := NewPortClsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PortCls_AudioEngine{
		PortClsEvent: tmp,
	}
	return
}

func NewPortCls_AudioEngineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PortCls_AudioEngine, err error) {
	tmp, err := NewPortClsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PortCls_AudioEngine{
		PortClsEvent: tmp,
	}
	return
}
