// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// BindingElement struct
type BindingElement struct {
	*cim.WmiInstance
}

func NewBindingElementEx1(instance *cim.WmiInstance) (newInstance *BindingElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BindingElement{
		WmiInstance: tmp,
	}
	return
}

func NewBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BindingElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BindingElement{
		WmiInstance: tmp,
	}
	return
}