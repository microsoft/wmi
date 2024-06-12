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

// QryInfoTrustDom_Start struct
type QryInfoTrustDom_Start struct {
	*QryInfoTrustDom
}

func NewQryInfoTrustDom_StartEx1(instance *cim.WmiInstance) (newInstance *QryInfoTrustDom_Start, err error) {
	tmp, err := NewQryInfoTrustDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom_Start{
		QryInfoTrustDom: tmp,
	}
	return
}

func NewQryInfoTrustDom_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QryInfoTrustDom_Start, err error) {
	tmp, err := NewQryInfoTrustDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QryInfoTrustDom_Start{
		QryInfoTrustDom: tmp,
	}
	return
}
