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

// PerfInfo struct
type PerfInfo struct {
	*MSNT_SystemTrace
}

func NewPerfInfoEx1(instance *cim.WmiInstance) (newInstance *PerfInfo, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PerfInfo{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewPerfInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PerfInfo, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PerfInfo{
		MSNT_SystemTrace: tmp,
	}
	return
}
