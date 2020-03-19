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

// SPACES_StorageNode struct
type SPACES_StorageNode struct {
	*MSFT_StorageNode
}

func NewSPACES_StorageNodeEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageNode, err error) {
	tmp, err := NewMSFT_StorageNodeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNode{
		MSFT_StorageNode: tmp,
	}
	return
}

func NewSPACES_StorageNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageNode, err error) {
	tmp, err := NewMSFT_StorageNodeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNode{
		MSFT_StorageNode: tmp,
	}
	return
}
