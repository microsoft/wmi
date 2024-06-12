// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamGetBootKeyInfo struct
type SamGetBootKeyInfo struct {
	*MSSAMTrace
}

func NewSamGetBootKeyInfoEx1(instance *cim.WmiInstance) (newInstance *SamGetBootKeyInfo, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetBootKeyInfo{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamGetBootKeyInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetBootKeyInfo, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetBootKeyInfo{
		MSSAMTrace: tmp,
	}
	return
}
