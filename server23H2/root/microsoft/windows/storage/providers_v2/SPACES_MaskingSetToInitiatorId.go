// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Storage.Providers_v2
//
// ////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_MaskingSetToInitiatorId struct
type SPACES_MaskingSetToInitiatorId struct {
	*MSFT_MaskingSetToInitiatorId
}

func NewSPACES_MaskingSetToInitiatorIdEx1(instance *cim.WmiInstance) (newInstance *SPACES_MaskingSetToInitiatorId, err error) {
	tmp, err := NewMSFT_MaskingSetToInitiatorIdEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_MaskingSetToInitiatorId{
		MSFT_MaskingSetToInitiatorId: tmp,
	}
	return
}

func NewSPACES_MaskingSetToInitiatorIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_MaskingSetToInitiatorId, err error) {
	tmp, err := NewMSFT_MaskingSetToInitiatorIdEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_MaskingSetToInitiatorId{
		MSFT_MaskingSetToInitiatorId: tmp,
	}
	return
}
