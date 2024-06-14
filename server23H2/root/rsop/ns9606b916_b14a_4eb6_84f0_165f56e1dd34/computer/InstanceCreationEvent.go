// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceCreationEvent struct
type __InstanceCreationEvent struct {
	*__InstanceOperationEvent
}

func New__InstanceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceCreationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceCreationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

func New__InstanceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceCreationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceCreationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}
