// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// NtlmValidateUser_Start struct
type NtlmValidateUser_Start struct {
	*NtlmValidateUser
}

func NewNtlmValidateUser_StartEx1(instance *cim.WmiInstance) (newInstance *NtlmValidateUser_Start, err error) {
	tmp, err := NewNtlmValidateUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmValidateUser_Start{
		NtlmValidateUser: tmp,
	}
	return
}

func NewNtlmValidateUser_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmValidateUser_Start, err error) {
	tmp, err := NewNtlmValidateUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmValidateUser_Start{
		NtlmValidateUser: tmp,
	}
	return
}
