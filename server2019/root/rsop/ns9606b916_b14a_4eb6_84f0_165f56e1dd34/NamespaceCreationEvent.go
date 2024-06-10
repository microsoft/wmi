// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34
//
// ////////////////////////////////////////////
package ns9606b916_b14a_4eb6_84f0_165f56e1dd34

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceCreationEvent struct
type __NamespaceCreationEvent struct {
	*__NamespaceOperationEvent
}

func New__NamespaceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceCreationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceCreationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceCreationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceCreationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}
