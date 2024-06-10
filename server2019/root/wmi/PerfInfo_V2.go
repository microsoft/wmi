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

// PerfInfo_V2 struct
type PerfInfo_V2 struct {
	*MSNT_SystemTrace
}

func NewPerfInfo_V2Ex1(instance *cim.WmiInstance) (newInstance *PerfInfo_V2, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PerfInfo_V2{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewPerfInfo_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PerfInfo_V2, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PerfInfo_V2{
		MSNT_SystemTrace: tmp,
	}
	return
}
