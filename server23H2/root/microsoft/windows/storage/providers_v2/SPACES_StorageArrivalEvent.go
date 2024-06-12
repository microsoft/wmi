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

// SPACES_StorageArrivalEvent struct
type SPACES_StorageArrivalEvent struct {
	*WSP_StorageArrivalEvent
}

func NewSPACES_StorageArrivalEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageArrivalEvent, err error) {
	tmp, err := NewWSP_StorageArrivalEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageArrivalEvent{
		WSP_StorageArrivalEvent: tmp,
	}
	return
}

func NewSPACES_StorageArrivalEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageArrivalEvent, err error) {
	tmp, err := NewWSP_StorageArrivalEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageArrivalEvent{
		WSP_StorageArrivalEvent: tmp,
	}
	return
}
