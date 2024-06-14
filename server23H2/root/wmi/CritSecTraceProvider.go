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

// CritSecTraceProvider struct
type CritSecTraceProvider struct {
	*EventTrace
}

func NewCritSecTraceProviderEx1(instance *cim.WmiInstance) (newInstance *CritSecTraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CritSecTraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewCritSecTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CritSecTraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CritSecTraceProvider{
		EventTrace: tmp,
	}
	return
}
