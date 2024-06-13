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

// SamEnumUsersInDom_End struct
type SamEnumUsersInDom_End struct {
	*SamEnumUsersInDom
}

func NewSamEnumUsersInDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamEnumUsersInDom_End, err error) {
	tmp, err := NewSamEnumUsersInDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamEnumUsersInDom_End{
		SamEnumUsersInDom: tmp,
	}
	return
}

func NewSamEnumUsersInDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamEnumUsersInDom_End, err error) {
	tmp, err := NewSamEnumUsersInDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamEnumUsersInDom_End{
		SamEnumUsersInDom: tmp,
	}
	return
}
