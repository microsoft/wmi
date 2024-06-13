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

// MSV1_0Trace struct
type MSV1_0Trace struct {
	*EventTrace
}

func NewMSV1_0TraceEx1(instance *cim.WmiInstance) (newInstance *MSV1_0Trace, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSV1_0Trace{
		EventTrace: tmp,
	}
	return
}

func NewMSV1_0TraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSV1_0Trace, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSV1_0Trace{
		EventTrace: tmp,
	}
	return
}
