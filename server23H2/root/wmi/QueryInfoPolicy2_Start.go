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

// QueryInfoPolicy2_Start struct
type QueryInfoPolicy2_Start struct {
	*QueryInfoPolicy2
}

func NewQueryInfoPolicy2_StartEx1(instance *cim.WmiInstance) (newInstance *QueryInfoPolicy2_Start, err error) {
	tmp, err := NewQueryInfoPolicy2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy2_Start{
		QueryInfoPolicy2: tmp,
	}
	return
}

func NewQueryInfoPolicy2_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryInfoPolicy2_Start, err error) {
	tmp, err := NewQueryInfoPolicy2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryInfoPolicy2_Start{
		QueryInfoPolicy2: tmp,
	}
	return
}
