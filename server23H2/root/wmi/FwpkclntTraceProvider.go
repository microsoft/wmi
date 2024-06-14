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

// FwpkclntTraceProvider struct
type FwpkclntTraceProvider struct {
	*EventTrace
}

func NewFwpkclntTraceProviderEx1(instance *cim.WmiInstance) (newInstance *FwpkclntTraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FwpkclntTraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewFwpkclntTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FwpkclntTraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FwpkclntTraceProvider{
		EventTrace: tmp,
	}
	return
}
