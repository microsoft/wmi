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

// MLNX_HostedService struct
type MLNX_HostedService struct {
	*CIM_HostedService
}

func NewMLNX_HostedServiceEx1(instance *cim.WmiInstance) (newInstance *MLNX_HostedService, err error) {
	tmp, err := NewCIM_HostedServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_HostedService{
		CIM_HostedService: tmp,
	}
	return
}

func NewMLNX_HostedServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_HostedService, err error) {
	tmp, err := NewCIM_HostedServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_HostedService{
		CIM_HostedService: tmp,
	}
	return
}
