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

// MLNX_Card struct
type MLNX_Card struct {
	*CIM_Card
}

func NewMLNX_CardEx1(instance *cim.WmiInstance) (newInstance *MLNX_Card, err error) {
	tmp, err := NewCIM_CardEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_Card{
		CIM_Card: tmp,
	}
	return
}

func NewMLNX_CardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_Card, err error) {
	tmp, err := NewCIM_CardEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_Card{
		CIM_Card: tmp,
	}
	return
}
