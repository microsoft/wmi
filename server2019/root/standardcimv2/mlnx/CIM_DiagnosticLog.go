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

// CIM_DiagnosticLog struct
type CIM_DiagnosticLog struct {
	*CIM_RecordLog
}

func NewCIM_DiagnosticLogEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticLog, err error) {
	tmp, err := NewCIM_RecordLogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticLog{
		CIM_RecordLog: tmp,
	}
	return
}

func NewCIM_DiagnosticLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticLog, err error) {
	tmp, err := NewCIM_RecordLogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticLog{
		CIM_RecordLog: tmp,
	}
	return
}
