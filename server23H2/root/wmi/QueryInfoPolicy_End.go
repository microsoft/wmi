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

// QueryInfoPolicy_End struct
type QueryInfoPolicy_End struct {
	*QueryInfoPolicy
}

func NewQueryInfoPolicy_EndEx1(instance *cim.WmiInstance) (newInstance *QueryInfoPolicy_End, err error) {
	tmp, err := NewQueryInfoPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy_End{
		QueryInfoPolicy: tmp,
	}
	return
}

func NewQueryInfoPolicy_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryInfoPolicy_End, err error) {
	tmp, err := NewQueryInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy_End{
		QueryInfoPolicy: tmp,
	}
	return
}
