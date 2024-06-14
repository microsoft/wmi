// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// NtlmLogonUser_Start struct
type NtlmLogonUser_Start struct {
	*NtlmLogonUser
}

func NewNtlmLogonUser_StartEx1(instance *cim.WmiInstance) (newInstance *NtlmLogonUser_Start, err error) {
	tmp, err := NewNtlmLogonUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser_Start{
		NtlmLogonUser: tmp,
	}
	return
}

func NewNtlmLogonUser_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmLogonUser_Start, err error) {
	tmp, err := NewNtlmLogonUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmLogonUser_Start{
		NtlmLogonUser: tmp,
	}
	return
}
