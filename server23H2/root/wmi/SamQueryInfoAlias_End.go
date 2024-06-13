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

// SamQueryInfoAlias_End struct
type SamQueryInfoAlias_End struct {
	*SamQueryInfoAlias
}

func NewSamQueryInfoAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamQueryInfoAlias_End, err error) {
	tmp, err := NewSamQueryInfoAliasEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoAlias_End{
		SamQueryInfoAlias: tmp,
	}
	return
}

func NewSamQueryInfoAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQueryInfoAlias_End, err error) {
	tmp, err := NewSamQueryInfoAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoAlias_End{
		SamQueryInfoAlias: tmp,
	}
	return
}
