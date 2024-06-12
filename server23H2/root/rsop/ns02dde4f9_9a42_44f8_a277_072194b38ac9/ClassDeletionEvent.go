// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS02DDE4F9_9A42_44F8_A277_072194B38AC9
//////////////////////////////////////////////
package ns02dde4f9_9a42_44f8_a277_072194b38ac9

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassDeletionEvent struct
type __ClassDeletionEvent struct {
	*__ClassOperationEvent
}

func New__ClassDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

func New__ClassDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}
