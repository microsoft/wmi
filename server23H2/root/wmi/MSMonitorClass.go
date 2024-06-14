// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSMonitorClass struct
type MSMonitorClass struct {
	*cim.WmiInstance
}

func NewMSMonitorClassEx1(instance *cim.WmiInstance) (newInstance *MSMonitorClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSMonitorClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSMonitorClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMonitorClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMonitorClass{
		WmiInstance: tmp,
	}
	return
}
