// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS3BC654AA_392A_4C67_BCCB_7765E6837951.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ExtrinsicEvent struct
type __ExtrinsicEvent struct {
	*__Event
}

func New__ExtrinsicEventEx1(instance *cim.WmiInstance) (newInstance *__ExtrinsicEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ExtrinsicEvent{
		__Event: tmp,
	}
	return
}

func New__ExtrinsicEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ExtrinsicEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ExtrinsicEvent{
		__Event: tmp,
	}
	return
}
