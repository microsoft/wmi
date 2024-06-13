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

// NlServerAuth struct
type NlServerAuth struct {
	*MSNetLogonTrace
}

func NewNlServerAuthEx1(instance *cim.WmiInstance) (newInstance *NlServerAuth, err error) {
	tmp, err := NewMSNetLogonTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NlServerAuth{
		MSNetLogonTrace: tmp,
	}
	return
}

func NewNlServerAuthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NlServerAuth, err error) {
	tmp, err := NewMSNetLogonTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NlServerAuth{
		MSNetLogonTrace: tmp,
	}
	return
}
