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

// MLNX_DriverService struct
type MLNX_DriverService struct {
	*CIM_Service
}

func NewMLNX_DriverServiceEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverService{
		CIM_Service: tmp,
	}
	return
}

func NewMLNX_DriverServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverService{
		CIM_Service: tmp,
	}
	return
}
