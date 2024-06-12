// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS6E517131_DD09_4014_BAEA_F9EF96C11D09
//////////////////////////////////////////////
package ns6e517131_dd09_4014_baea_f9ef96c11d09

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
