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

// SamUserCreate_End struct
type SamUserCreate_End struct {
	*SamUserCreate
}

func NewSamUserCreate_EndEx1(instance *cim.WmiInstance) (newInstance *SamUserCreate_End, err error) {
	tmp, err := NewSamUserCreateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamUserCreate_End{
		SamUserCreate: tmp,
	}
	return
}

func NewSamUserCreate_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamUserCreate_End, err error) {
	tmp, err := NewSamUserCreateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamUserCreate_End{
		SamUserCreate: tmp,
	}
	return
}
