// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// UseManagedPresentationBindingElement struct
type UseManagedPresentationBindingElement struct {
	*BindingElement
}

func NewUseManagedPresentationBindingElementEx1(instance *cim.WmiInstance) (newInstance *UseManagedPresentationBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UseManagedPresentationBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewUseManagedPresentationBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UseManagedPresentationBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UseManagedPresentationBindingElement{
		BindingElement: tmp,
	}
	return
}
