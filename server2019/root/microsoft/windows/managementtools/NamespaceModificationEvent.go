// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceModificationEvent struct
type __NamespaceModificationEvent struct {
	*__NamespaceOperationEvent

	//
	PreviousNamespace __Namespace
}

func New__NamespaceModificationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceModificationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceModificationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceModificationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceModificationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

// SetPreviousNamespace sets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) SetPropertyPreviousNamespace(value __Namespace) (err error) {
	return instance.SetProperty("PreviousNamespace", value)
}

// GetPreviousNamespace gets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) GetPropertyPreviousNamespace() (value __Namespace, err error) {
	retValue, err := instance.GetProperty("PreviousNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(__Namespace)
	if !ok {
		// TODO: Set an error
	}
	return
}
