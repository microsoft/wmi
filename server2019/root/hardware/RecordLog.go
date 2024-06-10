// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Hardware
//
// ////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RecordLog struct
type RecordLog struct {
	*CIM_RecordLog
}

func NewRecordLogEx1(instance *cim.WmiInstance) (newInstance *RecordLog, err error) {
	tmp, err := NewCIM_RecordLogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RecordLog{
		CIM_RecordLog: tmp,
	}
	return
}

func NewRecordLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RecordLog, err error) {
	tmp, err := NewCIM_RecordLogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RecordLog{
		CIM_RecordLog: tmp,
	}
	return
}
