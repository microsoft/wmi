// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS9E4D1FDD_7A96_4826_963E_F60A751B9B15
//////////////////////////////////////////////
package ns9e4d1fdd_7a96_4826_963e_f60a751b9b15

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
