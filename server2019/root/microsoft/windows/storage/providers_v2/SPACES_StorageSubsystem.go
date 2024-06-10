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

// SPACES_StorageSubsystem struct
type SPACES_StorageSubsystem struct {
	*MSFT_StorageSubSystem
}

func NewSPACES_StorageSubsystemEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystem, err error) {
	tmp, err := NewMSFT_StorageSubSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystem{
		MSFT_StorageSubSystem: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystem, err error) {
	tmp, err := NewMSFT_StorageSubSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystem{
		MSFT_StorageSubSystem: tmp,
	}
	return
}
