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

// SamCreateAliasInDom_End struct
type SamCreateAliasInDom_End struct {
	*SamCreateAliasInDom
}

func NewSamCreateAliasInDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamCreateAliasInDom_End, err error) {
	tmp, err := NewSamCreateAliasInDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamCreateAliasInDom_End{
		SamCreateAliasInDom: tmp,
	}
	return
}

func NewSamCreateAliasInDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamCreateAliasInDom_End, err error) {
	tmp, err := NewSamCreateAliasInDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamCreateAliasInDom_End{
		SamCreateAliasInDom: tmp,
	}
	return
}
