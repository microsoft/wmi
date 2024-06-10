// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RDComplete struct
type RDComplete struct {
	*EventTraceEvent
}

func NewRDCompleteEx1(instance *cim.WmiInstance) (newInstance *RDComplete, err error) {
	tmp, err := NewEventTraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RDComplete{
		EventTraceEvent: tmp,
	}
	return
}

func NewRDCompleteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RDComplete, err error) {
	tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RDComplete{
		EventTraceEvent: tmp,
	}
	return
}
