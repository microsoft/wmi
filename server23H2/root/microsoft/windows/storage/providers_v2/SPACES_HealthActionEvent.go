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

// SPACES_HealthActionEvent struct
type SPACES_HealthActionEvent struct {
	*MSFT_HealthActionEvent
}

func NewSPACES_HealthActionEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_HealthActionEvent, err error) {
	tmp, err := NewMSFT_HealthActionEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_HealthActionEvent{
		MSFT_HealthActionEvent: tmp,
	}
	return
}

func NewSPACES_HealthActionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_HealthActionEvent, err error) {
	tmp, err := NewMSFT_HealthActionEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_HealthActionEvent{
		MSFT_HealthActionEvent: tmp,
	}
	return
}
