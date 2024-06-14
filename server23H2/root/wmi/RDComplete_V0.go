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

// RDComplete_V0 struct
type RDComplete_V0 struct {
	*EventTraceEvent_V0
}

func NewRDComplete_V0Ex1(instance *cim.WmiInstance) (newInstance *RDComplete_V0, err error) {
	tmp, err := NewEventTraceEvent_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &RDComplete_V0{
		EventTraceEvent_V0: tmp,
	}
	return
}

func NewRDComplete_V0Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RDComplete_V0, err error) {
	tmp, err := NewEventTraceEvent_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RDComplete_V0{
		EventTraceEvent_V0: tmp,
	}
	return
}
