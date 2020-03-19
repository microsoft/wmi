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

// MLNX_DiagnosticConcreteJob struct
type MLNX_DiagnosticConcreteJob struct {
	*CIM_ConcreteJob
}

func NewMLNX_DiagnosticConcreteJobEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticConcreteJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticConcreteJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMLNX_DiagnosticConcreteJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticConcreteJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticConcreteJob{
		CIM_ConcreteJob: tmp,
	}
	return
}
