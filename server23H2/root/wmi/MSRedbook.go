// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSRedbook struct
type MSRedbook struct {
	*cim.WmiInstance
}

func NewMSRedbookEx1(instance *cim.WmiInstance) (newInstance *MSRedbook, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSRedbook{
		WmiInstance: tmp,
	}
	return
}

func NewMSRedbookEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSRedbook, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSRedbook{
		WmiInstance: tmp,
	}
	return
}