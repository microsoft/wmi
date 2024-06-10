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

// SamCloseHandle_End struct
type SamCloseHandle_End struct {
	*SamCloseHandle
}

func NewSamCloseHandle_EndEx1(instance *cim.WmiInstance) (newInstance *SamCloseHandle_End, err error) {
	tmp, err := NewSamCloseHandleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamCloseHandle_End{
		SamCloseHandle: tmp,
	}
	return
}

func NewSamCloseHandle_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamCloseHandle_End, err error) {
	tmp, err := NewSamCloseHandleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamCloseHandle_End{
		SamCloseHandle: tmp,
	}
	return
}
