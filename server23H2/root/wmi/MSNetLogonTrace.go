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

// MSNetLogonTrace struct
type MSNetLogonTrace struct {
	*EventTrace
}

func NewMSNetLogonTraceEx1(instance *cim.WmiInstance) (newInstance *MSNetLogonTrace, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNetLogonTrace{
		EventTrace: tmp,
	}
	return
}

func NewMSNetLogonTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNetLogonTrace, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNetLogonTrace{
		EventTrace: tmp,
	}
	return
}
