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

// SamGetMemInAlias_End struct
type SamGetMemInAlias_End struct {
	*SamGetMemInAlias
}

func NewSamGetMemInAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetMemInAlias_End, err error) {
	tmp, err := NewSamGetMemInAliasEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInAlias_End{
		SamGetMemInAlias: tmp,
	}
	return
}

func NewSamGetMemInAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetMemInAlias_End, err error) {
	tmp, err := NewSamGetMemInAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInAlias_End{
		SamGetMemInAlias: tmp,
	}
	return
}
