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

// WSP_HealthAction struct
type WSP_HealthAction struct {
	*MSFT_HealthAction
}

func NewWSP_HealthActionEx1(instance *cim.WmiInstance) (newInstance *WSP_HealthAction, err error) {
	tmp, err := NewMSFT_HealthActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_HealthAction{
		MSFT_HealthAction: tmp,
	}
	return
}

func NewWSP_HealthActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_HealthAction, err error) {
	tmp, err := NewMSFT_HealthActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_HealthAction{
		MSFT_HealthAction: tmp,
	}
	return
}
