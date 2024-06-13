// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageHealthStatusChangeEvent struct
type SPACES_StorageHealthStatusChangeEvent struct {
	*MSFT_StorageHealthStatusChangeEvent
}

func NewSPACES_StorageHealthStatusChangeEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageHealthStatusChangeEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageHealthStatusChangeEvent{
		MSFT_StorageHealthStatusChangeEvent: tmp,
	}
	return
}

func NewSPACES_StorageHealthStatusChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageHealthStatusChangeEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageHealthStatusChangeEvent{
		MSFT_StorageHealthStatusChangeEvent: tmp,
	}
	return
}
