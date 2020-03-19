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

// MLNX_ComputerSystem struct
type MLNX_ComputerSystem struct {
	*CIM_ComputerSystem
}

func NewMLNX_ComputerSystemEx1(instance *cim.WmiInstance) (newInstance *MLNX_ComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_ComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMLNX_ComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_ComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_ComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}
