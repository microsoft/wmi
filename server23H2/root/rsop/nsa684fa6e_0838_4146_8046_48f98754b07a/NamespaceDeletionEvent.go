// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA684FA6E_0838_4146_8046_48F98754B07A
//////////////////////////////////////////////
package nsa684fa6e_0838_4146_8046_48f98754b07a

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
