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

// SetDomainInfoPolicy struct
type SetDomainInfoPolicy struct {
	*MSLSATrace
}

func NewSetDomainInfoPolicyEx1(instance *cim.WmiInstance) (newInstance *SetDomainInfoPolicy, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetDomainInfoPolicy{
		MSLSATrace: tmp,
	}
	return
}

func NewSetDomainInfoPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetDomainInfoPolicy, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetDomainInfoPolicy{
		MSLSATrace: tmp,
	}
	return
}
