// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_Synchronized struct
type WSP_Synchronized struct {
	*MSFT_Synchronized
}

func NewWSP_SynchronizedEx1(instance *cim.WmiInstance) (newInstance *WSP_Synchronized, err error) {
	tmp, err := NewMSFT_SynchronizedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_Synchronized{
		MSFT_Synchronized: tmp,
	}
	return
}

func NewWSP_SynchronizedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_Synchronized, err error) {
	tmp, err := NewMSFT_SynchronizedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_Synchronized{
		MSFT_Synchronized: tmp,
	}
	return
}
