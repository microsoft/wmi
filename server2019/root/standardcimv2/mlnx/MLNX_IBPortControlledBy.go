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

// MLNX_IBPortControlledBy struct
type MLNX_IBPortControlledBy struct {
	*CIM_ControlledBy
}

func NewMLNX_IBPortControlledByEx1(instance *cim.WmiInstance) (newInstance *MLNX_IBPortControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPortControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}

func NewMLNX_IBPortControlledByEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_IBPortControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPortControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}
