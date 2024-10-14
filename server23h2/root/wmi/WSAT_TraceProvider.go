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

// WSAT_TraceProvider struct
type WSAT_TraceProvider struct {
	*EventTrace
}

func NewWSAT_TraceProviderEx1(instance *cim.WmiInstance) (newInstance *WSAT_TraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSAT_TraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewWSAT_TraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSAT_TraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSAT_TraceProvider{
		EventTrace: tmp,
	}
	return
}
