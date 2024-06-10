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

// SetDomainInfoPolicy_End struct
type SetDomainInfoPolicy_End struct {
	*SetDomainInfoPolicy
}

func NewSetDomainInfoPolicy_EndEx1(instance *cim.WmiInstance) (newInstance *SetDomainInfoPolicy_End, err error) {
	tmp, err := NewSetDomainInfoPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetDomainInfoPolicy_End{
		SetDomainInfoPolicy: tmp,
	}
	return
}

func NewSetDomainInfoPolicy_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetDomainInfoPolicy_End, err error) {
	tmp, err := NewSetDomainInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetDomainInfoPolicy_End{
		SetDomainInfoPolicy: tmp,
	}
	return
}
