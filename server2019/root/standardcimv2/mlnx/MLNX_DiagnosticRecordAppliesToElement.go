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

// MLNX_DiagnosticRecordAppliesToElement struct
type MLNX_DiagnosticRecordAppliesToElement struct {
	*CIM_RecordAppliesToElement
}

func NewMLNX_DiagnosticRecordAppliesToElementEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticRecordAppliesToElement, err error) {
	tmp, err := NewCIM_RecordAppliesToElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticRecordAppliesToElement{
		CIM_RecordAppliesToElement: tmp,
	}
	return
}

func NewMLNX_DiagnosticRecordAppliesToElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticRecordAppliesToElement, err error) {
	tmp, err := NewCIM_RecordAppliesToElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticRecordAppliesToElement{
		CIM_RecordAppliesToElement: tmp,
	}
	return
}
