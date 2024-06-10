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

// MLNX_DiagnosticUseOfLog struct
type MLNX_DiagnosticUseOfLog struct {
	*CIM_UseOfLog
}

func NewMLNX_DiagnosticUseOfLogEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticUseOfLog, err error) {
	tmp, err := NewCIM_UseOfLogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticUseOfLog{
		CIM_UseOfLog: tmp,
	}
	return
}

func NewMLNX_DiagnosticUseOfLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticUseOfLog, err error) {
	tmp, err := NewCIM_UseOfLogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticUseOfLog{
		CIM_UseOfLog: tmp,
	}
	return
}
