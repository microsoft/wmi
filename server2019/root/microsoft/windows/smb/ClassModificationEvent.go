// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassModificationEvent struct
type __ClassModificationEvent struct {
	*__ClassOperationEvent

	//
	PreviousClass interface{}
}

func New__ClassModificationEventEx1(instance *cim.WmiInstance) (newInstance *__ClassModificationEvent, err error) {
	tmp, err := New__ClassOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassModificationEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

func New__ClassModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassModificationEvent, err error) {
	tmp, err := New__ClassOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassModificationEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

// SetPreviousClass sets the value of PreviousClass for the instance
func (instance *__ClassModificationEvent) SetPropertyPreviousClass(value interface{}) (err error) {
	return instance.SetProperty("PreviousClass", value)
}

// GetPreviousClass gets the value of PreviousClass for the instance
func (instance *__ClassModificationEvent) GetPropertyPreviousClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousClass")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
