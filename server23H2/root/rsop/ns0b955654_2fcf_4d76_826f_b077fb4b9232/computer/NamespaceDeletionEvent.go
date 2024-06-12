// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS0B955654_2FCF_4D76_826F_B077FB4B9232.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceDeletionEvent struct
type __NamespaceDeletionEvent struct {
	*__NamespaceOperationEvent
}

func New__NamespaceDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}
