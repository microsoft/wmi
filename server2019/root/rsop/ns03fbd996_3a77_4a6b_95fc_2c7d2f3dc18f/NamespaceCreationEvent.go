// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS03FBD996_3A77_4A6B_95FC_2C7D2F3DC18F
//////////////////////////////////////////////
package ns03fbd996_3a77_4a6b_95fc_2c7d2f3dc18f
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceCreationEvent struct
type __NamespaceCreationEvent struct { 
	*__NamespaceOperationEvent
}

	func New__NamespaceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceCreationEvent, err error) {tmp, err := New__NamespaceOperationEventEx1(instance)
		
	if err != nil { return }
	newInstance = &__NamespaceCreationEvent {
	__NamespaceOperationEvent: tmp,
	}
	return
	}
	

	func New__NamespaceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__NamespaceCreationEvent, err error) {tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__NamespaceCreationEvent {
	__NamespaceOperationEvent: tmp,
	}
	return
	}
	

