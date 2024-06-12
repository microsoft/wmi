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

// EnumTrstedDomEx_End struct
type EnumTrstedDomEx_End struct {
	*EnumTrstedDomEx
}

func NewEnumTrstedDomEx_EndEx1(instance *cim.WmiInstance) (newInstance *EnumTrstedDomEx_End, err error) {
	tmp, err := NewEnumTrstedDomExEx1(instance)

	if err != nil {
		return
	}
	newInstance = &EnumTrstedDomEx_End{
		EnumTrstedDomEx: tmp,
	}
	return
}

func NewEnumTrstedDomEx_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *EnumTrstedDomEx_End, err error) {
	tmp, err := NewEnumTrstedDomExEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &EnumTrstedDomEx_End{
		EnumTrstedDomEx: tmp,
	}
	return
}
