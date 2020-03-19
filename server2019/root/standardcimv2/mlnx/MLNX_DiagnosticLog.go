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

// MLNX_DiagnosticLog struct
type MLNX_DiagnosticLog struct {
	*CIM_DiagnosticLog
}

func NewMLNX_DiagnosticLogEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticLog, err error) {
	tmp, err := NewCIM_DiagnosticLogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticLog{
		CIM_DiagnosticLog: tmp,
	}
	return
}

func NewMLNX_DiagnosticLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticLog, err error) {
	tmp, err := NewCIM_DiagnosticLogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticLog{
		CIM_DiagnosticLog: tmp,
	}
	return
}
