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

// OpTrustedDomByName struct
type OpTrustedDomByName struct {
	*MSLSATrace
}

func NewOpTrustedDomByNameEx1(instance *cim.WmiInstance) (newInstance *OpTrustedDomByName, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OpTrustedDomByName{
		MSLSATrace: tmp,
	}
	return
}

func NewOpTrustedDomByNameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OpTrustedDomByName, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OpTrustedDomByName{
		MSLSATrace: tmp,
	}
	return
}
