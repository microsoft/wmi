// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageModificationEvent struct
type SPACES_StorageModificationEvent struct {
	*WSP_StorageModificationEvent
}

func NewSPACES_StorageModificationEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageModificationEvent, err error) {
	tmp, err := NewWSP_StorageModificationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageModificationEvent{
		WSP_StorageModificationEvent: tmp,
	}
	return
}

func NewSPACES_StorageModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageModificationEvent, err error) {
	tmp, err := NewWSP_StorageModificationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageModificationEvent{
		WSP_StorageModificationEvent: tmp,
	}
	return
}
