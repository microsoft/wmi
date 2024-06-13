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

// SamDeleteUser_End struct
type SamDeleteUser_End struct {
	*SamDeleteUser
}

func NewSamDeleteUser_EndEx1(instance *cim.WmiInstance) (newInstance *SamDeleteUser_End, err error) {
	tmp, err := NewSamDeleteUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamDeleteUser_End{
		SamDeleteUser: tmp,
	}
	return
}

func NewSamDeleteUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamDeleteUser_End, err error) {
	tmp, err := NewSamDeleteUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamDeleteUser_End{
		SamDeleteUser: tmp,
	}
	return
}
