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

// IkeextTraceProvider struct
type IkeextTraceProvider struct {
	*EventTrace
}

func NewIkeextTraceProviderEx1(instance *cim.WmiInstance) (newInstance *IkeextTraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IkeextTraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewIkeextTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IkeextTraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IkeextTraceProvider{
		EventTrace: tmp,
	}
	return
}
