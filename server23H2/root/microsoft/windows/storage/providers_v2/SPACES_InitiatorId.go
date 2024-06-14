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

// SPACES_InitiatorId struct
type SPACES_InitiatorId struct {
	*MSFT_InitiatorId
}

func NewSPACES_InitiatorIdEx1(instance *cim.WmiInstance) (newInstance *SPACES_InitiatorId, err error) {
	tmp, err := NewMSFT_InitiatorIdEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_InitiatorId{
		MSFT_InitiatorId: tmp,
	}
	return
}

func NewSPACES_InitiatorIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_InitiatorId, err error) {
	tmp, err := NewMSFT_InitiatorIdEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_InitiatorId{
		MSFT_InitiatorId: tmp,
	}
	return
}
