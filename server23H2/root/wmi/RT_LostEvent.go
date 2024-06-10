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

// RT_LostEvent struct
type RT_LostEvent struct {
	*Lost_Event
}

func NewRT_LostEventEx1(instance *cim.WmiInstance) (newInstance *RT_LostEvent, err error) {
	tmp, err := NewLost_EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RT_LostEvent{
		Lost_Event: tmp,
	}
	return
}

func NewRT_LostEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RT_LostEvent, err error) {
	tmp, err := NewLost_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RT_LostEvent{
		Lost_Event: tmp,
	}
	return
}
