// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfRawData struct
type Win32_PerfRawData struct {
	*Win32_Perf
}

func NewWin32_PerfRawDataEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData, err error) {
	tmp, err := NewWin32_PerfEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData{
		Win32_Perf: tmp,
	}
	return
}

func NewWin32_PerfRawDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData, err error) {
	tmp, err := NewWin32_PerfEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData{
		Win32_Perf: tmp,
	}
	return
}
