// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetEventSession_Provider struct
type MSFT_NetEventSession_Provider struct {
	*CIM_Component
}

func NewMSFT_NetEventSession_ProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventSession_Provider, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventSession_Provider{
		CIM_Component: tmp,
	}
	return
}

func NewMSFT_NetEventSession_ProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventSession_Provider, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventSession_Provider{
		CIM_Component: tmp,
	}
	return
}
