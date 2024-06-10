// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SetInfoTrustedDom struct
type SetInfoTrustedDom struct {
	*MSLSATrace
}

func NewSetInfoTrustedDomEx1(instance *cim.WmiInstance) (newInstance *SetInfoTrustedDom, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SetInfoTrustedDom{
		MSLSATrace: tmp,
	}
	return
}

func NewSetInfoTrustedDomEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SetInfoTrustedDom, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SetInfoTrustedDom{
		MSLSATrace: tmp,
	}
	return
}
