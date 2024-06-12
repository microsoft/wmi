// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSC2FA7540_DF00_428F_9858_1AFA6AE34276.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceDeletionEvent struct
type __InstanceDeletionEvent struct {
	*__InstanceOperationEvent
}

func New__InstanceDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceDeletionEvent, err error) {
	tmp, err := New__InstanceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceDeletionEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

func New__InstanceDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceDeletionEvent, err error) {
	tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceDeletionEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}
