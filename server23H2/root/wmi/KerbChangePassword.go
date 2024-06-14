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

// KerbChangePassword struct
type KerbChangePassword struct {
	*MSKerbTrace
}

func NewKerbChangePasswordEx1(instance *cim.WmiInstance) (newInstance *KerbChangePassword, err error) {
	tmp, err := NewMSKerbTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword{
		MSKerbTrace: tmp,
	}
	return
}

func NewKerbChangePasswordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbChangePassword, err error) {
	tmp, err := NewMSKerbTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword{
		MSKerbTrace: tmp,
	}
	return
}
