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

// SPACES_ResiliencySetting struct
type SPACES_ResiliencySetting struct {
	*MSFT_ResiliencySetting
}

func NewSPACES_ResiliencySettingEx1(instance *cim.WmiInstance) (newInstance *SPACES_ResiliencySetting, err error) {
	tmp, err := NewMSFT_ResiliencySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_ResiliencySetting{
		MSFT_ResiliencySetting: tmp,
	}
	return
}

func NewSPACES_ResiliencySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_ResiliencySetting, err error) {
	tmp, err := NewMSFT_ResiliencySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_ResiliencySetting{
		MSFT_ResiliencySetting: tmp,
	}
	return
}
