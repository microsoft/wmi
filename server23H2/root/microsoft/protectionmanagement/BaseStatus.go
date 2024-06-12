// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.ProtectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// BaseStatus struct
type BaseStatus struct {
	*cim.WmiInstance
}

func NewBaseStatusEx1(instance *cim.WmiInstance) (newInstance *BaseStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BaseStatus{
		WmiInstance: tmp,
	}
	return
}

func NewBaseStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BaseStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BaseStatus{
		WmiInstance: tmp,
	}
	return
}
