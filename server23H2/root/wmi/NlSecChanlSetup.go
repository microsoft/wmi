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

// NlSecChanlSetup struct
type NlSecChanlSetup struct {
	*MSNetLogonTrace
}

func NewNlSecChanlSetupEx1(instance *cim.WmiInstance) (newInstance *NlSecChanlSetup, err error) {
	tmp, err := NewMSNetLogonTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NlSecChanlSetup{
		MSNetLogonTrace: tmp,
	}
	return
}

func NewNlSecChanlSetupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NlSecChanlSetup, err error) {
	tmp, err := NewMSNetLogonTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NlSecChanlSetup{
		MSNetLogonTrace: tmp,
	}
	return
}
