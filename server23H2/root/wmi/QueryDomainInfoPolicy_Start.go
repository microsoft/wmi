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

// QueryDomainInfoPolicy_Start struct
type QueryDomainInfoPolicy_Start struct {
	*QueryDomainInfoPolicy
}

func NewQueryDomainInfoPolicy_StartEx1(instance *cim.WmiInstance) (newInstance *QueryDomainInfoPolicy_Start, err error) {
	tmp, err := NewQueryDomainInfoPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryDomainInfoPolicy_Start{
		QueryDomainInfoPolicy: tmp,
	}
	return
}

func NewQueryDomainInfoPolicy_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryDomainInfoPolicy_Start, err error) {
	tmp, err := NewQueryDomainInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryDomainInfoPolicy_Start{
		QueryDomainInfoPolicy: tmp,
	}
	return
}
