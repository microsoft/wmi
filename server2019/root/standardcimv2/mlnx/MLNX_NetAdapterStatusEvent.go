// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_NetAdapterStatusEvent struct
type MLNX_NetAdapterStatusEvent struct {
	*CIM_InstModification
}

func NewMLNX_NetAdapterStatusEventEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterStatusEvent, err error) {
	tmp, err := NewCIM_InstModificationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterStatusEvent{
		CIM_InstModification: tmp,
	}
	return
}

func NewMLNX_NetAdapterStatusEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterStatusEvent, err error) {
	tmp, err := NewCIM_InstModificationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterStatusEvent{
		CIM_InstModification: tmp,
	}
	return
}
