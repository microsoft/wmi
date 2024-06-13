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

// SamAddMemToAlias_End struct
type SamAddMemToAlias_End struct {
	*SamAddMemToAlias
}

func NewSamAddMemToAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamAddMemToAlias_End, err error) {
	tmp, err := NewSamAddMemToAliasEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamAddMemToAlias_End{
		SamAddMemToAlias: tmp,
	}
	return
}

func NewSamAddMemToAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamAddMemToAlias_End, err error) {
	tmp, err := NewSamAddMemToAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamAddMemToAlias_End{
		SamAddMemToAlias: tmp,
	}
	return
}
