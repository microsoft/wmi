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

// BfeTraceProvider struct
type BfeTraceProvider struct {
	*EventTrace
}

func NewBfeTraceProviderEx1(instance *cim.WmiInstance) (newInstance *BfeTraceProvider, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BfeTraceProvider{
		EventTrace: tmp,
	}
	return
}

func NewBfeTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BfeTraceProvider, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BfeTraceProvider{
		EventTrace: tmp,
	}
	return
}
