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

// SamEnumUsersInDom struct
type SamEnumUsersInDom struct {
	*MSSAMTrace
}

func NewSamEnumUsersInDomEx1(instance *cim.WmiInstance) (newInstance *SamEnumUsersInDom, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamEnumUsersInDom{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamEnumUsersInDomEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamEnumUsersInDom, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamEnumUsersInDom{
		MSSAMTrace: tmp,
	}
	return
}
