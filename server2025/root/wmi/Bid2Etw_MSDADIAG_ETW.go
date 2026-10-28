// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Bid2Etw_MSDADIAG_ETW struct
type Bid2Etw_MSDADIAG_ETW struct {
	*EventTrace
}

func NewBid2Etw_MSDADIAG_ETWEx1(instance *cim.WmiInstance) (newInstance *Bid2Etw_MSDADIAG_ETW, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Bid2Etw_MSDADIAG_ETW{
		EventTrace: tmp,
	}
	return
}

func NewBid2Etw_MSDADIAG_ETWEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Bid2Etw_MSDADIAG_ETW, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Bid2Etw_MSDADIAG_ETW{
		EventTrace: tmp,
	}
	return
}
