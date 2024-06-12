// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetLbfoProvider struct
type MSFT_NetLbfoProvider struct {
	*MSFT_NetImPlatProvider
}

func NewMSFT_NetLbfoProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLbfoProvider, err error) {
	tmp, err := NewMSFT_NetImPlatProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoProvider{
		MSFT_NetImPlatProvider: tmp,
	}
	return
}

func NewMSFT_NetLbfoProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLbfoProvider, err error) {
	tmp, err := NewMSFT_NetImPlatProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoProvider{
		MSFT_NetImPlatProvider: tmp,
	}
	return
}
