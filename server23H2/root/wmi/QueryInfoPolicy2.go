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

// QueryInfoPolicy2 struct
type QueryInfoPolicy2 struct {
	*MSLSATrace
}

func NewQueryInfoPolicy2Ex1(instance *cim.WmiInstance) (newInstance *QueryInfoPolicy2, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy2{
		MSLSATrace: tmp,
	}
	return
}

func NewQueryInfoPolicy2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryInfoPolicy2, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy2{
		MSLSATrace: tmp,
	}
	return
}
