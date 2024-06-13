// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.Security.MicrosoftVolumeEncryption
//////////////////////////////////////////////
package microsoftvolumeencryption

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ProviderEx struct
type Win32_ProviderEx struct {
	*__Win32Provider
}

func NewWin32_ProviderExEx1(instance *cim.WmiInstance) (newInstance *Win32_ProviderEx, err error) {
	tmp, err := New__Win32ProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ProviderEx{
		__Win32Provider: tmp,
	}
	return
}

func NewWin32_ProviderExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ProviderEx, err error) {
	tmp, err := New__Win32ProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ProviderEx{
		__Win32Provider: tmp,
	}
	return
}
