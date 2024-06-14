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

// KerbInitSecurityContext struct
type KerbInitSecurityContext struct {
	*MSKerbTrace
}

func NewKerbInitSecurityContextEx1(instance *cim.WmiInstance) (newInstance *KerbInitSecurityContext, err error) {
	tmp, err := NewMSKerbTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext{
		MSKerbTrace: tmp,
	}
	return
}

func NewKerbInitSecurityContextEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbInitSecurityContext, err error) {
	tmp, err := NewMSKerbTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext{
		MSKerbTrace: tmp,
	}
	return
}
