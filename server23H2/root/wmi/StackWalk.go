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

// StackWalk struct
type StackWalk struct {
	*MSNT_SystemTrace
}

func NewStackWalkEx1(instance *cim.WmiInstance) (newInstance *StackWalk, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &StackWalk{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewStackWalkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *StackWalk, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &StackWalk{
		MSNT_SystemTrace: tmp,
	}
	return
}
