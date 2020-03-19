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

// MLNX_Realizes struct
type MLNX_Realizes struct {
	*CIM_Realizes
}

func NewMLNX_RealizesEx1(instance *cim.WmiInstance) (newInstance *MLNX_Realizes, err error) {
	tmp, err := NewCIM_RealizesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_Realizes{
		CIM_Realizes: tmp,
	}
	return
}

func NewMLNX_RealizesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_Realizes, err error) {
	tmp, err := NewCIM_RealizesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_Realizes{
		CIM_Realizes: tmp,
	}
	return
}
