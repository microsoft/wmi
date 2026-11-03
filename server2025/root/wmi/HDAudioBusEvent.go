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

// HDAudioBusEvent struct
type HDAudioBusEvent struct {
	*WMIEvent
}

func NewHDAudioBusEventEx1(instance *cim.WmiInstance) (newInstance *HDAudioBusEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HDAudioBusEvent{
		WMIEvent: tmp,
	}
	return
}

func NewHDAudioBusEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HDAudioBusEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HDAudioBusEvent{
		WMIEvent: tmp,
	}
	return
}
