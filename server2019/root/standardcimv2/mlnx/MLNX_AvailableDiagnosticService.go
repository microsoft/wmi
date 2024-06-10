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

// MLNX_AvailableDiagnosticService struct
type MLNX_AvailableDiagnosticService struct {
	*CIM_AvailableDiagnosticService
}

func NewMLNX_AvailableDiagnosticServiceEx1(instance *cim.WmiInstance) (newInstance *MLNX_AvailableDiagnosticService, err error) {
	tmp, err := NewCIM_AvailableDiagnosticServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_AvailableDiagnosticService{
		CIM_AvailableDiagnosticService: tmp,
	}
	return
}

func NewMLNX_AvailableDiagnosticServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_AvailableDiagnosticService, err error) {
	tmp, err := NewCIM_AvailableDiagnosticServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_AvailableDiagnosticService{
		CIM_AvailableDiagnosticService: tmp,
	}
	return
}
