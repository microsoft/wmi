// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// AMLIEvalData1 struct
type AMLIEvalData1 struct {
	*ACPITrace
}

func NewAMLIEvalData1Ex1(instance *cim.WmiInstance) (newInstance *AMLIEvalData1, err error) {
	tmp, err := NewACPITraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AMLIEvalData1{
		ACPITrace: tmp,
	}
	return
}

func NewAMLIEvalData1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AMLIEvalData1, err error) {
	tmp, err := NewACPITraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AMLIEvalData1{
		ACPITrace: tmp,
	}
	return
}
