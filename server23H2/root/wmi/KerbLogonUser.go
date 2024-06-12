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

// KerbLogonUser struct
type KerbLogonUser struct {
	*MSKerbTrace
}

func NewKerbLogonUserEx1(instance *cim.WmiInstance) (newInstance *KerbLogonUser, err error) {
	tmp, err := NewMSKerbTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbLogonUser{
		MSKerbTrace: tmp,
	}
	return
}

func NewKerbLogonUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbLogonUser, err error) {
	tmp, err := NewMSKerbTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbLogonUser{
		MSKerbTrace: tmp,
	}
	return
}
