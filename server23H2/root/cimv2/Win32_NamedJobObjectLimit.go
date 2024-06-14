// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_NamedJobObjectLimit struct
type Win32_NamedJobObjectLimit struct {
	*CIM_CollectionSetting
}

func NewWin32_NamedJobObjectLimitEx1(instance *cim.WmiInstance) (newInstance *Win32_NamedJobObjectLimit, err error) {
	tmp, err := NewCIM_CollectionSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectLimit{
		CIM_CollectionSetting: tmp,
	}
	return
}

func NewWin32_NamedJobObjectLimitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NamedJobObjectLimit, err error) {
	tmp, err := NewCIM_CollectionSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectLimit{
		CIM_CollectionSetting: tmp,
	}
	return
}