// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// LogRecord struct
type LogRecord struct {
	*CIM_LogRecord
}

func NewLogRecordEx1(instance *cim.WmiInstance) (newInstance *LogRecord, err error) {
	tmp, err := NewCIM_LogRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LogRecord{
		CIM_LogRecord: tmp,
	}
	return
}

func NewLogRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LogRecord, err error) {
	tmp, err := NewCIM_LogRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LogRecord{
		CIM_LogRecord: tmp,
	}
	return
}
