// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageAlertEvent struct
type SPACES_StorageAlertEvent struct {
	*WSP_StorageAlertEvent
}

func NewSPACES_StorageAlertEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageAlertEvent, err error) {
	tmp, err := NewWSP_StorageAlertEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageAlertEvent{
		WSP_StorageAlertEvent: tmp,
	}
	return
}

func NewSPACES_StorageAlertEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageAlertEvent, err error) {
	tmp, err := NewWSP_StorageAlertEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageAlertEvent{
		WSP_StorageAlertEvent: tmp,
	}
	return
}
