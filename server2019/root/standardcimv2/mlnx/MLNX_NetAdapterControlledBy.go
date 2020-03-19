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

// MLNX_NetAdapterControlledBy struct
type MLNX_NetAdapterControlledBy struct {
	*CIM_ControlledBy
}

func NewMLNX_NetAdapterControlledByEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}

func NewMLNX_NetAdapterControlledByEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}
