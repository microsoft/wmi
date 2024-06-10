// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// DebuggerEnabled struct
type DebuggerEnabled struct {
	*PerfInfo_V2
}

func NewDebuggerEnabledEx1(instance *cim.WmiInstance) (newInstance *DebuggerEnabled, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DebuggerEnabled{
		PerfInfo_V2: tmp,
	}
	return
}

func NewDebuggerEnabledEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DebuggerEnabled, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DebuggerEnabled{
		PerfInfo_V2: tmp,
	}
	return
}
