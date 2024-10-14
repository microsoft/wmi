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

// OpenTrustedDomain_Start struct
type OpenTrustedDomain_Start struct {
	*OpenTrustedDomain
}

func NewOpenTrustedDomain_StartEx1(instance *cim.WmiInstance) (newInstance *OpenTrustedDomain_Start, err error) {
	tmp, err := NewOpenTrustedDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OpenTrustedDomain_Start{
		OpenTrustedDomain: tmp,
	}
	return
}

func NewOpenTrustedDomain_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OpenTrustedDomain_Start, err error) {
	tmp, err := NewOpenTrustedDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OpenTrustedDomain_Start{
		OpenTrustedDomain: tmp,
	}
	return
}
