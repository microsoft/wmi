// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA7AA3F49_B28C_4B8E_8501_B2871A96094C
//////////////////////////////////////////////
package nsa7aa3f49_b28c_4b8e_8501_b2871a96094c

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SystemEvent struct
type __SystemEvent struct {
	*__ExtrinsicEvent
}

func New__SystemEventEx1(instance *cim.WmiInstance) (newInstance *__SystemEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__SystemEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func New__SystemEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SystemEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SystemEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}
