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

// MSIde struct
type MSIde struct {
	*cim.WmiInstance
}

func NewMSIdeEx1(instance *cim.WmiInstance) (newInstance *MSIde, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSIde{
		WmiInstance: tmp,
	}
	return
}

func NewMSIdeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSIde, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSIde{
		WmiInstance: tmp,
	}
	return
}
