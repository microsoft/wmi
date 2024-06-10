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

// NtlmLogonUser struct
type NtlmLogonUser struct {
	*MSV1_0Trace
}

func NewNtlmLogonUserEx1(instance *cim.WmiInstance) (newInstance *NtlmLogonUser, err error) {
	tmp, err := NewMSV1_0TraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser{
		MSV1_0Trace: tmp,
	}
	return
}

func NewNtlmLogonUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmLogonUser, err error) {
	tmp, err := NewMSV1_0TraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser{
		MSV1_0Trace: tmp,
	}
	return
}
