// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// QueryInfoPolicy_Start struct
type QueryInfoPolicy_Start struct {
	*QueryInfoPolicy
}

func NewQueryInfoPolicy_StartEx1(instance *cim.WmiInstance) (newInstance *QueryInfoPolicy_Start, err error) {
	tmp, err := NewQueryInfoPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy_Start{
		QueryInfoPolicy: tmp,
	}
	return
}

func NewQueryInfoPolicy_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryInfoPolicy_Start, err error) {
	tmp, err := NewQueryInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy_Start{
		QueryInfoPolicy: tmp,
	}
	return
}
