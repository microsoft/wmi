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

// SamOpenAlias_End struct
type SamOpenAlias_End struct {
	*SamOpenAlias
}

func NewSamOpenAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamOpenAlias_End, err error) {
	tmp, err := NewSamOpenAliasEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamOpenAlias_End{
		SamOpenAlias: tmp,
	}
	return
}

func NewSamOpenAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamOpenAlias_End, err error) {
	tmp, err := NewSamOpenAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamOpenAlias_End{
		SamOpenAlias: tmp,
	}
	return
}
