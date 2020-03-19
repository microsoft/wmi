// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiStream struct
type Msft_MiStream struct {
	*cim.WmiInstance
}

func NewMsft_MiStreamEx1(instance *cim.WmiInstance) (newInstance *Msft_MiStream, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msft_MiStream{
		WmiInstance: tmp,
	}
	return
}

func NewMsft_MiStreamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiStream, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiStream{
		WmiInstance: tmp,
	}
	return
}
