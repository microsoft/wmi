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

// WSP_ReplicationCapabilities struct
type WSP_ReplicationCapabilities struct {
	*MSFT_ReplicationCapabilities
}

func NewWSP_ReplicationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationCapabilities, err error) {
	tmp, err := NewMSFT_ReplicationCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationCapabilities{
		MSFT_ReplicationCapabilities: tmp,
	}
	return
}

func NewWSP_ReplicationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationCapabilities, err error) {
	tmp, err := NewMSFT_ReplicationCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationCapabilities{
		MSFT_ReplicationCapabilities: tmp,
	}
	return
}
