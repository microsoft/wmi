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
