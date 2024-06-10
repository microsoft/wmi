// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_PCIDeviceStatusEvent struct
type MLNX_PCIDeviceStatusEvent struct {
	*CIM_InstModification
}

func NewMLNX_PCIDeviceStatusEventEx1(instance *cim.WmiInstance) (newInstance *MLNX_PCIDeviceStatusEvent, err error) {
	tmp, err := NewCIM_InstModificationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceStatusEvent{
		CIM_InstModification: tmp,
	}
	return
}

func NewMLNX_PCIDeviceStatusEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_PCIDeviceStatusEvent, err error) {
	tmp, err := NewCIM_InstModificationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceStatusEvent{
		CIM_InstModification: tmp,
	}
	return
}
