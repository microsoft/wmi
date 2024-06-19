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

// QueryDomainInfoPolicy_End struct
type QueryDomainInfoPolicy_End struct {
	*QueryDomainInfoPolicy
}

func NewQueryDomainInfoPolicy_EndEx1(instance *cim.WmiInstance) (newInstance *QueryDomainInfoPolicy_End, err error) {
	tmp, err := NewQueryDomainInfoPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QueryDomainInfoPolicy_End{
		QueryDomainInfoPolicy: tmp,
	}
	return
}

func NewQueryDomainInfoPolicy_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QueryDomainInfoPolicy_End, err error) {
	tmp, err := NewQueryDomainInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QueryDomainInfoPolicy_End{
		QueryDomainInfoPolicy: tmp,
	}
	return
}
