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

// SPACES_StorageDepartureEvent struct
type SPACES_StorageDepartureEvent struct {
	*WSP_StorageDepartureEvent
}

func NewSPACES_StorageDepartureEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageDepartureEvent, err error) {
	tmp, err := NewWSP_StorageDepartureEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageDepartureEvent{
		WSP_StorageDepartureEvent: tmp,
	}
	return
}

func NewSPACES_StorageDepartureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageDepartureEvent, err error) {
	tmp, err := NewWSP_StorageDepartureEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageDepartureEvent{
		WSP_StorageDepartureEvent: tmp,
	}
	return
}
