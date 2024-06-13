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

// HeapTrace_V2 struct
type HeapTrace_V2 struct {
	*HeapTraceProvider
}

func NewHeapTrace_V2Ex1(instance *cim.WmiInstance) (newInstance *HeapTrace_V2, err error) {
	tmp, err := NewHeapTraceProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapTrace_V2{
		HeapTraceProvider: tmp,
	}
	return
}

func NewHeapTrace_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapTrace_V2, err error) {
	tmp, err := NewHeapTraceProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapTrace_V2{
		HeapTraceProvider: tmp,
	}
	return
}
