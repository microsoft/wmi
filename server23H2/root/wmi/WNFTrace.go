// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WNFTrace struct
type WNFTrace struct {
	*WnfProvider
}

func NewWNFTraceEx1(instance *cim.WmiInstance) (newInstance *WNFTrace, err error) {
	tmp, err := NewWnfProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WNFTrace{
		WnfProvider: tmp,
	}
	return
}

func NewWNFTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WNFTrace, err error) {
	tmp, err := NewWnfProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WNFTrace{
		WnfProvider: tmp,
	}
	return
}
