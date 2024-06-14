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

// MSLSA_LookupIsolatedNameInTrustedDomains struct
type MSLSA_LookupIsolatedNameInTrustedDomains struct {
	*MSLSATrace
}

func NewMSLSA_LookupIsolatedNameInTrustedDomainsEx1(instance *cim.WmiInstance) (newInstance *MSLSA_LookupIsolatedNameInTrustedDomains, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSLSA_LookupIsolatedNameInTrustedDomains{
		MSLSATrace: tmp,
	}
	return
}

func NewMSLSA_LookupIsolatedNameInTrustedDomainsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSLSA_LookupIsolatedNameInTrustedDomains, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSLSA_LookupIsolatedNameInTrustedDomains{
		MSLSATrace: tmp,
	}
	return
}
