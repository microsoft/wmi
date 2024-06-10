// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MuiTraceData struct
type MuiTraceData struct {
	*MuiTrace
}

func NewMuiTraceDataEx1(instance *cim.WmiInstance) (newInstance *MuiTraceData, err error) {
	tmp, err := NewMuiTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MuiTraceData{
		MuiTrace: tmp,
	}
	return
}

func NewMuiTraceDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MuiTraceData, err error) {
	tmp, err := NewMuiTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MuiTraceData{
		MuiTrace: tmp,
	}
	return
}
