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

// QryInfoTrustDom struct
type QryInfoTrustDom struct {
	*MSLSATrace
}

func NewQryInfoTrustDomEx1(instance *cim.WmiInstance) (newInstance *QryInfoTrustDom, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom{
		MSLSATrace: tmp,
	}
	return
}

func NewQryInfoTrustDomEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QryInfoTrustDom, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom{
		MSLSATrace: tmp,
	}
	return
}
