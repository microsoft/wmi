// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Object struct
type Object struct {
	*cim.WmiInstance
}

func NewObjectEx1(instance *cim.WmiInstance) (newInstance *Object, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Object{
		WmiInstance: tmp,
	}
	return
}

func NewObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Object, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Object{
		WmiInstance: tmp,
	}
	return
}
