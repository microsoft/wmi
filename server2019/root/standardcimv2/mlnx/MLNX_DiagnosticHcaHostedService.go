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

// MLNX_DiagnosticHcaHostedService struct
type MLNX_DiagnosticHcaHostedService struct {
	*MLNX_DiagnosticHostedService
}

func NewMLNX_DiagnosticHcaHostedServiceEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticHcaHostedService, err error) {
	tmp, err := NewMLNX_DiagnosticHostedServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticHcaHostedService{
		MLNX_DiagnosticHostedService: tmp,
	}
	return
}

func NewMLNX_DiagnosticHcaHostedServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticHcaHostedService, err error) {
	tmp, err := NewMLNX_DiagnosticHostedServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticHcaHostedService{
		MLNX_DiagnosticHostedService: tmp,
	}
	return
}
