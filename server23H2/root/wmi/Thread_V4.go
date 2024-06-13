// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Thread_V4 struct
type Thread_V4 struct {
	*MSNT_SystemTrace
}

func NewThread_V4Ex1(instance *cim.WmiInstance) (newInstance *Thread_V4, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Thread_V4{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewThread_V4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Thread_V4, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Thread_V4{
		MSNT_SystemTrace: tmp,
	}
	return
}
