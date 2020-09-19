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

// WSP_ReplicationSettings struct
type WSP_ReplicationSettings struct {
	*MSFT_ReplicationSettings
}

func NewWSP_ReplicationSettingsEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationSettings, err error) {
	tmp, err := NewMSFT_ReplicationSettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationSettings{
		MSFT_ReplicationSettings: tmp,
	}
	return
}

func NewWSP_ReplicationSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationSettings, err error) {
	tmp, err := NewMSFT_ReplicationSettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationSettings{
		MSFT_ReplicationSettings: tmp,
	}
	return
}
