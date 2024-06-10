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

// SPACES_StoragePoolToResiliencySetting struct
type SPACES_StoragePoolToResiliencySetting struct {
	*MSFT_StoragePoolToResiliencySetting
}

func NewSPACES_StoragePoolToResiliencySettingEx1(instance *cim.WmiInstance) (newInstance *SPACES_StoragePoolToResiliencySetting, err error) {
	tmp, err := NewMSFT_StoragePoolToResiliencySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StoragePoolToResiliencySetting{
		MSFT_StoragePoolToResiliencySetting: tmp,
	}
	return
}

func NewSPACES_StoragePoolToResiliencySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StoragePoolToResiliencySetting, err error) {
	tmp, err := NewMSFT_StoragePoolToResiliencySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StoragePoolToResiliencySetting{
		MSFT_StoragePoolToResiliencySetting: tmp,
	}
	return
}
