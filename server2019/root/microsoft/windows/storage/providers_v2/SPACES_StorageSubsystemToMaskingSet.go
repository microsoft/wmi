// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageSubsystemToMaskingSet struct
type SPACES_StorageSubsystemToMaskingSet struct {
	*MSFT_StorageSubSystemToMaskingSet
}

func NewSPACES_StorageSubsystemToMaskingSetEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToMaskingSet, err error) {
	tmp, err := NewMSFT_StorageSubSystemToMaskingSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToMaskingSet{
		MSFT_StorageSubSystemToMaskingSet: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemToMaskingSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystemToMaskingSet, err error) {
	tmp, err := NewMSFT_StorageSubSystemToMaskingSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToMaskingSet{
		MSFT_StorageSubSystemToMaskingSet: tmp,
	}
	return
}
