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

// SamSetSecurityObj_End struct
type SamSetSecurityObj_End struct {
	*SamSetSecurityObj
}

func NewSamSetSecurityObj_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetSecurityObj_End, err error) {
	tmp, err := NewSamSetSecurityObjEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetSecurityObj_End{
		SamSetSecurityObj: tmp,
	}
	return
}

func NewSamSetSecurityObj_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetSecurityObj_End, err error) {
	tmp, err := NewSamSetSecurityObjEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetSecurityObj_End{
		SamSetSecurityObj: tmp,
	}
	return
}
