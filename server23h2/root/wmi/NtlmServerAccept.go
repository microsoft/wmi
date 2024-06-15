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

// NtlmServerAccept struct
type NtlmServerAccept struct {
	*MSV1_0Trace
}

func NewNtlmServerAcceptEx1(instance *cim.WmiInstance) (newInstance *NtlmServerAccept, err error) {
	tmp, err := NewMSV1_0TraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmServerAccept{
		MSV1_0Trace: tmp,
	}
	return
}

func NewNtlmServerAcceptEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmServerAccept, err error) {
	tmp, err := NewMSV1_0TraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmServerAccept{
		MSV1_0Trace: tmp,
	}
	return
}
