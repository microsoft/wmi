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

// MLNX_OwningDiagnosticJobElement struct
type MLNX_OwningDiagnosticJobElement struct {
	*CIM_OwningJobElement
}

func NewMLNX_OwningDiagnosticJobElementEx1(instance *cim.WmiInstance) (newInstance *MLNX_OwningDiagnosticJobElement, err error) {
	tmp, err := NewCIM_OwningJobElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_OwningDiagnosticJobElement{
		CIM_OwningJobElement: tmp,
	}
	return
}

func NewMLNX_OwningDiagnosticJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_OwningDiagnosticJobElement, err error) {
	tmp, err := NewCIM_OwningJobElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_OwningDiagnosticJobElement{
		CIM_OwningJobElement: tmp,
	}
	return
}
