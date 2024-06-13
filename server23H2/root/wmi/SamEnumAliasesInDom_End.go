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

// SamEnumAliasesInDom_End struct
type SamEnumAliasesInDom_End struct {
	*SamEnumAliasesInDom
}

func NewSamEnumAliasesInDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamEnumAliasesInDom_End, err error) {
	tmp, err := NewSamEnumAliasesInDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamEnumAliasesInDom_End{
		SamEnumAliasesInDom: tmp,
	}
	return
}

func NewSamEnumAliasesInDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamEnumAliasesInDom_End, err error) {
	tmp, err := NewSamEnumAliasesInDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamEnumAliasesInDom_End{
		SamEnumAliasesInDom: tmp,
	}
	return
}
