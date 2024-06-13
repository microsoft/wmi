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

// SetInfoPolicy2_End struct
type SetInfoPolicy2_End struct {
	*SetInfoPolicy2
}

func NewSetInfoPolicy2_EndEx1(instance *cim.WmiInstance) (newInstance *SetInfoPolicy2_End, err error) {
	tmp, err := NewSetInfoPolicy2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SetInfoPolicy2_End{
		SetInfoPolicy2: tmp,
	}
	return
}

func NewSetInfoPolicy2_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetInfoPolicy2_End, err error) {
	tmp, err := NewSetInfoPolicy2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetInfoPolicy2_End{
		SetInfoPolicy2: tmp,
	}
	return
}
