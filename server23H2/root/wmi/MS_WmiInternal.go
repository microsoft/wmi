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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MS_WmiInternal struct
type MS_WmiInternal struct {
	*cim.WmiInstance
}

func NewMS_WmiInternalEx1(instance *cim.WmiInstance) (newInstance *MS_WmiInternal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_WmiInternal{
		WmiInstance: tmp,
	}
	return
}

func NewMS_WmiInternalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_WmiInternal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_WmiInternal{
		WmiInstance: tmp,
	}
	return
}
