// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SetInfoTrustedDom_End struct
type SetInfoTrustedDom_End struct {
	*SetInfoTrustedDom
}

func NewSetInfoTrustedDom_EndEx1(instance *cim.WmiInstance) (newInstance *SetInfoTrustedDom_End, err error) {
	tmp, err := NewSetInfoTrustedDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetInfoTrustedDom_End{
		SetInfoTrustedDom: tmp,
	}
	return
}

func NewSetInfoTrustedDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetInfoTrustedDom_End, err error) {
	tmp, err := NewSetInfoTrustedDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetInfoTrustedDom_End{
		SetInfoTrustedDom: tmp,
	}
	return
}
