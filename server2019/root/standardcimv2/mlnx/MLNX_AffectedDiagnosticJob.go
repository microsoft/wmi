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

// MLNX_AffectedDiagnosticJob struct
type MLNX_AffectedDiagnosticJob struct {
	*CIM_AffectedJobElement
}

func NewMLNX_AffectedDiagnosticJobEx1(instance *cim.WmiInstance) (newInstance *MLNX_AffectedDiagnosticJob, err error) {
	tmp, err := NewCIM_AffectedJobElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_AffectedDiagnosticJob{
		CIM_AffectedJobElement: tmp,
	}
	return
}

func NewMLNX_AffectedDiagnosticJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_AffectedDiagnosticJob, err error) {
	tmp, err := NewCIM_AffectedJobElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_AffectedDiagnosticJob{
		CIM_AffectedJobElement: tmp,
	}
	return
}
