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

// SamQuerySecurityObj_End struct
type SamQuerySecurityObj_End struct {
	*SamQuerySecurityObj
}

func NewSamQuerySecurityObj_EndEx1(instance *cim.WmiInstance) (newInstance *SamQuerySecurityObj_End, err error) {
	tmp, err := NewSamQuerySecurityObjEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQuerySecurityObj_End{
		SamQuerySecurityObj: tmp,
	}
	return
}

func NewSamQuerySecurityObj_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQuerySecurityObj_End, err error) {
	tmp, err := NewSamQuerySecurityObjEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQuerySecurityObj_End{
		SamQuerySecurityObj: tmp,
	}
	return
}
