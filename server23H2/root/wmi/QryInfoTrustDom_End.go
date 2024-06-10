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

// QryInfoTrustDom_End struct
type QryInfoTrustDom_End struct {
	*QryInfoTrustDom
}

func NewQryInfoTrustDom_EndEx1(instance *cim.WmiInstance) (newInstance *QryInfoTrustDom_End, err error) {
	tmp, err := NewQryInfoTrustDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom_End{
		QryInfoTrustDom: tmp,
	}
	return
}

func NewQryInfoTrustDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QryInfoTrustDom_End, err error) {
	tmp, err := NewQryInfoTrustDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom_End{
		QryInfoTrustDom: tmp,
	}
	return
}
