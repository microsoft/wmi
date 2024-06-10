// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// TransportSecurityBindingElement struct
type TransportSecurityBindingElement struct {
	*SecurityBindingElement
}

func NewTransportSecurityBindingElementEx1(instance *cim.WmiInstance) (newInstance *TransportSecurityBindingElement, err error) {
	tmp, err := NewSecurityBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransportSecurityBindingElement{
		SecurityBindingElement: tmp,
	}
	return
}

func NewTransportSecurityBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransportSecurityBindingElement, err error) {
	tmp, err := NewSecurityBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransportSecurityBindingElement{
		SecurityBindingElement: tmp,
	}
	return
}
