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

// MsftSil_Data struct
type MsftSil_Data struct {
	*cim.WmiInstance
}

func NewMsftSil_DataEx1(instance *cim.WmiInstance) (newInstance *MsftSil_Data, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftSil_Data{
		WmiInstance: tmp,
	}
	return
}

func NewMsftSil_DataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftSil_Data, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftSil_Data{
		WmiInstance: tmp,
	}
	return
}
