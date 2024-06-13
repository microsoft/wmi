// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerManagerServerComponentDescriptor struct
type MSFT_ServerManagerServerComponentDescriptor struct {
	*cim.WmiInstance
}

func NewMSFT_ServerManagerServerComponentDescriptorEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerServerComponentDescriptor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerServerComponentDescriptor{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerServerComponentDescriptor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerServerComponentDescriptor{
		WmiInstance: tmp,
	}
	return
}
