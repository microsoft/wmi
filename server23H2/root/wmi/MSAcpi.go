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

// MSAcpi struct
type MSAcpi struct {
	*cim.WmiInstance
}

func NewMSAcpiEx1(instance *cim.WmiInstance) (newInstance *MSAcpi, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSAcpi{
		WmiInstance: tmp,
	}
	return
}

func NewMSAcpiEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSAcpi, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSAcpi{
		WmiInstance: tmp,
	}
	return
}
