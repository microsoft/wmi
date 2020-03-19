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

// MLNX_DiagnosticLogManagesRecord struct
type MLNX_DiagnosticLogManagesRecord struct {
	*CIM_LogManagesRecord
}

func NewMLNX_DiagnosticLogManagesRecordEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticLogManagesRecord, err error) {
	tmp, err := NewCIM_LogManagesRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticLogManagesRecord{
		CIM_LogManagesRecord: tmp,
	}
	return
}

func NewMLNX_DiagnosticLogManagesRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticLogManagesRecord, err error) {
	tmp, err := NewCIM_LogManagesRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticLogManagesRecord{
		CIM_LogManagesRecord: tmp,
	}
	return
}
