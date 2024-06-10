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

// KerbSetPassword struct
type KerbSetPassword struct {
	*MSKerbTrace
}

func NewKerbSetPasswordEx1(instance *cim.WmiInstance) (newInstance *KerbSetPassword, err error) {
	tmp, err := NewMSKerbTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword{
		MSKerbTrace: tmp,
	}
	return
}

func NewKerbSetPasswordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbSetPassword, err error) {
	tmp, err := NewMSKerbTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword{
		MSKerbTrace: tmp,
	}
	return
}
